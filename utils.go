package main

import (
	"chit-chat/data"
	"fmt"
	"html/template"
	"net/http"
)

func generateHTML(writer http.ResponseWriter, threads []data.Thread, filenames ...string) {
	 var files []string
	 for _, file := range filenames {
	 	files = append(files, fmt.Sprintf("templates/%s.html", file))
	 }

	 templates := template.Must(template.ParseFiles(files...))
	 templates.ExecuteTemplate(writer,"layout", threads)
}