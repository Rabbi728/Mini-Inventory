package inventory

import (
	"mini-inventory/config"
	"time"
)

type InventoryService struct{}

func (s *InventoryService) CreateInventory(inv *Inventory) error {
	query := `INSERT INTO inventories (product_id, location_id, record_type, items, note, created_by, created_at, updated_at)
			  VALUES (:product_id, :location_id, :record_type, :items, :note, :created_by, :created_at, :updated_at) RETURNING id`

	now := time.Now()
	inv.CreatedAt = now
	inv.UpdatedAt = now

	nstmt, err := config.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Get(&inv.ID, inv)
	return err
}
