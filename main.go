package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"go_shortener/config"
	"go_shortener/controller"
	"go_shortener/handler"
	"go_shortener/models"
)

func main() {
	db, err := handler.Open()
	if err != nil {
		fmt.Println("Unable to connect to mongodb", err)
		return
	}
	connection := db.C(config.DB_URL_COLLECTION)
	fmt.Println("Server listening on port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case "GET":
			shortUrl:=r.URL.Query().Get("short")
			longUrl:=r.URL.Query().Get("long")
			fmt.Println(shortUrl)
			fmt.Println(longUrl)
			urlInfo:=models.UrlInfo{
				ShortUrl:shortUrl,
				LongUrl:longUrl,
			}
			response, statusCode := controller.Fetch(connection,urlInfo)
			w.WriteHeader(statusCode)
			data, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "%s", data)
			break
		case "POST":
			var url string = r.PostFormValue("url")
			response, statusCode := controller.CreateURL(connection, url)
			w.WriteHeader(statusCode)
			data, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "%s", data)
			break
		case "PUT":

			break
		case "DELETE":
			break
		default:

			break
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
