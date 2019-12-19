package generator

import (
	"coder/config"
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
)

type XDaoCreator struct {
	CommonCreator
}

func NewXDaoCreator(tpl string) *XDaoCreator {
	creator := XDaoCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *XDaoCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] =  goutils.CamelName(m.name) + "Dao"
	data["Name"] = goutils.CamelName(m.name)
	data["oplteq"] =  template.JS("<=")
	data["name"] = goutils.UnderscoreName(m.name)
	data["imports"] = config.Project.DaoImport
	
	return data
}
func unescaped (x string) interface{} { return template.HTML(x)}

func (m *XDaoCreator) Render(data IData) {
	tmpl := template.New("xdao.tpl")
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": unescaped})
	tmpl, _ = tmpl.ParseFiles(m.tpl)
	viewData := data.Get()
	if m.p {
		tmpl.Execute(os.Stdout, viewData)
	} else {
		file, err := os.Create(m.output + "/" + (m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = tmpl.Execute(file, viewData)
	}
	
}
