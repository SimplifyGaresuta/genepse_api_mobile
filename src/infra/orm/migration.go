package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}
