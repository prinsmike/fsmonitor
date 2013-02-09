package main

import (
	"github.com/DisposaBoy/JsonConfigReader"
	"html/template"
	"log"
	"net/http"
)

func main() {

}

func loadConfig() interface{} {
	// Read the config file.
	var v interface{}
	f, _ := os.Open("~/.fsmonitor.rc")
	// wrap our reader before passing it to the json decoder
	r := JsonConfigReader.New(f)
	json.NewDecoder(r).Decode(&v)
	return v
}
