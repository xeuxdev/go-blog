package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xeuxdev/go-blog/internal/database"
)

func main() {

	database.ConnectDB()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /")
		w.Write([]byte("Hello World!!!!"))
	})

	// router.Route("/users", handlers.HandleUsers)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}

}
