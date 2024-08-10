package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xeuxdev/go-blog/internal/database"
	"github.com/xeuxdev/go-blog/internal/handlers"
)

func main() {

	db := database.ConnectDB()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /")
		w.Write([]byte("Hello World!!!!"))
	})

	service := &handlers.Service{DB: db}

	service.HandleUsers(router)

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
