package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func openMysql() (err error) {
	db, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	db, err = gorm.Open("mysql", user+":"+pass+"@/"+dbName+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc)

}
