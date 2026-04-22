package report

import (
	"mini-inventory/config"
	"time"
)

type ReportService struct{}

func (s *ReportService) GetReceiveReport(productID uint, startDate, endDate string) ([]InventoryReportRecord, error) {
	return s.getInventoryRecords("IN", productID, startDate, endDate)
}

func (s *ReportService) GetDeliveryReport(productID uint, startDate, endDate string) ([]InventoryReportRecord, error) {
	return s.getInventoryRecords("OUT", productID, startDate, endDate)
}

func (s *ReportService) getInventoryRecords(recordType string, productID uint, startDate, endDate string) ([]InventoryReportRecord, error) {
	startDate, endDate = s.resolveDates(startDate, endDate)

	query := `SELECT i.id, i.product_id, p.title as product_title, p.product_code, 
                     i.location_id, l.title as location_title, i.items, i.note, 
                     u.name as created_by, i.created_at
              FROM inventories i
              JOIN products p ON i.product_id = p.id
              JOIN locations l ON i.location_id = l.id
              JOIN users u ON i.created_by = u.id
              WHERE i.record_type = ? 
              AND i.created_at >= ? AND i.created_at <= ?`

	args := []interface{}{recordType, startDate, endDate}
	if productID > 0 {
		query += " AND i.product_id = ?"
		args = append(args, productID)
	}

	query = config.DB.Rebind(query)
	var records []InventoryReportRecord
	err := config.DB.Select(&records, query, args...)
	return records, err
}

func (s *ReportService) GetStockRegister(productID uint, startDate, endDate string) ([]StockRegisterItem, error) {
	startDate, endDate = s.resolveDates(startDate, endDate)

	query := `
		SELECT 
			p.id as product_id, p.product_code, p.title, p.uom,
			COALESCE(opening.total, 0) as opening,
			COALESCE(current_in.total, 0) as in_total,
			COALESCE(current_out.total, 0) as out_total,
			(COALESCE(opening.total, 0) + COALESCE(current_in.total, 0) - COALESCE(current_out.total, 0)) as balance
		FROM products p
		LEFT JOIN (
			SELECT product_id, SUM(CASE WHEN record_type = 'IN' THEN items ELSE -items END) as total
			FROM inventories WHERE created_at < ? GROUP BY product_id
		) opening ON p.id = opening.product_id
		LEFT JOIN (
			SELECT product_id, SUM(items) as total
			FROM inventories WHERE record_type = 'IN' AND created_at >= ? AND created_at <= ? GROUP BY product_id
		) current_in ON p.id = current_in.product_id
		LEFT JOIN (
			SELECT product_id, SUM(items) as total
			FROM inventories WHERE record_type = 'OUT' AND created_at >= ? AND created_at <= ? GROUP BY product_id
		) current_out ON p.id = current_out.product_id
	`
	
	args := []interface{}{startDate, startDate, endDate, startDate, endDate}
	if productID > 0 {
		query += " WHERE p.id = ?"
		args = append(args, productID)
	}

	query = config.DB.Rebind(query)
	var records []StockRegisterItem
	err := config.DB.Select(&records, query, args...)
	return records, err
}

func (s *ReportService) resolveDates(startDate, endDate string) (string, string) {
	if startDate == "" {
		startDate = time.Now().Format("2006-01-02")
	}
	if endDate == "" {
		endDate = startDate
	}

	if len(startDate) == 10 {
		startDate += " 00:00:00"
	}
	if len(endDate) == 10 {
		endDate += " 23:59:59"
	}

	return startDate, endDate
}
