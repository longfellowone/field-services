package postgres

import (
	"database/sql"
	"field/supply"
)

type ProductRepository struct {
	db    *sql.DB
	cache []supply.Product
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Save(p *supply.Product) error {
	return nil
}

func (r *ProductRepository) Find(id string) (*supply.Product, error) {
	var product supply.Product

	return &product, nil
}

func (r *ProductRepository) FindAll() ([]supply.Product, error) {
	var products []supply.Product

	return products, nil
}
