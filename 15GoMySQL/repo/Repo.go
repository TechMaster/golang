package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB //Global Variable

func Connect(conf config.Configuration) {
	dsn := conf.Db.User + ":" + conf.Db.Pass + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMasterData() {
	//initCountry()
	//initCategory()
	//initManufacturer()
	//initProduct()
}
