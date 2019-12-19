package dto

type {{.Name}} struct {
	ID             uint      `json:"id"`
    {{range  $i,$v := .attrs}}
                 {{$v.Name}}  {{$v.Kind}}    `json:"{{$v.DBName}}"`
    {{end}}
    CreatedAt      time.Time `json:"created_at"`

}