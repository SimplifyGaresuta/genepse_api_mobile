package orm

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenMysql() (err error) {
	if isDevelop() {
		db, err = gorm.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASS")+"@/"+os.Getenv("GENEPSE_DBNAME")+"?charset="+os.Getenv("MYSQL_CHARSET")+"&parseTime="+os.Getenv("MYSQL_PARSETIME")+"&loc="+os.Getenv("MYSQL_LOC"))
	} else {
		log.Println("ほんとに開くよ！")
		db, err = gorm.Open("mysql", os.Getenv("GENEPSE_MYSQL_CONNECTION"))
	}
	return
}
func CloseMysql() {
	db.Close()
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}
