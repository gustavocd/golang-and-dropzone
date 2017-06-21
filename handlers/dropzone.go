package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gustavocd/golang-and-dropzone/models"
	"github.com/gustavocd/golang-and-dropzone/views"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		logrus.Errorf("Something went wrong %v", err)
		return
	}
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := views.Render(w, "index.html", nil)
	checkError(err)
}

// Store ...
func Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, header, err := r.FormFile("file")
	checkError(err)
	defer file.Close()
	mimeType := header.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		saveFile(w, file, header)
	case "image/png":
		saveFile(w, file, header)
	default:
		jsonResponse(w, http.StatusBadRequest, map[string]string{"message": "Por favor envie un archivo válido"})
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	path := "./uploads/" + handle.Filename

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	image := &models.File{
		ID:   bson.NewObjectId(),
		Name: handle.Filename,
		Path: path,
	}

	if err := image.Create(); err != nil {
		panic(err)
	}
	id := image.ID.Hex()
	jsonResponse(w, http.StatusCreated, map[string]string{"id": id})
}

func jsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", json)
}
