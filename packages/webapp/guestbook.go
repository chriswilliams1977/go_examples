package webapp

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Subscriber struct{
	Name string
	Rate float64
	Active bool
}

//func for reporting errors
func check(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}

//handler func
//func type takes response and request
func ViewHandler(writer http.ResponseWriter, request *http.Request)  {
	//http/template is based on text/template package but has security features for html
	//create template using the contents of view.html
	html, err := template.ParseFiles("./packages/webapp/view.html")
	//report errors
	check(err)
	//write template content to HTTP response and pass data to template (currently nil)
	//to add data to template you add actions to template text
	//actions are {{}} - inside insert data or operation
	//{{.}} - this reference data passed into the execute method
	err = html.Execute(writer, nil)
	//check for errors
	check(err)
}

//html/template is based on the text/template package
//to check package go doc text/template Template.Execute - is type io.writer
//go doc io writer - is an interface with write method - note second param is interface{} means you can pass any type into it
//both http.ResponseWriter and os.Stdout have write method thus fulfil this interface
//thus can pass both to Template.Execute
func TextTemplateExample(text string){
	myText := text
	tmpl, err := template.New("test").Parse(myText)
	check(err)
	//pass os.stdout to the template execute
	//os.stdout writes data to terminal  - like a file that writes content to terminal
	//fmt.println ect use os.stdout
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}

//Example of how to pass data into text template
//
//you can also do conditional statements in template
//{{if .}}Dot is true{{end}},true - will print Dot is true
//{{if .}}Dot is true{{end}},false - will print blank
//
//can also use range to loop over slice, array, map or channel
//{{range}}{{end}}
//
//can also insert structs
//use struct field name {.Name}
func TextTemplateWithData(text string, data interface{}){
	//templateText := "Template start\nAction: {{.}}\nTemplate end \n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func TextTemplateWithStruct(){
	templateText := "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber1 := Subscriber{Name: "Chris Williams", Rate: 4.99, Active: false}
	subscriber2 := Subscriber{Name: "Luna Williams", Rate: 2.33, Active: true}
	TextTemplateWithData(templateText,subscriber1)
	TextTemplateWithData(templateText,subscriber2)
}

