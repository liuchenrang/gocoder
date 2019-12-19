
package generator

import (
	"coder/config"
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
)

type IDaoCreator struct {
	CommonCreator
}

func NewIDaoCreator(tpl string) *IDaoCreator {
	creator := IDaoCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *IDaoCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] =  goutils.CamelName(m.name) + "Dao"
	data["Name"] = goutils.CamelName(m.name)
	data["opLteq"] =  template.JS("<=")
	data["name"] = goutils.UnderscoreName(m.name)
	data["imports"] = config.Project.IDaoImport
	return data
}
func (m *IDaoCreator) Render(data IData) {
	tmpl, err := template.ParseFiles(m.tpl)
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		file, err := os.Create(m.output + "/" + "/idao" + goutils.CamelName(m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
