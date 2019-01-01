package material

import (
	"errors"
)

var (
	ErrPOalreadyExists = errors.New("PO already exists")
	ErrPOnotFound      = errors.New("PO not found")
)

type PurchaseOrder struct {
	PONumber PONumber
	Supplier Supplier
}

//func (p *PurchaseOrder) update(n PONumber, s Supplier) PurchaseOrder {
//	return PurchaseOrder{
//		PONumber: n,
//		Supplier: s,
//	}
//}
//
//func (p *PurchaseOrder) remove() *PurchaseOrder {
//	return &PurchaseOrder{}
//}

type (
	PONumber string
	Supplier string
)

type OrderPOs struct {
	POs []PurchaseOrder
}

func (p *OrderPOs) add(n PONumber, s Supplier) error {
	if p.exists(n) {
		return ErrPOalreadyExists
	}

	p.POs = append(p.POs, PurchaseOrder{
		PONumber: n,
		Supplier: s,
	})
	return nil
}

func (p *OrderPOs) remove(n PONumber) error {
	for i, po := range p.POs {
		if po.PONumber == n {
			copy(p.POs[i:], p.POs[i+1:])
			p.POs[len(p.POs)-1] = PurchaseOrder{}
			p.POs = p.POs[:len(p.POs)-1]
			return nil
		}
	}
	return ErrPOnotFound
}

func (p *OrderPOs) exists(n PONumber) bool {
	for _, po := range p.POs {
		if po.PONumber == n {
			return true
		}
	}
	return false
}
