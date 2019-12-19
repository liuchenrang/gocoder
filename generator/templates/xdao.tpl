package daos

import (
     {{range $element := .imports}}
        "{{$element}}"
     {{end}}
)



type {{.CreatorName}} struct {
    conf *config.ServiceConfig
    db *xorm.Engine
}
func (c {{.CreatorName}}) Create(model models.{{.Name}})   (models.{{.Name}}, error)  {
    _,err := c.db.Insert(&model)
   	return model, err
}
func (c {{.CreatorName}}) Incr(id uint,attr map[string]interface{}) (bool, error) {
	if id >= 0 {
		session := c.db.Id(id)
		for k,v := range attr {
			session.Incr(k,v)
		}
		acc, err  := session.Update(models.{{.Name}}{})
		if err == nil && acc >0 {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, errors.New("id miss")
}
func (c {{.CreatorName}}) Delete(model models.{{.Name}})   (int, error)  {
  	if model.Id  {{printf "%s" .oplteq | unescaped}} 0 {
  		return 0,errors.New("id miss")
  	}
  	affected, err := c.db.Id(model.Id).Delete(&model)
  	if affected > 0 {
  		return int(affected), err
  	}
  	return 0, err
}

func (c {{.CreatorName}}) DeleteById(id uint) (int, error) {
   acc, err := c.db.Id(id).Delete(new(models.{{.Name}}))
   	return int(acc), err
}

func (c {{.CreatorName}}) Update(id uint, attr map[string]interface{})   (bool, error)  {
    if id >= 0 {
    		_, err := c.db.Id(id).Table(models.{{.Name}}{}).Update(attr)
    		if err != nil {
    			return false, err
    		} else {
    			return true, err
    		}
    }
    return false, errors.New("id miss")
}
func (c {{.CreatorName}}) UpdateByModel(model models.{{.Name}})   (bool, error)  {
    if model.Id >= 0 {
    		_, err := c.db.Id(model.Id).Update(model)
    		if err != nil {
    			return false, err;
    		} else {
    			return true, err
    		}
    }
    return false, errors.New("id miss")
}

func (c {{.CreatorName}}) Find(id uint)  ( *models.{{.Name}}, error) {
    var object models.{{.Name}}
    model := &object
    _, err := c.db.Id(id).Get(model)
    if err != nil {
        return nil, err
    }
    if model.Id > 0 {
        return model, nil
    }
    return nil, nil
}


func (c {{.CreatorName}}) FindByKey(name string,value interface{}) (*models.{{.Name}}, error) {
    model := &models.{{.Name}}{}
   	has,err := c.db.Where(name+" = ?", value).Limit(1).Get(model)
   	if err == nil {
   		if has {
   			return model, nil
   		}else{
   			return nil, nil
   		}
   	} else {
   		return model, err
   	}
}



func (c {{.CreatorName}}) FindAll(page components.Page) (*[]*models.{{.Name}}, error){
    var list []*models.{{.Name}}
   	limitInfo := page.ToLimitInfo()
   	err := c.db.OrderBy(page.ToOrder()).Limit(limitInfo.Limit, limitInfo.Offset).Find(&list)
   	if len(list) {{printf "%s" .oplteq | unescaped}} 0  {
   		return nil,nil
   	}
   	return &list, err
}
func (c {{.CreatorName}}) FindWithID(ids []uint) (*[]*models.{{.Name}}, error) {
        var list []*models.{{.Name}}
     	err := c.db.In("id",ids).Find(&list)
     	if len(list)  {{printf "%s" .oplteq | unescaped}} 0  {
     		return nil,nil
     	}
     	return &list, err
}

func (c {{.CreatorName}}) CountWhereForAdmin(wh string,bind ...interface{}) (int64,  error) {
    model := new(models.{{.Name}})
   	return  c.db.Where(wh, bind).Count(model)
}

func (c {{.CreatorName}}) FindWhereForAdmin(wh string,page components.Page,bind ...interface{})  (*[]*models.{{.Name}}, error) {
        limitInfo := page.ToLimitInfo()
    	var list []*models.{{.Name}}
    	err := c.db.Where(wh,bind).OrderBy(page.Order).Limit(limitInfo.Limit,limitInfo.Offset).Find(&list)
    	if len(list)  {{printf "%s" .oplteq | unescaped}} 0  {
    		return nil,nil
    	}
    	return &list, err
}



func (c {{.CreatorName}}) Count() (int64, error) {
    return c.db.Count(&models.{{.Name}}{})
}

func New{{.CreatorName}}(conf *config.ServiceConfig, db *xorm.Engine) *{{.CreatorName}} {
    return &{{.CreatorName}}{conf: conf, db: db}
}
