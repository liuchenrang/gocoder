package generator

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
)


type ControllerCreator struct {
	CommonCreator
}

func NewControllerCreator(tpl string) *ControllerCreator {
	creator := ControllerCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *ControllerCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] = "Controller" + goutils.CamelName(m.name)
	data["Name"] = goutils.CamelName(m.name)
	data["name"] = goutils.UnderscoreName(m.name)
	return data
}

func (m *ControllerCreator) Render(data IData) {
	//println(m.tpl)
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
