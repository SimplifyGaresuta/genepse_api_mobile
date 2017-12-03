package orm

import (
	"fmt"
	"genepse_api/src/infra/config"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	if isDevelop() {
		db, err = gorm.Open("mysql", config.MysqlUserName+":"+config.MysqlPass+"@/"+config.MysqlDbName+"?charset="+config.MysqlCharset+"&parseTime="+config.MysqlParseTime+"&loc="+config.MysqlLoc)
	} else {
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
