package user

import (
	"time"
)

type User struct {
	ID        uint      `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	Email     string    `db:"email" json:"email" binding:"required,email"`
	Password  string    `db:"password" json:"password" binding:"required,min=6"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}