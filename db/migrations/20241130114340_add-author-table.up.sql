create table if not exists "Author" (
  "ID" serial primary key,
  "PenName" varchar(255) not null default 'Pen Name',
  "FirstName" varchar(255),
  "MiddleName" varchar(255),
  "LastName" varchar(255),
  "Gender" char
)