package daos

import (

)


var (
    {{.Name}} {{.CreatorName}}
)
type {{.CreatorName}} struct {
    conf *config.ServerConfig
    db *Db
}
func (c {{.CreatorName}}) Create(model models.{{.Name}})  (models.{{.Name}})  {
    c.Db.Save(&model)
    return model
}
func (c {{.CreatorName}}) Delete(model models.{{.Name}})  (models.{{.Name}})  {
    c.Db.Delete(&model)
    return model
}

func (c {{.CreatorName}}) DeleteById(id uint) bool {
    return c.Db.Where(" id = ?", id).Delete(&models.{{.Name}}{}).RowsAffected > 0
}
func (c {{.CreatorName}}) Save(model models.{{.Name}})  (models.{{.Name}})  {
    c.Db.Save(&model)
    return model
}

func (c {{.CreatorName}}) Update(id uint, attr map[string]interface{})  bool  {
    {{$op := "<="}}
    if id >= 0 {
        model := &models.{{.Name}}{}
        model.ID = id
        return c.Db.Model(model).Updates(attr).RowsAffected > 0
    }
    return false
}

func (c {{.CreatorName}}) Find(id uint)  (model models.{{.Name}})  {
    c.Db.First(&model, id)
    return model
}


func (c {{.CreatorName}}) FindByKey(name string,value interface{}) ({{.name}} models.{{.Name}}, noFound bool) {
    noFound = c.Db.First(&{{.name}}, name + " = ?", value).RecordNotFound()
    return {{.name}},noFound
}



func (c {{.CreatorName}}) FindAll(page components.Page) []models.{{.Name}} {
    var list []models.{{.Name}}
    limitInfo := page.ToLimitInfo()
    c.Db.Where("id > 0").Order(page.ToOrder()).Limit(limitInfo.Limit).Offset(limitInfo.Offset).Find(&list)
    return list
}
func (c {{.CreatorName}}) FindWithID(ids []uint) (models []models.{{.Name}}) {
     c.Db.Where(ids).Find(&models)
     return models
}

func (c {{.CreatorName}}) CountWhere(wh string,bind ...string) (models []models.{{.Name}}) {
     c.Db.Where(wh,bind).Find(&models)
     return models
}

func (c {{.CreatorName}}) FindWhere(wh string,page components.Page,bind ...string) (models []models.{{.Name}}) {
     limitInfo := page.ToLimitInfo()
     c.Db.Where(wh,bind).Order(page.ToOrder()).Limit(limitInfo.Limit).Offset(limitInfo.Offset).Find(&models)
     return models
}



func (c {{.CreatorName}}) Count() uint {
    var count uint
    c.Db.Model(&models.{{.Name}}{}).Where("id > 0").Count(&count)
    return count
}

func New{{.CreatorName}}(conf *config.ServerConfig) *{{.CreatorName}} {
    return &{{.CreatorName}}{conf: conf}
}
