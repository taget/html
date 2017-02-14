package main

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
)

var decoder *schema.Decoder

func main() {
	decoder = schema.NewDecoder()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(home))
	ws.Route(ws.POST("/").Consumes("application/x-www-form-urlencoded").To(power))
	restful.Add(ws)
	print("open browser on http://localhost:8080/\n")
	http.ListenAndServe(":8080", nil)
}

// Message passint go html template
type Message struct {
	Text string
}

// PowerState from the html form
type PowerState struct {
	Power string
}

func home(req *restful.Request, resp *restful.Response) {
	p := &Message{"Click the buttons"}
	// you might want to cache compiled templates
	t, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, p)
}

func power(req *restful.Request, resp *restful.Response) {

	err := req.Request.ParseForm()
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	po := new(PowerState)
	err = decoder.Decode(po, req.Request.PostForm)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("got <%s> ", po.Power)

	current := time.Now()
	m := &Message{"I'v done power " + po.Power + " at " + current.Format(time.RFC3339)}
	// you might want to cache compiled templates
	t, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, m)
}
