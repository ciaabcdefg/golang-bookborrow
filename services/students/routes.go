package students

import (
	"borrow/internal/prefixedrouter"
	"borrow/internal/writer"
	"borrow/repo"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) RegisterRoutes(router prefixedrouter.CommonRouter) *Handler {
	router.GET("/student/id/:id", h.handleGetStudentByID)
	router.GET("/student/get", h.handleGetStudents)
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

func (h *Handler) handleGetStudentByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to parse id - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	student, err := h.repo.GetStudentByID(context.Background(), int32(id))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to get student - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, student, http.StatusOK)
}

func (h *Handler) handleGetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()

	limitStr := query.Get("limit")
	limit := 2147483647

	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			writer.WriteError(w, fmt.Errorf("failed to parse limit - %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	offsetStr := query.Get("offset")
	offset := 0

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			writer.WriteError(w, fmt.Errorf("failed to parse offset - %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	students, err := h.repo.GetStudents(context.Background(), repo.GetStudentsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch students - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	total, err := h.repo.GetTotalStudents(context.Background())
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch student count - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%d %d\n", limit, offset)

	writer.WriteJson(w, map[string]any{
		"data":  students,
		"total": int64(total),
	}, http.StatusOK)
}
