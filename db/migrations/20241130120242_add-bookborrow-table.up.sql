do $$ begin
  create type "BorrowStatus" as ENUM ('Borrowing', 'Returned');
exception
  when duplicate_object then null; 
end $$;

create table if not exists "BookBorrow" (
  "BookID" integer references "Book"("ID"),
  "BorrowListID" integer references "BorrowList"("ID"),
  "Status" "BorrowStatus" not null default 'Borrowing',
  "ReturnedAt" timestamp,
  primary key("BookID", "BorrowListID")
);