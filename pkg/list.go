package supply

import (
	"log"
)

type MaterialList struct {
	Items []Item
}

func (m MaterialList) UpdateQuantityRequested(uuid ProductUUID, quantity uint) {
	i, item := m.findItem(uuid)

	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return
	}

	m.Items[i].QuantityRequested = quantity
}

func (m MaterialList) UpdatePO(uuid ProductUUID, po string) {
	i, item := m.findItem(uuid)

	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return
	}

	if po == "" {
		po = "N/A"
	}

	m.Items[i].PONumber = po
}

func (m MaterialList) receiveItem(uuid ProductUUID, quantity uint) {
	i, item := m.findItem(uuid)

	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return
	}

	m.Items[i] = m.Items[i].receive(quantity)
}

func (m MaterialList) addItem(uuid ProductUUID, name string) MaterialList {
	_, item := m.findItem(uuid)
	if item.ProductUUID != "" {
		log.Println(ErrItemAlreadyOnList)
		return m
	}

	m.Items = append(m.Items, newItem(uuid, name))
	return m
}

func (m MaterialList) removeItem(uuid ProductUUID) MaterialList {
	i, item := m.findItem(uuid)
	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return m
	}

	m.Items = append(m.Items[:i], m.Items[i+1:]...)
	return m
}

func (m MaterialList) findItem(uuid ProductUUID) (int, Item) {
	for i := range m.Items {
		if m.Items[i].ProductUUID == uuid {
			return i, m.Items[i]
		}
	}
	return -1, Item{}
}

func (m MaterialList) receivedAll() bool {
	for i := range m.Items {
		if m.Items[i].ItemStatus != Filled && m.Items[i].ItemStatus != OrderExceeded {
			return false
		}
	}
	return true
}

//func (m MaterialList) AdjustQuantityRequested(id ProductUUID, qr QuantityRequested) {
//	if qr <= 0 {
//		log.Println(ErrQuantityZero)
//		return
//	}
//
//	i := m.FindItem(id)
//	if i < 0 {
//		log.Println(ErrItemNotFound)
//		return
//	}
//
//	m[i].QuantityRequested = qr
//}

//func (m MaterialList) FindItem(uuid ProductUUID) Item {
//	for i := range m {
//		if m[i].ProductUUID == uuid {
//			m[i].Index = i
//			return m[i]
//		}
//	}
//	return Item{}
//}
