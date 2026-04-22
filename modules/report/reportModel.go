package report

import "time"

type StockRegisterItem struct {
	ProductID   uint    `json:"product_id" db:"product_id"`
	ProductCode string  `json:"product_code" db:"product_code"`
	Title       string  `json:"title" db:"title"`
	Opening     int     `json:"opening" db:"opening"`
	InTotal     int     `json:"in_total" db:"in_total"`
	OutTotal    int     `json:"out_total" db:"out_total"`
	Balance     int     `json:"balance" db:"balance"`
	UOM         string  `json:"uom" db:"uom"`
}

type InventoryReportRecord struct {
	ID            uint      `json:"id" db:"id"`
	ProductID     uint      `json:"product_id" db:"product_id"`
	ProductTitle  string    `json:"product_title" db:"product_title"`
	ProductCode   string    `json:"product_code" db:"product_code"`
	LocationID    uint      `json:"location_id" db:"location_id"`
	LocationTitle string    `json:"location_title" db:"location_title"`
	Items         int       `json:"items" db:"items"`
	Note          string    `json:"note" db:"note"`
	CreatedBy     string    `json:"created_by" db:"created_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type ReportFilter struct {
	ProductID uint      `json:"product_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
