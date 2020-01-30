package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
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

func getFilteredQuotes(filter []string) ([]*Quote, error) {
	resp, err := http.Get("https://programming-quotes-api.herokuapp.com/quotes/lang/en")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var quotes []*Quote
	if jsonErr := json.NewDecoder(resp.Body).Decode(&quotes); jsonErr != nil {
		return nil, jsonErr
	}
	res := make([]*Quote, 0, len(quotes))
	// append only those the contain all strings from the filter expression
	for _, quote := range quotes {
		// transform author to lower
		text := strings.ToLower(quote.Author)
		// now test if all are contained
		containsAll := true
		for _, filterEntry := range filter {
			if !strings.Contains(text, filterEntry) {
				containsAll = false
				break
			}
		}
		if containsAll {
			res = append(res, quote)
		}
	}
	// now sort according to raiting
	sort.Slice(res, func(i, j int) bool {
		return res[i].Rating > res[j].Rating
	})
	return res, nil
}

type QuoteFilterHandler struct {
	Template *template.Template
	FilterBy []string
}

func NewQuoteFilterHandler(filter []string) QuoteFilterHandler {
	t := template.Must(template.ParseFiles("./templates/filter.html"))
	return QuoteFilterHandler{t, filter}
}

func (h QuoteFilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// usually handle the error, see https://stackoverflow.com/questions/30821745/idiomatic-way-to-handle-template-errors-in-golang
	quotes, quotesErr := getFilteredQuotes(h.FilterBy)
	if quotesErr != nil {
		http.Error(w, "Error getting quote", 400)
		log.Println("Error while getting quote:", quotesErr.Error())
		return
	}
	h.Template.Execute(w, quotes)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/random/", handleRandomQuote)
	randomTemplate := NewRandomTemplateHandler()
	http.Handle("/random-template/", randomTemplate)
	filterHandler := NewQuoteFilterHandler([]string{"alan", "turing"})
	http.Handle("/filter/", filterHandler)
	addr := "localhost:8080"
	log.Printf("Running server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
