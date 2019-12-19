package form

import (
	"html/template"
)

type CommonElement struct {
	field Field
}




func (e *CommonElement) RenderElement() template.HTML {
	
	return  template.HTML("")
	
}
func (e *CommonElement) RenderMethod() template.JS {

	return template.JS("")
	
}
func (e *CommonElement) RenderInit() template.JS  {

	return template.JS("")
}
func (e *CommonElement) RenderData() template.JS {

	return template.JS("")
}

