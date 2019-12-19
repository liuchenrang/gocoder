package generator

import (
	"coder/config"
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
	"strings"
)

type ServiceCreator struct {
	CommonCreator
}

func NewServiceCreator(tpl string) *ServiceCreator {
	creator := ServiceCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *ServiceCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] = "Service" + goutils.CamelName(m.name)
	data["Name"] = goutils.CamelName(m.name)
	data["name"] = goutils.UnderscoreName(m.name)
	data["imports"] = config.Project.ServiceImport
	
	return data
}
func importFix (x string) interface{} {
	if strings.Index(x, " ") > -1 {
		return template.HTML(x)
	}else{
		return template.HTML("\"" + x + "\"")
		
	}
}

func (m *ServiceCreator) Render(data IData) {
	//println(m.tpl)
	tmpl := template.New("service.tpl")
	tmpl = tmpl.Funcs(template.FuncMap{"importFix": importFix})
	tmpl, err := tmpl.ParseFiles(m.tpl)
	
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		file, err := os.Create(m.output + "/service" + goutils.CamelName(m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
