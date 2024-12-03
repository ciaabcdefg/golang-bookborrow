create table if not exists "Student" (
  "ID" serial primary key,
  "FirstName" varchar(255) not null default 'First Name',
  "MiddleName" varchar(255),
  "LastName" varchar(255) not null default 'Last Name',
  "Password" varchar(255) not null,
  "Gender" char
)