package webapp

import (
	"bufio"
	"fmt"
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

type Guestbook struct{
	SignatureCount int
	Signatures []string
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
	//add call to getStrings
	signatures := getStrings("./packages/webapp/signatures.txt")
	//Display loaded signatures
	//fmt.Printf("%#v\n",signatures)

	//http/template is based on text/template package but has security features for html
	//create template using the contents of view.html
	html, err := template.ParseFiles("./packages/webapp/view.html")
	//report errors
	check(err)

	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures: signatures,
	}

	//write template content to HTTP response and pass data to template (currently nil)
	//to add data to template you add actions to template text
	//actions are {{}} - inside insert data or operation
	//{{.}} - this reference data passed into the execute method
	err = html.Execute(writer, guestbook)
	//check for errors
	check(err)
}

//handler func for new.html
func NewHandler(writer http.ResponseWriter, request *http.Request){
	//loads the contents of new.html as the text template
	html, err := template.ParseFiles("./packages/webapp/new.html")
	//report errors
	check(err)

	//write template tom response and return no data
	err = html.Execute(writer, nil)
	//check for errors
	check(err)
}

//handler for the form submission
func CreateHandler(writer http.ResponseWriter, request *http.Request){
	//forms pass data on the http.request
	//get the value of the form submission
	signature := request.FormValue("signature")

	//write out output
	//_,err := writer.Write([]byte(signature))

	//open signatures.txt file, creating it if it does not exist
	//add line of text to end of the file
	//close the file

	//options for opening the file
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	//open the file
	file, err := os.OpenFile("./packages/webapp/signatures.txt",options,os.FileMode(0600))

	//report errors
	check(err)

	//add line to file
	//returns number of bytes written and any error
	_, err = fmt.Fprintln(file,signature)

	//report errors
	check(err)

	//we dont use defer as we are writting to the file not reading from it
	//calling close on a file being written to can cause errors that we need to handle which we cannot do with defer
	err = file.Close()

	//report errors
	check(err)

	//redirects page to guestbook
	//takes http.ResponseWriter, *http.Request, path for browser, status code
	//StatusFound means request was successful
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

//html/template is based on the text/template package
//to check package go doc text/template Template.Execute - is type io.writer
//note second param is interface{} means you can pass any type into it
//go doc io writer - is an interface with write method
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
	//remember to use range to iterate over a slice of structs
	templateText := "{{range .}}Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}{{end}}"
	subscriber1 := Subscriber{Name: "Chris Williams", Rate: 4.99, Active: false}
	subscriber2 := Subscriber{Name: "Luna Williams", Rate: 2.33, Active: true}
	TextTemplateWithData(templateText,[]Subscriber{subscriber1,subscriber2})
}

//func to read file and return slice of strings
func getStrings(fileName string) []string{
	//create slice
	var lines []string
	//read file
	file, err := os.Open(fileName)
	//if err = file does not exist return nil instead of slice
	//template treats nil as empty data
	if os.IsNotExist(err){
		return nil
	}

	//check for any other errors
	check(err)

	//after all other operations are complete close file
	defer file.Close()

	//scan the file
	scanner := bufio.NewScanner(file)
	//for each line read in file
	for scanner.Scan(){
		//add file text to string slice
		lines = append(lines, scanner.Text())
	}

	//report any scanning error and exit
	check(scanner.Err())

	//return slice
	return lines
}


