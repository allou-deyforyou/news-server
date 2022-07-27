package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"news/handler"

	_ "github.com/mattn/go-sqlite3"
)

var server *handler.Handler

func init() {
	server = handler.NewHandler(handler.NewEntClient())
}

func init() {
	server.Handle("/news/article/content", handler.ParseHandler(server.ArticleContent))
	server.Handle("/news/article/post", handler.ParseHandler(server.ArticlePostList))

	server.Handle("/news/media/content", handler.ParseHandler(server.MediaContent))
	server.Handle("/news/media/post", handler.ParseHandler(server.MediaPostList))
}

func main() {
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server))
}
