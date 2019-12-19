package providers

import (
	"coder/config"
	"xiaoshijie.com/go-xorm/xorm"
)

//var Models = []interface{}{ &models.Menu{}}
var (
	Engine *xorm.Engine
)

type XOrmDatabase struct {
}

func NewXOrmDatabase() *XOrmDatabase {
	return &XOrmDatabase{}
}

func (d *XOrmDatabase) Boot() {
	log.Infof("db %s \r\n", config.Project.Database)
	log.Infof("XOrmDatabaseType %s \r\n", config.Project.DatabaseType)
	eg, err := xorm.NewEngine(config.Project.DatabaseType, config.Project.Database)
	eg.ShowSQL(true)
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	eg.SetSoftDeleteHandler(&xorm.DefaultSoftDeleteHandler{})

	eg.SetMaxIdleConns(10)
	eg.SetMaxOpenConns(50)
	Engine = eg
}
func (d *XOrmDatabase) Register() {

}
func (d *XOrmDatabase) Shutdown() {
	Engine.Close()
}
func (d *XOrmDatabase)AfterBoot()  {

}