package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8000", "http service address")
var templ = template.Must(template.New("hello").Parse(templateStr))

func main() {
	http.Handle("/", http.HandlerFunc(index))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("name"))
}

const templateStr = `
<html>
<head>
<title>Hello World!</title>
<head>
<body>
{{if .}}
Hello {{.}}
{{end}}
<form action="/" name=hello method="GET">
<input name=name value="" title="Enter your name">
<input type=submit value="Submit" name=name>
</form>
</body>
</html>
`
