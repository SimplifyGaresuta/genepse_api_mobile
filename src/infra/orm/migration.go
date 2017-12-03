package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	//db.DropTable(&User{}, &FacebookAccount{}, &Skill{}, &SkillUser{})
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&FacebookAccount{}) {
		if err := db.CreateTable(&FacebookAccount{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&Skill{}) {
		if err := db.CreateTable(&Skill{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&SkillUser{}) {
		if err := db.CreateTable(&SkillUser{}).Error; err != nil {
			panic(err)
		}
	}
}
