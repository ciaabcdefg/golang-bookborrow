package main

import (
	"borrow/internal/db"
	"borrow/internal/env"
	"borrow/internal/prefixedrouter"
	"borrow/repo"
	"borrow/services/auth"
	"borrow/services/books"
	"borrow/services/borrow"
	"borrow/services/students"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	router := httprouter.New()

	subrouter := prefixedrouter.New(
		"/api/v1", router, prefixedrouter.EmptyMiddleware,
	)

	// router.GET("/", Index)

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)

	books.NewHandler(repo, db).RegisterRoutes(subrouter)
	students.NewHandler(repo, db).RegisterRoutes(subrouter)
	borrow.NewHandler(repo, db).RegisterRoutes(subrouter)
	auth.NewHandler(repo, db).RegisterRoutes(subrouter)

	log.Println("Connected to the database")
	log.Printf("Listening on %s:%s\n", env.Hostname, env.Port)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
	})

	corsHandler := corsOptions.Handler(router)
	http.ListenAndServe(env.Hostname+":"+env.Port, corsHandler)
}
