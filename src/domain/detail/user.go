package detail

import (
	"errors"
	"genepse_api/src/domain"
	"genepse_api/src/infra/orm"
	"log"
)

// User is 詳細画面に表示するユーザー
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	AvatarURL    string    `json:"avatar_url"`
	Attribute    string    `json:"attribute"`
	Skills       []string  `json:"skills"`
	Overview     string    `json:"overview"`
	Awards       []string  `json:"awards"`
	Products     []Product `json:"products"`
	Sns          []Sns     `json:"sns"`
	Licenses     []string  `json:"licenses"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	SchoolCareer string    `json:"school_career"`
	ActivityBase string    `json:"activity_base"`
}

type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Sns struct {
	Provider string `json:"provider"`
	URL      string `json:"url"`
}

func GetUser(id int) (user *User, err error) {
	rawUser := &orm.User{}
	if err = rawUser.Find(id); err != nil {
		return
	}
	gender := domain.GetGender(rawUser.Gender)
	products, err := getProducts(id)
	if err != nil {
		log.Println("ユーザーの作品取得時にエラー", err)
	}

	// TODO 抽象化
	facebookID, err := getFacebookURL(id)
	if err != nil {
		log.Println("ユーザーのfacebook取得時にエラー", err)
	}

	// TODO awardとlicenseで抽象化
	awardNames, err := getAwards(id)
	if err != nil {
		log.Println("ユーザーの受賞歴取得時にエラー", err)
	}

	licenseNames, err := getLicenses(id)
	if err != nil {
		log.Println("ユーザーの資格取得時にエラー", err)
	}

	skillNames, err := getSkills(id)
	if err != nil {
		log.Println("ユーザーのスキル取得時にエラー", err)
	}

	user = &User{
		ID:        int(rawUser.Model.ID),
		Name:      rawUser.Name,
		AvatarURL: rawUser.AvatarUrl,
		Attribute: domain.GetAttribute(rawUser.AttributeId),
		// TODO しっかり取る
		Skills:   skillNames,
		Overview: rawUser.Overview,
		Awards:   awardNames,
		Products: products,
		// TODO 抽象化
		Sns:          []Sns{Sns{Provider: "facebook", URL: facebookID}},
		Licenses:     licenseNames,
		Gender:       gender,
		Age:          rawUser.Age,
		Address:      rawUser.Address,
		SchoolCareer: rawUser.SchoolCarrer,
	}
	return
}

func getProducts(userID int) (products []Product, err error) {
	productUsers := orm.ProductUsers{}
	if err = productUsers.Where("user_id = ?", userID); err != nil {
		return
	}
	for _, productUser := range productUsers {
		p := &orm.Product{}
		if err = p.Find(int(productUser.Model.ID)); err != nil {
			return
		}
		products = append(products, Product{Title: p.Title, URL: p.ReferenceUrl})
	}
	return
}

// TODO providerで抽象化
func getFacebookURL(userID int) (url string, err error) {
	user := &orm.User{}
	if err = user.Find(userID); err != nil {
		return
	}
	fbID := user.FacebookAccountId
	if fbID == 0 {
		err = errors.New("facebookが登録されていません")
		return
	}
	fb := &orm.FacebookAccount{}
	if err = fb.Find(int(fbID)); err != nil {
		return
	}
	url = fb.MypageUrl
	return
}

func getAwards(userID int) (awardNames []string, err error) {
	awards := orm.Awards{}
	if err = awards.FindByUser(userID); err != nil {
		return
	}
	for _, award := range awards {
		awardNames = append(awardNames, award.Name)
	}
	return
}
func getLicenses(userID int) (licenseNames []string, err error) {
	licenses := orm.Licenses{}
	if err = licenses.FindByUser(userID); err != nil {
		return
	}
	for _, license := range licenses {
		licenseNames = append(licenseNames, license.Name)
	}
	return
}

// アソシエーションしたら直す
func getSkills(userID int) (skillNames []string, err error) {
	skillUsers := orm.SkillUsers{}
	if err = skillUsers.Where("user_id = ?", userID); err != nil {
		return
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
