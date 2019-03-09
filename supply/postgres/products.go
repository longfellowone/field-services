package postgres

import (
	"database/sql"
	"field/supply"
	"fmt"
	"log"
)

const (
	findProducts = iota
)

var sqlStmts = []string{
	"SELECT productid, category, name, uom FROM products", // findProducts
}

type ProductRepository struct {
	db            *sql.DB
	preparedStmts []*sql.Stmt
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	r := &ProductRepository{
		db:            db,
		preparedStmts: make([]*sql.Stmt, 0, len(sqlStmts)),
	}

	if err := r.createPreparedStmts(); err != nil {
		db.Close()
		log.Fatal(err)
	}
	return r
}

func (r *ProductRepository) createPreparedStmts() error {
	r.preparedStmts = []*sql.Stmt{}
	for _, stmt := range sqlStmts {
		ps, err := r.db.Prepare(stmt)
		if err != nil {
			return fmt.Errorf("unable to prepare statement %q: %v", stmt, err)
		}
		r.preparedStmts = append(r.preparedStmts, ps)
	}
	return nil
}

func (r *ProductRepository) Save(p *supply.Product) error {
	return nil
}

func (r *ProductRepository) Find(id string) (*supply.Product, error) {
	var product supply.Product

	return &product, nil
}

func (r *ProductRepository) FindAll() ([]*supply.Product, error) {
	rows, err := r.db.Query(sqlStmts[findProducts])
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var products []*supply.Product
	for rows.Next() {
		var p supply.Product
		err := rows.Scan(&p.ProductID, &p.Category, &p.Name, &p.UOM)
		if err != nil {
			log.Println(err)
		}
		products = append(products, &p)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return products, nil
}
