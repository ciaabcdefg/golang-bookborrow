-- name: GetAllBooks :many
select * from "Book";

-- name: GetBooks :many
select * from "Book"
where "ID" = ANY(@book_ids::int[]);

-- name: GetBookByID :one
select * from "Book" where "ID" = $1;

-- name: CreateBook :one
insert into "Book" ("Title", "Description", "Status", "PublishedAt") 
values ($1, $2, $3, $4) returning "ID", "AddedAt"; 

-- name: DeleteBookByID :execresult
delete from "Book" where "ID" = $1;

-- name: GetStudentByID :one
select * from "Student" where "ID" = $1;

-- name: GetAllStudents :many
select * from "Student";

-- name: GetStudents :many
select * from "Student" limit $1 offset $2;

-- name: GetTotalStudents :one
select count("ID") from "Student";