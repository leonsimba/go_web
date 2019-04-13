package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 迭代动作示例
	t, _ := template.ParseFiles("tmpl-1.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	// 带备选结构的迭代示例
	// t, _ := template.ParseFiles("tmpl-2.html")	
	// daysOfWeek := []string{}

	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
