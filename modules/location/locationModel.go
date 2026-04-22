package location

import (
	"time"
)

type Location struct {
	ID          uint      `db:"id" json:"id"`
	Title       string    `db:"title" json:"title" binding:"required"`
	Description string    `db:"description" json:"description"`
	CreatedBy   uint      `db:"created_by" json:"created_by"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}