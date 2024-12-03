alter table if exists "BorrowList"
drop column if exists "Status";

drop type if exists "BorrowListStatus";