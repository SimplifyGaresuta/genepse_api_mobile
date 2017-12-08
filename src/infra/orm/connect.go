package orm

import (
	"fmt"
	"genepse_api/src/infra/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	log.Println("今から開くよ！")
	if isDevelop() {
		db, err = gorm.Open("mysql", config.MysqlUserName+":"+config.MysqlPass+"@/"+config.MysqlDbName+"?charset="+config.MysqlCharset+"&parseTime="+config.MysqlParseTime+"&loc="+config.MysqlLoc)
	} else {
		log.Println("ほんとに開くよ！")
		// TODO ここで詰まってる
		db, err = gorm.Open("mysql", fmt.Sprintf("root@cloudsql(%s:%s)/%s", config.ProjectID, config.InstanceName, config.DatabaseName))
	}
	return
}
func CloseMysql() {
	db.Close()
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}
