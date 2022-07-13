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
	server.Handle("/news/category/article/post", handler.ParseHandler(server.NewsCategoryArticlePost))
	server.Handle("/news/latest/article/post", handler.ParseHandler(server.NewsLatestArticlePost))
	server.Handle("/news/article", handler.ParseHandler(server.NewsArticle))

	server.Handle("/news/tv/post", handler.ParseHandler(server.NewsTvPost))
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
