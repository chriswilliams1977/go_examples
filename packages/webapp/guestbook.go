package webapp

import (
	"html/template"
	"log"
	"net/http"
)

//func for reporting errors
func check(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}

//handler func
//func type takes response and request
func ViewHandler(writer http.ResponseWriter, request *http.Request)  {
	//create template using the contents of view.html
	html, err := template.ParseFiles("./packages/webapp/view.html")
	//report errors
	check(err)
	//write template content to writer and pass data to template (currently nil)
	err = html.Execute(writer, nil)
	//check for errors
	check(err)
}
