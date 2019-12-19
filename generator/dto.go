package generator

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
	"coder/generator/library"
)

type DtoCreator struct {
	CommonCreator
}

func NewDtoCreator(tpl string) *DtoCreator {
	creator := DtoCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *DtoCreator) Get() (data map[string]interface{}) {
	
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
	data["fields"] = m.GetModelInfo(camelName).StructFields
	return data
}

func (m *DtoCreator) Render(data IData) {
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
