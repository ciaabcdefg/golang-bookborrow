package books

import (
	"borrow/internal/prefixedrouter"
	"borrow/internal/writer"
	"borrow/repo"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) RegisterRoutes(router prefixedrouter.CommonRouter) *Handler {
	router.GET("/book", h.handleGetAllBooks)
	router.GET("/book/:id", h.handleGetBookByID)
	router.POST("/book", h.handleCreateBook)
	router.DELETE("/book/:id", h.handleDeleteBookByID)
	return h
}

type Handler struct {
	repo *repo.Queries
	db   *sql.DB
}

func NewHandler(repo *repo.Queries, db *sql.DB) *Handler {
	return &Handler{repo: repo, db: db}
}

// Handlers

func (h *Handler) handleGetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := h.repo.GetAllBooks(context.Background())

	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, books, http.StatusOK)
}

func (h *Handler) handleGetBookByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("invalid parameter - %s", err.Error()), http.StatusBadRequest)
		return
	}

	book, err := h.repo.GetBookByID(context.Background(), int32(id))

	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch book - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, book, http.StatusOK)
}

func (h *Handler) handleCreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := repo.CreateBookParams{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("invalid json - %s", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := h.repo.CreateBook(context.Background(), body)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to create - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, result, http.StatusCreated)
}

func (h *Handler) handleDeleteBookByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("invalid parameter - %s", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := h.repo.DeleteBookByID(context.Background(), int32(id))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to delete book - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	writer.WriteOk(w, fmt.Sprintf("deleted %d row(s)", rowsAffected), http.StatusOK)
}
