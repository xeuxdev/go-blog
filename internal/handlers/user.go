package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/xeuxdev/go-blog/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data"`
}

type Service struct {
	DB *sql.DB
}

func (s *Service) HandleUsers(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/create", s.createUserHandler)
	})
}

func (s *Service) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if r.ContentLength == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		fmt.Println("JSON decoding error:", err)
		return
	}

	// Generate a unique ID for the user
	user.ID = strings.Join(strings.Split(uuid.New().String(), "-"), "")

	// Register the user
	response := s.registerUser(user)

	// Set the appropriate status code and send the response
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

// func HandleUsers(r chi.Router) {

// r.Get("/", func(w http.ResponseWriter, r *http.Request) {

// 	var users []User
// 	var db = database.DB

// 	rows, err := db.Query("SELECT id, email, name, password FROM user")
// 	if err != nil {
// 		http.Error(w, "Failed to query the database", http.StatusInternalServerError)
// 		fmt.Println("Query error:", err)
// 		return
// 	}
// 	defer rows.Close() // Ensure rows are closed after function exits

// 	// Iterate through the result set
// 	for rows.Next() {
// 		var user User
// 		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
// 			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
// 			fmt.Println("Scan error:", err)
// 			return
// 		}
// 		users = append(users, user)
// 	}

// 	// Check for errors after iteration
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, "Error encountered during iteration", http.StatusInternalServerError)
// 		fmt.Println("Row iteration error:", err)
// 		return
// 	}

// 	// Set the content type and encode the response
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(users); err != nil {
// 		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		fmt.Println("JSON encoding error:", err)
// 	}
// })

// r.Get("/{email}", func(w http.ResponseWriter, r *http.Request) {

// 	email := chi.URLParam(r, "email")

// 	var user User

// 	if email == "" {
// 		http.Error(w, "email is required", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := getUserByEmail(email)

// 	if err != nil {
// 		if err.Error() == fmt.Sprintf("getUserByEmail %s: no such user", email) {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)

// })

// r.Post("/create", func(w http.ResponseWriter, r *http.Request) {

// 	var user models.User

// 	if r.ContentLength == 0 {
// 		http.Error(w, "Request body is empty", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
// 		fmt.Println("JSON decoding error:", err)
// 		return
// 	}

// 	user.ID = strings.Join(strings.Split(uuid.New().String(), "-"), "")

// 	response := registerUser(user)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(response)
// })

// 	r.Post("/create", )

// }

/*
Register a new user
*/
func (s *Service) registerUser(user models.User) Response {
	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return Response{
			Message: "Failed to hash password",
			Status:  http.StatusInternalServerError,
		}
	}

	// Insert the new user into the database
	_, err = s.DB.Exec("INSERT INTO user (id, email, name, password) VALUES (?, ?, ?, ?)", user.ID, user.Email, user.Name, hashedPassword)
	if err != nil {
		// Handle duplicate email error
		if strings.Contains(err.Error(), "Duplicate entry") {
			return Response{
				Message: "Email already exists",
				Status:  http.StatusConflict,
			}
		}

		// Log the error and return a general failure response
		log.Println("Database error:", err.Error())
		return Response{
			Message: "Failed to register user",
			Status:  http.StatusInternalServerError,
		}
	}

	// Return a successful response
	return Response{
		Message: "User created successfully",
		Status:  http.StatusCreated,
	}
}

// func getUserByEmail(email string) (User, error) {
// 	var user User
// 	var db = database.DB

// 	row := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

// 	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, fmt.Errorf("getUserByEmail %s: no such user", email)
// 		}
// 		return user, fmt.Errorf("getAlbumById %s: %v", email, err)
// 	}

// 	return user, nil
// }
