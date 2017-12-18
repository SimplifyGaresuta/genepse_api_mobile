package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	dropTable()
	createTable()
}

func dropTable() {
	db.DropTable(&User{}, &FacebookAccount{}, &Skill{}, &SkillUser{}, &Product{}, &Award{}, &License{})
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

	if !db.HasTable(&TwitterAccount{}) {
		if err := db.CreateTable(&TwitterAccount{}).Error; err != nil {
			panic(err)
		}
		if err := insertTwitterAccount(); err != nil {
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
		if err := insertProduct(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Award{}) {
		if err := db.CreateTable(&Award{}).Error; err != nil {
			panic(err)
		}
		if err := insertAward(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&License{}) {
		if err := db.CreateTable(&License{}).Error; err != nil {
			panic(err)
		}
		if err := insertLicense(); err != nil {
			panic(err)
		}
	}
}

func insertUser() (err error) {
	users := []User{
		User{
			Name:              "スティーブ・ジョブズ",
			AvatarUrl:         "https://i.gzn.jp/img/2016/01/06/blinq/a08.jpg",
			AttributeId:       2,
			Overview:          "頑張ります。",
			Gender:            1,
			Age:               20,
			Address:           "埼玉県さいたま市",
			SchoolCarrer:      "中央大学",
			ActivityBase:      "新宿",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "田中みな実",
			AvatarUrl:         "https://i2.wp.com/anincline.com/wp-content/uploads/2015/07/7b3cc41fc129710daee4f623415c93c6.png?fit=400%2C400",
			AttributeId:       2,
			Gender:            1,
			Age:               23,
			FacebookAccountId: 2,
		},
		User{
			Name:              "石原さとみ",
			AvatarUrl:         "https://i.pinimg.com/736x/17/54/d9/1754d97f559c4fc0dc1504ad59032140--satomi-ishihara-smartphone.jpg",
			AttributeId:       2,
			Overview:          "いえいえいえいえいえいえいえい！！！",
			Gender:            1,
			Age:               23,
			Address:           "埼玉県さいたま市",
			SchoolCarrer:      "明星大学",
			ActivityBase:      "恵比寿",
			FacebookAccountId: 3,
		},
		User{
			Name:              "未設定マン",
			Gender:            1,
			Age:               18,
			FacebookAccountId: 4,
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
		FacebookAccount{
			AccountId: "981239829389822",
			MypageUrl: "heyeheyehye.com",
		},
		FacebookAccount{
			AccountId: "283987237992971",
			MypageUrl: "ieieieiei.com",
		},
	}
	for _, f := range facebooks {
		if err = f.Insert(); err != nil {
			return
		}
	}
	return
}

func insertTwitterAccount() (err error) {
	twitters := []TwitterAccount{
		TwitterAccount{
			AccountId: "nakaoboy",
			MypageUrl: "https://twitter.com/nakaoboy",
		},
	}
	for _, f := range twitters {
		if err = f.Insert(); err != nil {
			return
		}
	}
	return
}

func insertSkill() (err error) {
	skills := []Skill{
		Skill{
			Name: "Ruby",
		},
		Skill{
			Name: "Java",
		},
		Skill{
			Name: "Python",
		},
		Skill{
			Name: "Go",
		},
		Skill{
			Name: "MySQL",
		},
		Skill{
			Name: "PHP",
		},
		Skill{
			Name: "AE",
		},
		Skill{
			Name: "営業",
		},
		Skill{
			Name: "NLP",
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
		SkillUser{
			SkillId:   3,
			UserId:    1,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   4,
			UserId:    1,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   4,
			UserId:    2,
			DispOrder: 1,
		},
	}
	for _, s := range skillUsers {
		if err = s.Insert(); err != nil {
			return
		}
	}
	return
}

func insertProduct() (err error) {
	products := []Product{
		Product{
			Title:        "リア充無双",
			UserId:       1,
			ReferenceUrl: "https://appsto.re/jp/26J0gb.i",
			ImageUrl:     "http://is2.mzstatic.com/image/thumb/Purple111/v4/27/d8/0c/27d80cef-fc79-c8ba-e18c-1b700dc79bc5/source/750x750bb.jpeg",
			DispOrder:    1,
		},
	}
	for _, p := range products {
		if err = p.Insert(); err != nil {
			return
		}
	}
	return
}

func insertAward() (err error) {
	awards := []Award{
		Award{
			UserId: 1,
			Name:   "ISUCON 優勝",
		},
		Award{
			UserId: 1,
			Name:   "アドテクチャレンジ 優勝",
		},
	}
	for _, a := range awards {
		if err = a.Insert(); err != nil {
			return
		}
	}
	return
}

func insertLicense() (err error) {
	licenses := []License{
		License{
			UserId: 1,
			Name:   "TOEIC 900点",
		},
		License{
			UserId: 1,
			Name:   "普通自動車運転免許",
		},
	}
	for _, l := range licenses {
		if err = l.Insert(); err != nil {
			return
		}
	}
	return
}
