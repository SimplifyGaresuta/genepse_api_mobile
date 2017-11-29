package orm

import (
	"genepse_api/src/infra/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	db, err = gorm.Open("mysql", config.MysqlUserName+":"+config.MysqlPass+"@/"+config.MysqlDbName+"?charset="+config.MysqlCharset+"&parseTime="+config.MysqlParseTime+"&loc="+config.MysqlLoc)
	return
}
func CloseMysql() {
	db.Close()
}
