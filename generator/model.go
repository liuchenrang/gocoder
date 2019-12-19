package generator

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"os"
	"coder/generator/library"
)

type ModelCreator struct {
	CommonCreator
}


func NewModelCreator(tpl string) *ModelCreator {
	creator := ModelCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *ModelCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] =  goutils.CamelName(m.name) + "Dao"
	data["Name"] = goutils.CamelName(m.name)
	data["name"] = goutils.UnderscoreName(m.name)
	return data
}

func (m *ModelCreator) Render(data IData) {
	//println(m.tpl)
	tmpl, err := template.ParseFiles(m.tpl)
	//tmpl = tmpl.Funcs(template.FuncMap{"escaped": func(x string) interface{} { return x }})
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		file, err := os.Create(m.output + "/" + library.LowerWord(m.name) + ".go")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
