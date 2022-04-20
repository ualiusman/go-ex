package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/freshman-tech/news-demo-starter-files/news"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page ==""{
			page = "1"
		}

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v", results)

		fmt.Println("Search query is: ", searchQuery)
		fmt.Println("Page is: ", page)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w,nil)
}



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == ""{
		log.Fatal(" Env: api key must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}

	fs := http.FileServer(http.Dir("assets"))
	newsapi := news.NewClient(myClient,apiKey,20)

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
