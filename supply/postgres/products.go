package postgres

import (
	"database/sql"
	"github.com/longfellowone/field-services/supply"
	"log"
)

const (
	findProducts = iota
)

var productSqlStmts = []string{
	"SELECT productid, category, name, uom FROM products", // findProducts
}

type ProductRepository struct {
	db            *sql.DB
	preparedStmts []*sql.Stmt
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	r := &ProductRepository{db: db, preparedStmts: make([]*sql.Stmt, 0, len(productSqlStmts))}
	r.createPreparedStmts()
	return r
}

func (r *ProductRepository) createPreparedStmts() {
	for _, stmt := range productSqlStmts {
		ps, err := r.db.Prepare(stmt)
		if err != nil {
			r.db.Close()
			log.Fatalf("unable to prepare statement %q: %v", stmt, err)
		}
		r.preparedStmts = append(r.preparedStmts, ps)
	}
}

func (r *ProductRepository) Save(p *supply.Product) error {
	return nil
}

func (r *ProductRepository) Find(id string) (*supply.Product, error) {
	var product supply.Product
	return &product, nil
}

func (r *ProductRepository) FindAll() ([]*supply.Product, error) {
	rows, err := r.preparedStmts[findProducts].Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var products []*supply.Product
	for rows.Next() {
		var p supply.Product
		err := rows.Scan(&p.ID, &p.Category, &p.Name, &p.UOM)
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
