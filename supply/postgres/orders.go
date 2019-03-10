package postgres

import (
	"database/sql"
	"field/supply"
	"field/supply/ordering"
	"log"
)

const (
	findOrder = iota
	findOrderItems
)

var orderSqlStmts = []string{
	"SELECT orderid,projectid,sentdate,status FROM orders WHERE orderid=$1",                                             // findOrder
	"SELECT oi.productid, oi.name FROM orders o INNER JOIN order_items oi ON o.orderid = oi.orderid WHERE o.orderid=$1", // findOrderItems
}

type OrderRepository struct {
	db            *sql.DB
	preparedStmts []*sql.Stmt
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	r := &OrderRepository{db: db, preparedStmts: make([]*sql.Stmt, 0, len(productSqlStmts))}
	r.createPreparedStmts()
	return r
}

func (r *OrderRepository) createPreparedStmts() {
	for _, stmt := range orderSqlStmts {
		ps, err := r.db.Prepare(stmt)
		if err != nil {
			r.db.Close()
			log.Fatalf("unable to prepare statement %q: %v", stmt, err)
		}
		r.preparedStmts = append(r.preparedStmts, ps)
	}
}

func (r *OrderRepository) Save(o *supply.Order) error {
	return nil
}

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	o := supply.Order{Items: make([]*supply.Item, 0)}

	tx, err := r.db.Begin()
	if err != nil {
		return &supply.Order{}, err
	}
	// defer tx.Rollback()

	err = tx.Stmt(r.preparedStmts[findOrder]).QueryRow(id).Scan(&o.OrderID, &o.ProjectID, &o.SentDate, &o.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return &supply.Order{}, err
		} else {
			return &supply.Order{}, err
		}
	}

	rows, err := tx.Stmt(r.preparedStmts[findOrderItems]).Query(id)
	if err != nil {
		return &supply.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var i supply.Item
		err := rows.Scan(&i.ProductID, &i.Name)
		if err != nil {
			return &supply.Order{}, err
		}
		o.Items = append(o.Items, &i)
	}
	err = rows.Err()
	if err != nil {
		return &supply.Order{}, err
	}

	err = tx.Commit()
	if err != nil {
		return &supply.Order{}, err
	}

	return &o, nil
}

func (r *OrderRepository) FindDates(projectid string) ([]ordering.ProjectOrder, error) {
	var orders []ordering.ProjectOrder

	return orders, nil
}
