create table if not exists "BookAuthor" (
  "BookID" integer references "Book"("ID"),
  "AuthorID" integer references "Author"("ID"),
  primary key("BookID", "AuthorID")
)