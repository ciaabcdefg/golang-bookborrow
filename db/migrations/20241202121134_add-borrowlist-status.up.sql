do $$ begin
  create type "BorrowListStatus" as ENUM ('Draft', 'Borrowing', 'Returned');
exception
  when duplicate_object then null; 
end $$;

alter table if exists "BorrowList"
add "Status" "BorrowListStatus" not null default 'Draft';