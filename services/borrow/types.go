package borrow

import "time"

type CreateBorrowListParams struct {
	StudentID int32
	BookIDs   []int32
}

type BorrowListReturn struct {
	ID         int32
	StudentID  int32
	BorrowedAt time.Time
	Status     string
	CreatedAt  time.Time
}
