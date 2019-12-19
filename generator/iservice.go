package generator

import (
	"coder/config"
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
)

type IServiceCreator struct {
	CommonCreator
}

func NewIServiceCreator(tpl string) *IServiceCreator {
	creator := IServiceCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *IServiceCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] = "Service" + goutils.CamelName(m.name)
	data["Name"] = goutils.CamelName(m.name)
	data["name"] = goutils.UnderscoreName(m.name)
	data["imports"] = config.Project.IServiceImport
	
	return data
}
func (m *IServiceCreator) Render(data IData) {
	//println(m.tpl)
	tmpl, err := template.ParseFiles(m.tpl)
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		file, err := os.Create(m.output + "/iservice" + goutils.CamelName(m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
