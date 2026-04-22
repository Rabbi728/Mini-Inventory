package product

import (
	"mini-inventory/config"
	"time"
)

type ProductService struct{}

func (s *ProductService) GetAllProducts() ([]Product, error) {
	var products []Product
	query := `SELECT id, title, description, color, size, uom, product_code, created_at, updated_at FROM products`
	err := config.DB.Select(&products, query)
	return products, err
}

func (s *ProductService) GetProductByID(id string) (Product, error) {
	var p Product
	query := `SELECT id, title, description, color, size, uom, product_code, created_at, updated_at FROM products WHERE id = $1`
	err := config.DB.Get(&p, query, id)
	return p, err
}

func (s *ProductService) CreateProduct(p *Product) error {
	query := `INSERT INTO products (title, description, color, size, uom, product_code, created_at, updated_at) 
			  VALUES (:title, :description, :color, :size, :uom, :product_code, :created_at, :updated_at) RETURNING id`

	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now

	nstmt, err := config.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Get(&p.ID, p)
	return err
}

func (s *ProductService) UpdateProduct(p *Product) error {
	query := `UPDATE products SET title = :title, description = :description, color = :color, 
			  size = :size, uom = :uom, product_code = :product_code, updated_at = :updated_at WHERE id = :id`

	p.UpdatedAt = time.Now()
	_, err := config.DB.NamedExec(query, p)
	return err
}

func (s *ProductService) DeleteProduct(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}
