package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"news/handler"
	"news/internal"

	_ "github.com/mattn/go-sqlite3"
)

var server *handler.Handler

func init() {
	server = handler.NewHandler(internal.NewEntClient())
}

func init() {
	server.Handle("/news/category", handler.ParseHandler(server.NewsCategoryPost))
	server.Handle("/news/latest", handler.ParseHandler(server.NewsLatestPost))
	server.Handle("/news/article", handler.ParseHandler(server.NewsArticle))

	server.Handle("/tv/post", handler.ParseHandler(server.TvPost))
}

func main() {
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server))
}
