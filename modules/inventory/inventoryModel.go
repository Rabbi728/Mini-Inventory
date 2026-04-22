package inventory

import "time"

type Inventory struct {
	ID         uint       `db:"id" json:"id"`
	ProductID  uint       `db:"product_id" json:"product_id" binding:"required"`
	LocationID uint       `db:"location_id" json:"location_id" binding:"required"`
	RecordType string 	  `db:"record_type" json:"record_type"`
	Items      int        `db:"items" json:"items" binding:"required"`
	Note       string     `db:"note" json:"note"`
	CreatedBy  uint       `db:"created_by" json:"created_by"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
}
