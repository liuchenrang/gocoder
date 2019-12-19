package generator

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
	"coder/providers"
)

type CommonCreator struct {
	tpl    string
	p      bool
	name   string
	output string
	modelPath string
	mapInfo map[string]*gorm.ModelStruct
}

func (m *CommonCreator) SetPreView(p bool) {
	m.p = p
}
func (m *CommonCreator) SetOutputPath(path string) {
	m.output = path
}
func (m *CommonCreator) SetCreateName(name string) {
	m.name = name
}
func (m *CommonCreator) SetTpl(name string) {
	m.tpl = name
}
func (m *CommonCreator) SetModelPath(modelPath string) {
	m.modelPath = modelPath
}
func (m *CommonCreator) parseFormTag(tag string) (formInfo map[string]string ) {
	formElementInfo := strings.Split(tag, ";")
	formInfo = make( map[string]string )
	for _,v := range  formElementInfo {
			attrInfo := strings.Split(v, ":")
		if len(attrInfo) == 2 {
			
			formInfo[attrInfo[0]] = attrInfo[1]
		}
	}
	return formInfo
}
func (m *CommonCreator) GetModelInfo(name string) *gorm.ModelStruct{

	
	if len(m.mapInfo) == 0 {
		m.mapInfo = make(map[string]*gorm.ModelStruct)
	
		for _,model := range providers.Models {
		
			name := reflect.TypeOf(model).Elem().Name()
			println("table info " + name + " init ")
			modelStruct := providers.Db.NewScope(model).GetModelStruct()
			j := 0
			for _,m :=  range modelStruct.StructFields  {
				if m.DBName == "id" || m.DBName == "created_at" || m.DBName == "updated_at" || m.DBName == "deleted_at" {
					j++
				}
			}
			i := len(modelStruct.StructFields)
			newFields := make([]*gorm.StructField, i-j)
			k := 0
			for _,structField := range modelStruct.StructFields  {
				
				
				attrInfo := m.parseFormTag(structField.Tag.Get("form"))
				if structField.DBName == "id" || structField.DBName == "created_at" || structField.DBName == "updated_at" || structField.DBName == "deleted_at" {
					continue
				}
				if val, ok  := structField.TagSettings["COMMENT"]; ok {
					attrInfo["COMMENT"] = strings.Trim(val, "\"'");
					structField.TagSettings["COMMENT"] = strings.Trim(val, "\"'");
				}
				structField.TagSettings = attrInfo
				newFields[k] = structField
				k++
			}
			modelStruct.StructFields = newFields
			m.mapInfo[name] = modelStruct
		}
	}
	
	return m.mapInfo[name];
}
