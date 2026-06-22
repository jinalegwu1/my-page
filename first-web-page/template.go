package main

import "html/template"
//var Tmpl = template.Must(template.ParseFiles("templates/index.html"))
var Tmpl = template.Must(
    template.ParseFiles("templates/index.html"),
)
