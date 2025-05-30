package model

type User struct {
	ID       int64  `json:"id" db:"id"`
	FullName string `json:"full_name" db:"name"`
	Email    string `json:"email" db:"email"`
	Age      int32  `json:"age" db:"age"`
}
