create table if not exists "BorrowList" (
  "ID" serial primary key,
  "StudentID" integer not null references "Student"("ID"),
  "BorrowedAt" timestamp not null default CURRENT_TIMESTAMP
);