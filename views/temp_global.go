package views

import "embed"
import "html/template"

var (
	//go:embed *.html
	embedTmpl embed.FS
	// 自定义的函数必须在调用ParseFiles() ParseFS()之前创建。
	funcMap = template.FuncMap{}
	GoTpl   = template.Must(
		template.New("").
			Funcs(funcMap).
			ParseFS(embedTmpl, "*.html"))
)
