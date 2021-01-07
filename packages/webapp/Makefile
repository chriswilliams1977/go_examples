#these days requests go to a program running on a server that generates response for server to return
#net/http package handles requests to an app - included a small web server
#when you call http.HandleFunc with handler funcs upi are passing the handler func itself to be stored and called later

#in go functions are first class citizens meaning they can be assigned to vars then called by those vars
#when passing a func as param it has to be the right type
#for example http.HandleFunc expect func(ResponseWriter, *Request)
#a functions params and return value are part of its type
#a var that holds a function needs to specify its params and return values

#html/template package allows you to include html responses
#this package can load html in from and file and insert signatures
#template.Parsefiles takes a file
#it uses the content of this file to create a template value
#ParseFiles returns a pointer to this template and possibly an error
#to get output from template we call execute with two args
#1 - ResponseWriter place to write the text
#2 - data to insert into template