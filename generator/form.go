package generator

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
	"coder/generator/library"
)

type FormCreator struct {
	CommonCreator
}

func NewFormCreator(tpl string) *FormCreator {
	creator := FormCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *FormCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	parse := library.ParseModel{}
	err := parse.Parse(m.modelPath + "/"+ m.name + ".go")
	if err != nil {
		println(err.Error())
		return nil
	}
	fields,_ := parse.GetFields()
	data["attrs"] = fields
	camelName := goutils.CamelName(m.name)
	data["CreatorName"] = camelName
	data["Name"] = camelName
	data["name"] = goutils.UnderscoreName(m.name)
	return data
}

func (m *FormCreator) Render(data IData) {
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
