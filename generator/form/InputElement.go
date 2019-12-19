package form

import (
	"github.com/polaris1119/goutils"
	"html/template"
)

type InputElement struct {
	CommonElement
	field Field
}


func NewInputElement(field Field) *InputElement {
	return &InputElement{field: field}
}


func (e *InputElement) RenderElement() template.HTML {
	temp := `<el-select v-model="`+e.field.DBName+`" placeholder="请选择"`+e.field.TagSettings["COMMENT"]+`"">
    <el-option
      v-for="item in `+e.field.DBName+`_list"
      :key="item.value"
      :label="item.label"
      :value="item.value">
    </el-option>
  </el-select>`;
	return  template.HTML(temp)
	
}
func (e *InputElement) RenderMethod() template.JS {
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
func (e *InputElement) RenderInit() template.JS  {
	methodJs := `this.onLoadSelect`+goutils.CamelName(e.field.Name)+"Data()\r\n"
	if value,ok := e.field.FormSettings["DATA_TYPE"]; ok {
		if value == "1" {
			return template.JS(methodJs)
		}
	}
	return template.JS("")
}
func (e *InputElement) RenderData() template.JS {
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
