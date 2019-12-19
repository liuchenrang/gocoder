package generator

import "github.com/jinzhu/gorm"

type Generator struct {
	tplPath string
}

func NewGenerator(tplPath string) *Generator {
	return &Generator{tplPath: tplPath}
}
type Creator interface {
	Render(data IData)
	SetCreateName(name string)
	SetOutputPath(path string)
	SetPreView(p bool)
	SetModelPath(path string)
	GetModelInfo(name string) *gorm.ModelStruct
}

type IData interface {
	Get()  map[string]interface{}
}
func (g *Generator) GetInstance(creatorType string) Creator  {
	switch creatorType {
	case "xdao":
		return NewXDaoCreator(g.tplPath + "/xdao.tpl")
	case "dao":
		return NewDaoCreator(g.tplPath + "/dao.tpl")
	case "idao":
		return NewIDaoCreator(g.tplPath + "/idao.tpl")
	case "iservice":
		return NewIServiceCreator(g.tplPath + "/iservice.tpl")
	case "form":
		return NewFormCreator(g.tplPath + "/form.tpl")
	case "list":
		return NewListCreator(g.tplPath + "/list.tpl")
	case "edit":
		return NewEditCreator(g.tplPath + "/edit.tpl")
	case "dto":
		return NewDtoCreator(g.tplPath + "/dto.tpl")
	case "service":
		return NewServiceCreator(g.tplPath + "/service.tpl")
	case "xorm":
		return NewXOrmCreator(g.tplPath + "/xorm.go.tpl")
	case "controller":
		return NewControllerCreator(g.tplPath + "/controller.tpl")
	}
	return nil
}
