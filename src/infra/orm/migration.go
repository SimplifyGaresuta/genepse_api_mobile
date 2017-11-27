package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	err := OpenMysql()
	if err != nil {
		log.Println(err)
	}
	defer CloseMysql()
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}
