package main

import (
	"net/http"

	"github.com/otiai10/gosseract"
)

// Index do home handler
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

// OcrHandler do home handler
func OcrHandler(w http.ResponseWriter, r *http.Request) {
	client := gosseract.NewClient()
	defer client.Close()
	// client.SetLanguage("eng+tha")
	client.SetImage("./assets/TestNaja.png")
	text, _ := client.Text()

	w.Header().Set("Content-Type", "text/html; charset=utf-8;")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}
