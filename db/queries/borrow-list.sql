-- name: GetAllBorrowLists :many
select * from "BorrowList";

-- name: CreateBorrowList :one
insert into "BorrowList" ("StudentID") values ($1) returning "ID";

-- name: AddBooksToBorrowList :many
insert into "BookBorrow" ("BookID", "BorrowListID")
values (unnest(@book_ids::int[]), @borrow_list_id::int)
returning "BookID";

-- name: GetAllMyBorrowLists :many
select * from "BorrowList"
where "StudentID" = $1;