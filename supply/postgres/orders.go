package postgres

import (
	"database/sql"
	"field/supply"
	"field/supply/ordering"
)

type OrderRepository struct {
	db            *sql.DB
	preparedStmts []*sql.Stmt
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	r := &OrderRepository{db: db, preparedStmts: make([]*sql.Stmt, 0, len(productSqlStmts))}

	return r
}

const saveOrder = `
	INSERT INTO orders
		(orderid, projectid, sentdate, status)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT orders_pk
	DO UPDATE SET 
		projectid=EXCLUDED.projectid,
		sentdate=EXCLUDED.sentdate,
		status=EXCLUDED.status`

const saveItems = `
	INSERT INTO order_items
		(orderid, productid, name, uom, requested, received, remaining, status, ponumber)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (orderid, productid)
	DO UPDATE SET
		name=EXCLUDED.name,
		uom=EXCLUDED.uom,
		requested=EXCLUDED.requested,
		received=EXCLUDED.received,
		remaining=EXCLUDED.remaining,
		status=EXCLUDED.status,
		ponumber=EXCLUDED.ponumber`

func (r *OrderRepository) Save(o *supply.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(saveOrder, o.OrderID, o.ProjectID, o.SentDate, o.Status)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(saveItems)
	if err != nil {
		return err
	}
	for _, item := range o.Items {
		_, err = stmt.Exec(
			o.OrderID,
			item.ProductID,
			item.Name,
			item.UOM,
			item.QuantityRequested,
			item.QuantityReceived,
			item.QuantityRemaining,
			item.ItemStatus,
			item.PONumber)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

const findOrder = `
	SELECT orderid,projectid,sentdate,status 
	FROM orders 
	WHERE orderid=$1`

const findOrderItems = `
	SELECT oi.productid, oi.name, oi.uom, oi.requested, oi.received, oi.remaining, oi.status, oi.ponumber
	FROM orders o 
	INNER JOIN order_items oi 
	ON o.orderid = oi.orderid 
	WHERE o.orderid=$1`

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	o := supply.Order{Items: make([]*supply.Item, 0)}

	tx, err := r.db.Begin()
	if err != nil {
		return &supply.Order{}, err
	}

	err = tx.QueryRow(findOrder, id).Scan(&o.OrderID, &o.ProjectID, &o.SentDate, &o.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return &supply.Order{}, err
		} else {
			return &supply.Order{}, err
		}
	}

	rows, err := tx.Query(findOrderItems, id)
	if err != nil {
		return &supply.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var i supply.Item
		err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.UOM,
			&i.QuantityRequested,
			&i.QuantityReceived,
			&i.QuantityRemaining,
			&i.ItemStatus,
			&i.PONumber)
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

const findOrderDates = `
	SELECT orderid, sentdate, status
	FROM orders
	WHERE projectid = $1`

func (r *OrderRepository) FindDates(projectid string) ([]ordering.ProjectOrder, error) {
	orders := make([]ordering.ProjectOrder, 0)

	rows, err := r.db.Query(findOrderDates, projectid)
	if err != nil {
		return []ordering.ProjectOrder{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var o ordering.ProjectOrder
		err := rows.Scan(&o.OrderID, &o.SentDate, &o.Status)
		if err != nil {
			return []ordering.ProjectOrder{}, nil
		}
		orders = append(orders, o)
	}
	err = rows.Err()
	if err != nil {
		return []ordering.ProjectOrder{}, nil
	}

	return orders, nil
}
