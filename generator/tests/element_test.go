package tests

import (
	"fmt"
	"testing"
	"coder/generator/library"
)

func TestRender(t *testing.T)  {
	println(1)
	pr := library.ParseModel{}
	pr.Parse("/Users/chen/IdeaProjects/weibang/models/roleMenuMapping.go")
	fields, _ := pr.GetFields()
	fmt.Printf("%+v", fields)
}
