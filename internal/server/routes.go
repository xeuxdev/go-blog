package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

// func (s *Server) GetAllUsers() ([]models.User, error) {
// 	users := []models.User{}
// 	err := s.db.Select(&users, "SELECT id, email, name, password FROM users")
// 	return users, err
// }

// func CreateUser(user models.User) error {
// 	_, err := db.DB.NamedExec(`INSERT INTO users (id, email, name, password) VALUES (:id, :email, :name, :password)`, &user)
// 	return err
// }

// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := repositories.GetAllUsers()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(users)
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if err := repositories.CreateUser(user); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// func GetAllPosts() ([]models.Post, error) {
// 	posts := []models.Post{}
// 	err := db.DB.Select(&posts, "SELECT id, title, content, full_content, image, likesCount, viewCount, createdAt, updatedAt, user_id FROM posts")
// 	return posts, err
// }

// func CreatePost(post models.Post) error {
// 	_, err := db.DB.NamedExec(`INSERT INTO posts (id, title, content, full_content, image, likesCount, viewCount, createdAt, updatedAt, user_id) VALUES (:id, :title, :content, :full_content, :image, :likesCount, :viewCount, :createdAt, :updatedAt, :user_id)`, &post)
// 	return err
// }

// func GetAllPosts(w http.ResponseWriter, r *http.Request) {
// 	posts, err := repositories.GetAllPosts()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(posts)
// }

// func CreatePost(w http.ResponseWriter, r *http.Request) {
// 	var post models.Post
// 	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if err := repositories.CreatePost(post); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }
