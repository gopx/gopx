package handler

import (
	"net/http"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
