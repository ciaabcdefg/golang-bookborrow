package auth

import (
	"borrow/internal/prefixedrouter"
	"borrow/internal/writer"
	"borrow/repo"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func (h *Handler) RegisterRoutes(router prefixedrouter.CommonRouter) *Handler {
	router.POST("/auth/login/student", h.handleStudentLogin)
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

func (h *Handler) handleStudentLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body StudentLoginParams

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to parse json - %s", err.Error()), http.StatusBadRequest)
		return
	}

	studentID, err := strconv.Atoi(body.StudentID)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to parse studentID - %s", err.Error()), http.StatusBadRequest)
		return
	}

	student, err := h.repo.GetStudentByID(r.Context(), int32(studentID))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch student - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if student.Password != body.Password {
		writer.WriteError(w, fmt.Errorf("incorrect password"), http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      student,
		"user_type": "Student",
	})

	signedToken, err := token.SignedString([]byte("hello"))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to sign JWT - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    signedToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	writer.WriteJson(w, map[string]string{
		"acesss_token": signedToken,
	}, http.StatusOK)
}
