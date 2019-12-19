package form

type ApiForm{{.Name}} struct {
                ID   uint `form:"id"`

    {{range  $i,$v := .attrs}}
                {{$v.Name}}   {{$v.Kind}} `form:"{{$v.DBName}}" {{if eq (index $v.FormSettings "REQUIRE") "1"}}binding:"required"{{end}}`
    {{end}}
}