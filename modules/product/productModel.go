package product

import (
	"time"
)

type Product struct {
	ID          uint      `db:"id" json:"id"`
	Title       string    `db:"title" json:"title" binding:"required"`
	Description string    `db:"description" json:"description"`
	Color       string    `db:"color" json:"color"`
	Size        string    `db:"size" json:"size"`
	Uom			string	  `db:"uom" json:"uom" binding:"required"`	
	ProductCode string    `db:"product_code" json:"product_code" binding:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
