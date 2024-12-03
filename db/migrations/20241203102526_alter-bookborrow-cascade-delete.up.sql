alter table "BookBorrow"
drop constraint "BookBorrow_BookID_fkey";

alter table "BookBorrow"
add constraint "BookBorrow_BookID_fkey"
foreign key ("BookID")
references "Book" ("ID")
on delete cascade;

alter table "BookBorrow"
drop constraint "BookBorrow_BorrowListID_fkey";

alter table "BookBorrow"
add constraint "BookBorrow_BorrowListID_fkey"
foreign key ("BorrowListID")
references "BorrowList" ("ID")
on delete cascade;