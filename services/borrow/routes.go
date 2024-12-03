package borrow

import (
	"borrow/internal/array"
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
	// "borrow/repo"
)

func (h *Handler) RegisterRoutes(router prefixedrouter.CommonRouter) *Handler {
	router.POST("/borrow", h.handleCreateBorrowList)
	router.GET("/borrow-lists", h.handleGetMyBorrowLists)
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

func (h *Handler) handleCreateBorrowList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var body CreateBorrowListParams
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to parse body - %s", err.Error()), http.StatusBadRequest)
		return
	}

	// begin transaction
	tx, err := h.db.Begin()
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to start transaction - %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// all queries returned by WithTx will be inside the transaction
	qtx := h.repo.WithTx(tx)

	// create borrow list
	borrowListID, err := qtx.CreateBorrowList(context.Background(), body.StudentID)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to create borrow list - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// checking if all books are available
	//// all books must be available to be borrowed
	booksToBorrow, err := qtx.GetBooks(r.Context(), body.BookIDs)
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to create data of books to borrow - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	unavailableBooks := array.Filter(booksToBorrow, func(book repo.Book) bool {
		return book.Status != repo.BookStatusAvailable
	})

	if len(unavailableBooks) > 0 {
		writer.WriteJson(w, map[string]any{
			"error":             "some books are not available (do not have the 'Available' status)",
			"unavailable_books": unavailableBooks,
		}, http.StatusBadRequest)
		return
	}

	// create book borrow entries (belongs to the borrow list)
	bookIDs, err := qtx.AddBooksToBorrowList(context.Background(), repo.AddBooksToBorrowListParams{
		BookIds:      body.BookIDs,
		BorrowListID: borrowListID,
	})

	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to borrow books - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to commit transaction - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, map[string]any{
		"BookIDs": bookIDs,
	}, http.StatusCreated)
}

func (h *Handler) handleGetMyBorrowLists(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	query := r.URL.Query()

	var id int
	var err error

	{
		idString := query.Get("id")
		id, err = strconv.Atoi(idString)
		if err != nil {
			writer.WriteError(w, fmt.Errorf("failed to parse id - %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	borrowLists, err := h.repo.GetAllMyBorrowLists(context.Background(), int32(id))
	if err != nil {
		writer.WriteError(w, fmt.Errorf("failed to fetch borrow lists - %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteJson(w, array.Map(borrowLists, func(list repo.BorrowList) BorrowListReturn {
		return BorrowListReturn{
			ID:         list.ID,
			StudentID:  list.StudentID,
			BorrowedAt: list.BorrowedAt,
			Status:     string(list.Status.([]uint8)),
			CreatedAt:  list.CreatedAt,
		}
	}), http.StatusOK)
}
