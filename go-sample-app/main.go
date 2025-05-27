package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image/", imageHandler)
	
	// Health endpointni ham main ichida ro'yxatdan o'tkazamiz
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

type Index struct {
	Title string
	Body  string
	Links []Link
}

type Link struct {
	URL, Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Image gallery",
		Body:  "Welcome to the image gallery.",
	}
	for name, img := range images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("image.tmpl"))

type Image struct {
	Title string
	URL   string
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := images[strings.TrimPrefix(r.URL.Path, "/image/")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

var images = map[string]*Image{
	"go":     {"The Go Gopher", "https://golang.org/doc/gopher/frontpage.png"},
	"google": {"The Google Logo", "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"},
}
