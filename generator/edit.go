package generator

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"coder/generator/library"
)

type EditCreator struct {
	CommonCreator
}

func NewEditCreator(tpl string) *EditCreator {
	creator := EditCreator{}
	creator.SetTpl(tpl)
	return &creator
}

func (m *EditCreator) Get() (data map[string]interface{}) {
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
	name := library.LowerWord(camelName)
	data["name"] = name
	data["jsName"] = template.JS(name)

	return data
}

func (m *EditCreator) Render(data IData) {
	name := filepath.Base(m.tpl)
	tmpl := template.New(name)
	funcMap := make(map[string]interface{})
	funcMap["js"] = func(item string) template.JS{
		return template.JS(item)
	}
	tmpl.Funcs(funcMap)
	tmpl.Delims("<%","%>")
	b, err := ioutil.ReadFile(m.tpl)
	if err != nil {
		println(err.Error())
	}
	s := string(b)
	tmpl, err = tmpl.Parse(s)
	if err != nil {
		println(err.Error())
		return
	}
	if m.p {
		tmpl.Execute(os.Stdout, data.Get())
	} else {
		os.MkdirAll(m.output,os.ModePerm)
		file, err := os.Create(m.output + "/edit.vue")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
}
