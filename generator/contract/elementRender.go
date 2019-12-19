package contract

import (
	"html/template"
)

type ElementRender interface {
	RenderElement() template.HTML
	RenderData() template.JS
	RenderMethod() template.JS
	RenderInit() template.JS
}
