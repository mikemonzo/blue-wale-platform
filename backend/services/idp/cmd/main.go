package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/config"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/persistence/impl/postgres"
)

var userRepo *postgres.PostgresUserRepository

func main() {
	cfg := config.NewConfig()

	// Configure DB connection
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Host, cfg.DB.Port, cfg.DB.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB connection: %v", err)
	}
	defer db.Close()

	// Run DB migrations
	if err := postgres.RunMigrations(db); err != nil {
		log.Fatalf("Error running DB migrations: %v", err)
	}

	// Initialize repositories
	userRepo = postgres.NewPostgresUserRepository(db).(*postgres.PostgresUserRepository)

	setupRoutes()

	log.Printf("Starting IDP service on :%d", cfg.Server.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/users", usersHandler)
	//http.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	//http.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	//http.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Implement GetAllUsers method in postgresUserRepository
	switch r.Method {
	case http.MethodGet:
		getAllUsersHandler(w, r)
	case http.MethodPost:
		createUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := userRepo.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	fmt.Println("Creating user: createUserHandler")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Creating user: %+v", user)
	if err := userRepo.Create(r.Context(), &user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
