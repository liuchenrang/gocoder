package form

import (
	"github.com/polaris1119/goutils"
	"html/template"
	"coder/generator/contract"
)

type SelectElement struct {
	field Field
}
type Field struct {
	TagSettings  map[string]string
	FormSettings map[string]string
	Name         string
	DBName       string
	Kind         string
	Render       contract.ElementRender
}

func NewSelectElement(field Field) *SelectElement {
	return &SelectElement{field: field}
}


func (e *SelectElement) RenderElement() template.HTML {
	temp := `<el-select v-model="`+e.field.DBName+`" placeholder="请选择`+e.field.FormSettings["COMMENT"]+`">
    <el-option
      v-for="item in `+e.field.DBName+`_list"
      :key="item.value"
      :label="item.label"
      :value="item.value">
    </el-option>
  </el-select>`;
	return  template.HTML(temp)
	
}
func (e *SelectElement) RenderMethod() template.JS {
	methodJs := `onLoadSelect`+goutils.CamelName(e.field.Name)+`Data() {
      var that = this;
      this.indexResource(Api.menuResource).then(resp => {
        console.log(resp);
        that.`+e.field.DBName+`_list = resp.data.result;
      });
    },`
	if value,ok := e.field.FormSettings["DATA_TYPE"]; ok {
		if value == "1" {
			return template.JS(methodJs)
		}
	}
	return template.JS("")
	
}
func (e *SelectElement) RenderInit() template.JS  {
	methodJs := `this.onLoadSelect`+goutils.CamelName(e.field.Name)+"Data()\r\n"
	if value,ok := e.field.FormSettings["DATA_TYPE"]; ok {
		if value == "1" {
			return template.JS(methodJs)
		}
	}
	return template.JS("")
}
func (e *SelectElement) RenderData() template.JS {
	if value,ok := e.field.FormSettings["DATA_TYPE"]; ok {
		if value == "1" {
			return template.JS(e.field.DBName+"_list:[]")
		}else{
			if value, ok := e.field.FormSettings["DATA"]; ok{
				return template.JS(value)
			}
		}
	}
	return template.JS("")
}
