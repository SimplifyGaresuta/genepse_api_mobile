package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	//db.DropTable(&User{}, &FacebookAccount{})
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

}
