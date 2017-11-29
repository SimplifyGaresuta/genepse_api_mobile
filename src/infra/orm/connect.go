package orm

import (
	"genepse_api/src/infra/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func OpenMysql() (err error) {
	viper.SetConfigName("config")
	db, err = gorm.Open("mysql", config.UserName+":"+config.Pass+"@/"+config.DbName+"?charset="+config.Charset+"&parseTime="+config.ParseTime+"&loc="+config.Loc)
	return
}
func CloseMysql() {
	db.Close()
}
