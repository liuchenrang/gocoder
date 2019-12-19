package services

import (
 {{range $element := .imports}}
        {{$element | importFix}}
 {{end}}
)

var (
    {{.Name}} contractService.I{{.CreatorName}}
)

type {{.CreatorName}} struct {
    dao contractDao.I{{.Name}}Dao
    conf *config.ServiceConfig

}

func (c {{.CreatorName}}) Add(object models.{{.Name}})  (models.{{.Name}}, error) {
	return c.dao.Create(object)
}
func (c {{.CreatorName}}) Delete(object models.{{.Name}}) (int, error){
	return c.dao.Delete(object)
}
func (c {{.CreatorName}}) DeleteById(id uint) (int, error) {
	return c.dao.DeleteById(id)
}



func (c {{.CreatorName}}) Update(id uint,  attr map[string]interface{})  (bool, error) {
	return c.dao.Update(id, attr)
}
func (c {{.CreatorName}}) UpdateByModel(model models.{{.Name}})  (bool, error){
	return c.dao.UpdateByModel(model)
}

func (c {{.CreatorName}}) Find(id uint) ( *models.{{.Name}}, error){
	return c.dao.Find(id)
}
func (c {{.CreatorName}}) Incr(id uint,attr map[string]interface{}) (bool, error){
    return c.dao.Incr(id,attr)
}
func (c {{.CreatorName}}) FindAll(page components.Page) (*[]*models.{{.Name}}, error){
	return c.dao.FindAll(page)
}
func (c {{.CreatorName}}) FindWithID(ids []uint) (*[]*models.{{.Name}}, error){
	return c.dao.FindWithID(ids)
}





func (c {{.CreatorName}}) CountWhereForAdmin(wh string,bind ...interface{}) (int64,  error) {
     return  c.dao.CountWhereForAdmin(wh, bind...)
}

func (c {{.CreatorName}}) FindWhereForAdmin(wh string,page components.Page,bind ...interface{}) (*[]*models.{{.Name}}, error){
     return  c.dao.FindWhereForAdmin(wh,page, bind...)
}
func (c {{.CreatorName}}) Count() (int64, error) {
	return c.dao.Count()
}


func New{{.CreatorName}}(conf *config.ServiceConfig) {{.CreatorName}} {
	engine := conf.ConnectGroup.MysqlGroup["site_master"].(*xorm.Engine)

    return {{.CreatorName}}{conf: conf, dao: daos.New{{.Name}}Dao(conf,engine)}
}
