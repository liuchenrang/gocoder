package providers

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"coder/config"

	"github.com/op/go-logging"
)

var Models = []interface{}{}
//var Models = []interface{}{ &models.Menu{}}
var (
	Db  *gorm.DB
	log = logging.MustGetLogger("controller")
)

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Boot() {
	log.Infof("db %s \r\n", config.Project.Database)
	log.Infof("DatabaseType %s \r\n", config.Project.DatabaseType)
	db, err := gorm.Open(config.Project.DatabaseType, config.Project.Database)

	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	db.LogMode(true)

	if err = db.AutoMigrate(Models...).Error; nil != err {
		print("auto migrate tables failed: " + err.Error())
	}


	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	Db = db
}
func (d *Database) Register() {

}
func (d *Database) Shutdown() {
	Db.Close()

}
func (d *Database)AfterBoot()  {

}