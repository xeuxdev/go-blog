package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func HandleUsers(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /users")
		w.Write([]byte("Hello World!!!!"))
	})
}
