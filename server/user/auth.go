package user

import (
	"encoding/json"
	"net/http"

	"github.com/RajaeDev6/AlphaChat/server/database"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	dbHandler := NewDbHandler(db)

	var user User
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)

	err := dbHandler.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User Create"))

}
