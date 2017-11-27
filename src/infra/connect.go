package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	db, err = gorm.Open("mysql", user+":"+pass+"@/"+dbName+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc)
	return
}
