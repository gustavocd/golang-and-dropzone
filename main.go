package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gustavocd/golang-and-dropzone/handlers"
	"github.com/gustavocd/golang-and-dropzone/middlewares"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func staticFileServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fileServer := http.FileServer(http.Dir("public"))
	r.URL.Path = ps.ByName("filepath")
	f, err := os.Stat("./public/" + r.URL.Path)
	if err != nil || f.IsDir() {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	fileServer.ServeHTTP(w, r)
}

func main() {
	r := httprouter.New()

	r.GET("/", handlers.Index)
	r.POST("/upload", middlewares.IsMethodPost(handlers.Store))
	r.GET("/public/*filepath", staticFileServer)

	logrus.Println("Server running at localhost:8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
