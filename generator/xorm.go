package generator

import (
	xorm_generator "coder/generator/xorm-generator"
	"github.com/polaris1119/goutils"
	"os"
)

var (
	genJson                                      bool     = false
	ignoreColumnsJSON, created, updated, deleted []string = []string{}, []string{"created_at"}, []string{"updated_at"}, []string{"deleted_at"}
)

func dirExists(dir string) bool {
	d, e := os.Stat(dir)
	switch {
	case e != nil:
		return false
	case !d.IsDir():
		return false
	}
	
	return true
}

type XOrmCreator struct {
	CommonCreator
}

func NewXOrmCreator(tpl string) *XOrmCreator {
	creator := XOrmCreator{}
	creator.SetTpl(tpl)
	return &creator
}
func (m *XOrmCreator) Get() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["CreatorName"] = goutils.CamelName(m.name) + "Dao"
	data["Name"] = goutils.CamelName(m.name)
	data["name"] = goutils.UnderscoreName(m.name)
	return data
}

func (m *XOrmCreator) Render(data IData) {
	
	xorm_generator.Render(m.tpl, m.output, goutils.CamelName(m.name))
	
}
