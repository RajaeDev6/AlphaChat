package routes

import (
	"fmt"
	"net/http"

	"github.com/RajaeDev6/AlphaChat/server/user"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	str := []byte("hello world")
	w.Write(str)
}

func Configure() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", user.LoginHandler)
	fmt.Println("server running")
	http.ListenAndServe(":8080", nil)
}
