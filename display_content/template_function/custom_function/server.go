package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}

	// 创建一个模板, 将funcMap与模板文件tmpl.html绑定
	t := template.New("tmpl.html").Funcs(funcMap)

	// 解析模板文件
	t, _ = t.ParseFiles("tmpl.html")

	// 执行模板
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
