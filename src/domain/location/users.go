package location

import (
	"genepse_api/src/domain"
	"genepse_api/src/infra/cache"
	"genepse_api/src/infra/orm"
	"log"
	"math"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatar_url"`
	Attribute string   `json:"attribute"`
	Skills    []string `json:"skills"`
	Distance  int      `json:"distance"`
	Sns       []Sns    `json:"sns"`
}

type Sns struct {
	Provider string `json:"provider"`
	URL      string `json:"url"`
}

const query = `
select distinct u.id, u.name, u.avatar_url, u.attribute_id
from users as u left join skill_users as s on u.id=s.user_id
where u.id=? and (u.attribute_id != 0 and u.overview != "" and s.user_id is not null);`

func GetNearUsers(con *redis.Conn, userID string, distance int) (response *Response, err error) {
	ids, err := getUserIDs(con, userID, distance)
	if err != nil {
		c := *con
		log.Println("georadiusbymember時にエラー", c.Err())
		return
	}

	users, err := getUsers(con, ids, userID)

	response = &Response{
		Users: users,
	}
	return
}

func getUserIDs(con *redis.Conn, userID string, distance int) (userIDs []string, err error) {
	userIDs, err = cache.GeoRadiusByMember(con, key, userID, distance)
	return
}

func getUsers(con *redis.Conn, ids []string, userID string) (users []User, err error) {
	var i int
	for _, id := range ids {
		if id == userID {
			continue
		}
		distance, err := getDistance(con, userID, id, "m")
		if err != nil {
			log.Println(err)
			continue
		}
		i, err = strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			continue
		}

		rawUser := &orm.User{}
		// 位置情報だけ残ってるけどユーザーが削除されている可能性があるから
		if err = rawUser.RawQuery(query, i); err != nil {
			continue
		}
		skillNames, err := getSkills(i, 3)
		if err != nil {
			log.Println("ユーザーのスキル取得時にエラー", err)
		}

		fb := &orm.FacebookAccount{}
		fbURL, err := rawUser.ProviderURL(fb)
		if err != nil {
			log.Println(err)
		}
		tw := &orm.TwitterAccount{}
		twURL, err := rawUser.ProviderURL(tw)
		if err != nil {
			log.Println(err)
		}
		user := User{
			ID:        i,
			Name:      rawUser.Name,
			AvatarURL: rawUser.AvatarUrl,
			Attribute: domain.GetAttribute(rawUser.AttributeId),
			Skills:    skillNames,
			Distance:  distance,
			Sns:       []Sns{Sns{Provider: fb.ProviderName(), URL: fbURL}, Sns{Provider: tw.ProviderName(), URL: twURL}},
		}
		users = append(users, user)
	}
	return
}

// TODO アソシエーションしたら直す
func getSkills(userID, limit int) (skillNames []string, err error) {
	skillUsers := orm.SkillUsers{}
	if limit == 0 {
		if err = skillUsers.Where("user_id = ?", userID); err != nil {
			return
		}
	} else {
		if err = skillUsers.WhereLimit("user_id = ? and disp_order <= ?", 3, userID, limit); err != nil {
			return
		}
	}
	for _, skillUser := range skillUsers {
		skill := &orm.Skill{}
		if err = skill.Find(int(skillUser.SkillId)); err != nil {
			return
		}
		skillNames = append(skillNames, skill.Name)
	}
	return
}

func getDistance(con *redis.Conn, member1, member2, unit string) (int, error) {
	d, err := cache.GeoDist(con, key, member1, member2, unit)
	return int(math.Floor(d)), err
}
