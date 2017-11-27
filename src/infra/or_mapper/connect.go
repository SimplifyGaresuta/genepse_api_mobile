package or_mapper

import (
	"genepse_api/src/infra/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	db, err = gorm.Open("mysql", config.UserName+":"+config.Pass+"@/"+config.DbName+"?charset="+config.Charset+"&parseTime="+config.ParseTime+"&loc="+config.Loc)
	return
}
