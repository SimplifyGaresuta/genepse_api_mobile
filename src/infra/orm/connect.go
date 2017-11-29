package orm

import (
	"genepse_api/src/infra/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func OpenMysql() (err error) {
	err = config.Read("mysql")
	if err != nil {
		return
	}
	db, err = gorm.Open("mysql", viper.GetString("user_name")+":"+viper.GetString("pass")+"@/"+viper.GetString("db_name")+"?charset="+viper.GetString("charset")+"&parseTime="+viper.GetString("parseTime")+"&loc="+viper.GetString("loc"))
	return
}
func CloseMysql() {
	db.Close()
}
