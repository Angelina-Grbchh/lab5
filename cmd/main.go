package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"user-crud-api/internal/handler"
	"user-crud-api/internal/repository"
	"user-crud-api/internal/service"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/users?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.ListUsers)
	r.Get("/users/{id}", userHandler.GetUser)
	r.Put("/users/{id}", userHandler.UpdateUser)
	r.Delete("/users/{id}", userHandler.DeleteUser)

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
