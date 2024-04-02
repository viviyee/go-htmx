package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Song struct {
	Title  string
	Artist string
}

var data map[string][]Song = map[string][]Song{
	"songs": {
		{Title: "Positions", Artist: "Ariana Grande"},
		{Title: "Blank Space", Artist: "Taylor Swift"},
		{Title: "Sunflower", Artist: "Post Malone"},
	},
}

func main() {
	// Hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		t.Execute(w, data)
	})

	http.HandleFunc("/songs", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(1 * time.Second)
		log.Print("htmx request was received")
		log.Print(req.Header.Get("HX-Request"))

		title := req.PostFormValue("title")
		artist := req.PostFormValue("artist")

		// li := fmt.Sprintf("<li class='list-group-item'><p>%s - %s</p></li>", title, artist)
		// t, _ := template.New("t").Parse(li)
		// t.Execute(w, nil)

		t := template.Must(template.ParseFiles("templates/index.html"))
		t.ExecuteTemplate(w, "film-list-element", Song{Title: title, Artist: artist})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
