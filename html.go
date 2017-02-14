package main

import (
	"log"
    "time"
	"net/http"
	"text/template"

	"github.com/emicklei/go-restful"
)

// This example shows how to serve a HTML page using the standard Go template engine.
//
// GET http://localhost:8080/

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(home))
    ws.Route(ws.POST("/").To(power))
	restful.Add(ws)
	print("open browser on http://localhost:8080/\n")
	http.ListenAndServe(":8080", nil)
}

type Message struct {
	Text string
}

func home(req *restful.Request, resp *restful.Response) {
	p := &Message{"restful-html-template demo"}
	// you might want to cache compiled templates
	t, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, p)
}

func power(req *restful.Request, resp *restful.Response) {

    current := time.Now()
	p := &Message{"I'v done power on " + current.Format(time.RFC3339)}

    log.Println("Received ", current)

	// you might want to cache compiled templates
	t, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, p)
}
