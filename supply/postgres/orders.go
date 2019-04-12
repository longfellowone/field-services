package postgres

import (
	"database/sql"
	"field/supply"
	"field/supply/ordering"
	"github.com/pkg/errors"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

const saveOrder = `
	INSERT INTO orders
		(orderid, projectid, project_name, foreman_email, sentdate, status, comments)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT ON CONSTRAINT orders_pk
	DO UPDATE SET 
		projectid=EXCLUDED.projectid,
		project_name=EXCLUDED.project_name,
	    foreman_email=EXCLUDED.foreman_email,
		sentdate=EXCLUDED.sentdate,
		status=EXCLUDED.status,
		comments=EXCLUDED.comments`

const saveItems = `
	INSERT INTO order_items
		(orderid, productid, name, uom, requested, received, remaining, status, ponumber, dateadded)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT ON CONSTRAINT order_items_orderid_productid_unique
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

	_, err = tx.Exec(saveOrder, o.ID, o.Project.ID, o.Name, o.ForemanEmail, o.SentDate, o.Status, o.Comments)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(saveItems)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range o.Items {
		_, err = stmt.Exec(
			o.ID,
			item.ID,
			item.Name,
			item.UOM,
			item.QuantityRequested,
			item.QuantityReceived,
			item.QuantityRemaining,
			item.ItemStatus,
			item.PONumber,
			item.DateAdded)
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
	SELECT orderid, projectid, project_name, foreman_email, sentdate, status, comments
	FROM orders 
	WHERE orderid=$1`

const findOrderItems = `
	SELECT oi.productid, oi.name, oi.uom, oi.requested, oi.received, oi.remaining, oi.status, oi.ponumber, oi.dateadded
	FROM orders o 
	INNER JOIN order_items oi 
	ON o.orderid = oi.orderid 
	WHERE o.orderid=$1
	ORDER BY oi.dateadded DESC`

func (r *OrderRepository) Find(id string) (*supply.Order, error) {
	o := supply.Order{Items: make([]*supply.Item, 0)}

	tx, err := r.db.Begin()
	if err != nil {
		return &supply.Order{}, err
	}

	err = tx.QueryRow(findOrder, id).Scan(
		&o.ID,
		&o.Project.ID,
		&o.Name,
		&o.ForemanEmail,
		&o.SentDate,
		&o.Status,
		&o.Comments)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return &supply.Order{}, errors.New("order not found")
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
			&i.ID,
			&i.Name,
			&i.UOM,
			&i.QuantityRequested,
			&i.QuantityReceived,
			&i.QuantityRemaining,
			&i.ItemStatus,
			&i.PONumber,
			&i.DateAdded)

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
	WHERE projectid=$1
	ORDER BY sentdate DESC`

func (r *OrderRepository) FindDates(projectid string) ([]ordering.ProjectOrder, error) {
	orders := make([]ordering.ProjectOrder, 0)

	rows, err := r.db.Query(findOrderDates, projectid)
	if err != nil {
		return []ordering.ProjectOrder{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var o ordering.ProjectOrder
		err := rows.Scan(&o.ID, &o.SentDate, &o.Status)
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
