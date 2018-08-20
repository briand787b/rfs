package http

import "net/http"

func HandleReceiveMedia(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "couldnt parse form", 400)
		return
	}

	var holder string
	for _, media := range r.MultipartForm.Value["media"] {
		holder += media
	}

	w.Write([]byte(holder))
}
