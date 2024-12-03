create type "BookStatus" as ENUM ('Unavailable', 'Available', 'Borrowed');

create table if not exists "Book" (
  "ID" serial primary key,
  "Title" varchar(255) not null default 'Untitled',
  "Description" text not null default '',
  "Status" "BookStatus" not null default 'Unavailable',
  "PublishedAt" timestamp,
  "AddedAt" timestamp not null default CURRENT_TIMESTAMP
);