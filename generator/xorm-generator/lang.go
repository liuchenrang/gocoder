package xorm_generator

import (
	"io/ioutil"
	"strings"
	"text/template"
	
	"github.com/go-xorm/core"
)

type LangTmpl struct {
	Funcs      template.FuncMap
	Formater   func(string) (string, error)
	GenImports func([]*core.Table) map[string]string
}

var (
	mapper    = &core.SnakeMapper{}
	langTmpls = map[string]LangTmpl{
		"go":   GoLangTmpl,
	}
)
type Tmpl struct {
	Tables  []*core.Table
	Imports map[string]string
	Models  string
}


func loadConfig(f string) map[string]string {
	bts, err := ioutil.ReadFile(f)
	if err != nil {
		return nil
	}
	configs := make(map[string]string)
	lines := strings.Split(string(bts), "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		vs := strings.Split(line, "=")
		if len(vs) == 2 {
			configs[strings.TrimSpace(vs[0])] = strings.TrimSpace(vs[1])
		}
	}
	return configs
}

func unTitle(src string) string {
	if src == "" {
		return ""
	}
	
	if len(src) == 1 {
		return strings.ToLower(string(src[0]))
	} else {
		return strings.ToLower(string(src[0])) + src[1:]
	}
}

func upTitle(src string) string {
	if src == "" {
		return ""
	}
	
	return strings.ToUpper(src)
}

