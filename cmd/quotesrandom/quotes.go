package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Quote struct {
	Text string `json:"en"`
	Author string `json:"author"`
	Rating float64 `json:"rating"`
}

func getRandomQuote() (*Quote, error) {
	resp, err := http.Get("https://programming-quotes-api.herokuapp.com/quotes/random/lang/en")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var quote Quote
	if jsonErr := json.NewDecoder(resp.Body).Decode(&quote); jsonErr != nil {
		return nil, jsonErr
	}
	return &quote, nil
}

func handleRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, quoteErr := getRandomQuote()
	if quoteErr != nil {
		http.Error(w, "Error getting quote", 400)
		log.Println("Error while getting quote:", quoteErr.Error())
		return
	}
	fmt.Fprintf(w, "%s, %s", quote.Author, quote.Text)
}

type RandomTemplateHandler struct {
	Template *template.Template
}

func NewRandomTemplateHandler() RandomTemplateHandler {
	t := template.Must(template.ParseFiles("./templates/random.html"))
	return RandomTemplateHandler{t}
}

func (h RandomTemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// usually handle the error, see https://stackoverflow.com/questions/30821745/idiomatic-way-to-handle-template-errors-in-golang
	quote, quoteErr := getRandomQuote()
	if quoteErr != nil {
		http.Error(w, "Error getting quote", 400)
		log.Println("Error while getting quote:", quoteErr.Error())
		return
	}
	h.Template.Execute(w, quote)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/random/", handleRandomQuote)
	randomTemplate := NewRandomTemplateHandler()
	http.Handle("/random-template/", randomTemplate)
	addr := ":8080"
	log.Printf("Running server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
