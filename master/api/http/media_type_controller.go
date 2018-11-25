package http

import "net/http"

func handleMediaTypeGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handleMediaTypeGetAll"))
}
