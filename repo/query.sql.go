// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createBook = `-- name: CreateBook :one
insert into "Book" ("Title", "Description", "Status", "PublishedAt") 
values ($1, $2, $3, $4) returning "ID", "AddedAt"
`

type CreateBookParams struct {
	Title       string
	Description string
	Status      BookStatus
	PublishedAt sql.NullTime
}

type CreateBookRow struct {
	ID      int32
	AddedAt time.Time
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (CreateBookRow, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.PublishedAt,
	)
	var i CreateBookRow
	err := row.Scan(&i.ID, &i.AddedAt)
	return i, err
}

const deleteBookByID = `-- name: DeleteBookByID :execresult
delete from "Book" where "ID" = $1
`

func (q *Queries) DeleteBookByID(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteBookByID, id)
}

const getAllBooks = `-- name: GetAllBooks :many
select "ID", "Title", "Description", "Status", "PublishedAt", "AddedAt" from "Book"
`

func (q *Queries) GetAllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.PublishedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllStudents = `-- name: GetAllStudents :many
select "ID", "FirstName", "MiddleName", "LastName", "Password", "Gender" from "Student"
`

func (q *Queries) GetAllStudents(ctx context.Context) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getAllStudents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Password,
			&i.Gender,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookByID = `-- name: GetBookByID :one
select "ID", "Title", "Description", "Status", "PublishedAt", "AddedAt" from "Book" where "ID" = $1
`

func (q *Queries) GetBookByID(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByID, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.PublishedAt,
		&i.AddedAt,
	)
	return i, err
}

const getBooks = `-- name: GetBooks :many
select "ID", "Title", "Description", "Status", "PublishedAt", "AddedAt" from "Book"
where "ID" = ANY($1::int[])
`

func (q *Queries) GetBooks(ctx context.Context, bookIds []int32) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getBooks, pq.Array(bookIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.PublishedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentByID = `-- name: GetStudentByID :one
select "ID", "FirstName", "MiddleName", "LastName", "Password", "Gender" from "Student" where "ID" = $1
`

func (q *Queries) GetStudentByID(ctx context.Context, id int32) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByID, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Password,
		&i.Gender,
	)
	return i, err
}

const getStudents = `-- name: GetStudents :many
select "ID", "FirstName", "MiddleName", "LastName", "Password", "Gender" from "Student" limit $1 offset $2
`

type GetStudentsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetStudents(ctx context.Context, arg GetStudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getStudents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Password,
			&i.Gender,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalStudents = `-- name: GetTotalStudents :one
select count("ID") from "Student"
`

func (q *Queries) GetTotalStudents(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalStudents)
	var count int64
	err := row.Scan(&count)
	return count, err
}
