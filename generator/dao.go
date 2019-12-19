
package generator

import (
	"coder/config"
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
)

type DaoCreator struct {
	CommonCreator
}

func NewDaoCreator(tpl string) *DaoCreator {
	creator := DaoCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *DaoCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] =  goutils.CamelName(m.name) + "Dao"
	data["Name"] = goutils.CamelName(m.name)
	data["opLteq"] =  template.JS("<=")
	data["name"] = goutils.UnderscoreName(m.name)
	data["imports"] = config.Project.DaoImport
	return data
}
func (m *DaoCreator) Render(data IData) {
	tmpl, err := template.ParseFiles(m.tpl)
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		file, err := os.Create(m.output + "/" + (m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
