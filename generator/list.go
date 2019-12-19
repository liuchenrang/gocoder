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

type ListCreator struct {
	CommonCreator
}

func NewListCreator(tpl string) *ListCreator {
	creator := ListCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *ListCreator) Get() (data map[string]interface{}) {
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
	name := library.LowerWord(m.name)
	data["name"] = name
	data["jsName"] = template.JS(name)
	return data
}

func (m *ListCreator) Render(data IData) {
	name := filepath.Base(m.tpl)
	tmpl := template.New(name)
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
		file, err := os.Create(m.output + "/base.vue")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmpl.Execute(file, data.Get())
	}
	
}
