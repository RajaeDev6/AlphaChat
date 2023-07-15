package user

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	text := []byte("login page")
	w.Write(text)
}
