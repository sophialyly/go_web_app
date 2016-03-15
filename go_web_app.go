package main

import (
	"net/http"
	"text/template"
	//"io/ioutil"
	//"bufio"
	//"os"
	//"strings"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Header().Add("Content Type","text/html")
		//tmpl, err := template.New("test").Parse(doc)
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		//if err == nil {
			context := Context{
				[3]string{"Lemon", "Orange", "Apple"},
				"the title",
			}
			//tmpl.Execute(w,context)
			templates.Lookup("test").Execute(w, context)
		//}

	})
	//http.HandleFunc("/", render)
	http.ListenAndServe(":8000", nil)
}

func render(w http.ResponseWriter, req *http.Request){
	w.Header().Add("Content Type","text/html")
	tmpl, err := template.New("test").Parse(doc)
	if err == nil {
		//context := Context{"the message"}
		//tmpl.Execute(w,context)
		tmpl.Execute(w,req.URL.Path)
	}
}

//type Context struct  {
//	Message string
//
//}

const doc = `
{{template "header" .Title}}
 <body>
  <h1>List of Fruits</h1>
  <ul>
    {{range .Fruit}}
      <li>{{.}}</li>
    {{end}}
  </ul>
 </body>
 {{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
 <head><title>{{.}}</title></head>
 `

const footer = `
</html>
`

type Context struct {
	Fruit [3] string
	Title string
}
