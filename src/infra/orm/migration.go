package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	//dropTable()
	createTable()
}

func dropTable() {
	db.DropTable(&User{}, &FacebookAccount{}, &Skill{}, &SkillUser{}, &Product{}, &ProductUser{}, &Award{})
}

func createTable() {
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
		if err := insertUser(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&FacebookAccount{}) {
		if err := db.CreateTable(&FacebookAccount{}).Error; err != nil {
			panic(err)
		}
		if err := insertFacebookAccount(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Skill{}) {
		if err := db.CreateTable(&Skill{}).Error; err != nil {
			panic(err)
		}
		if err := insertSkill(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&SkillUser{}) {
		if err := db.CreateTable(&SkillUser{}).Error; err != nil {
			panic(err)
		}
		if err := insertSkillUser(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Product{}) {
		if err := db.CreateTable(&Product{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ProductUser{}) {
		if err := db.CreateTable(&ProductUser{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Award{}) {
		if err := db.CreateTable(&Award{}).Error; err != nil {
			panic(err)
		}
	}

}

func insertUser() (err error) {
	users := []User{
		User{
			Name:              "中尾涼",
			AvatarUrl:         "https://scontent.xx.fbcdn.net/v/t1.0-1/p50x50/14368752_227659240970532_4518352865855223562_n.jpg?oh=64beb0ab4d6b59bc1ab43c3ccf041976&oe=5AD6F164",
			AttributeId:       1,
			Overview:          "頑張ります。",
			License:           "TOEIC 900点",
			Gender:            1,
			Age:               20,
			Address:           "埼玉県さいたま市",
			SchoolCarrer:      "中央大学",
			ActivityBase:      "新宿",
			FacebookAccountId: 1,
		},
		User{
			Name:              "岩見建太",
			AvatarUrl:         "hey.com",
			AttributeId:       2,
			Gender:            1,
			Age:               23,
			FacebookAccountId: 2,
		},
	}
	for _, u := range users {
		if err = u.Insert(); err != nil {
			return
		}
	}
	return
}

func insertFacebookAccount() (err error) {
	facebooks := []FacebookAccount{
		FacebookAccount{
			AccountId: "429618797441241",
			MypageUrl: "fb.com",
		},
		FacebookAccount{
			AccountId: "291289898398981",
			MypageUrl: "eowioew.com",
		},
	}
	for _, f := range facebooks {
		if err = f.Insert(); err != nil {
			return
		}
	}
	return
}

func insertSkill() (err error) {
	skills := []Skill{
		Skill{
			Name: "ruby",
		},
		Skill{
			Name: "java",
		},
		Skill{
			Name: "python",
		},
		Skill{
			Name: "go",
		},
		Skill{
			Name: "mysql",
		},
	}
	for _, s := range skills {
		if err = s.Insert(); err != nil {
			return
		}
	}
	return
}

func insertSkillUser() (err error) {
	skillUsers := []SkillUser{
		SkillUser{
			SkillId:   1,
			UserId:    1,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   2,
			UserId:    1,
			DispOrder: 2,
		},
	}
	for _, s := range skillUsers {
		if err = s.Insert(); err != nil {
			return
		}
	}
	return
}
