package field

type MaterialList []Item

func (m MaterialList) NewItem(uuid ProductUUID, quantity int) (Item, error) {
	if m.findItem(uuid).ProductUUID != "" {
		return Item{}, ErrItemAlreadyOnList
	}
	return newItem(uuid, quantity), nil
}

func (m MaterialList) findItem(uuid ProductUUID) Item {
	for i := range m {
		if m[i].ProductUUID == uuid {
			return m[i]
		}
	}
	return Item{}
}

func (m MaterialList) removeItem(id ProductUUID) MaterialList {
	//i := o.findItem(id)
	//if i < 0 {
	//	log.Println(ErrItemNotFound)
	//	return
	//}
	//o.MaterialList = append(o.MaterialList[:i], o.MaterialList[i+1:]...)
	return MaterialList{}
}

func (m MaterialList) receivedAll() bool {
	for i := range m {
		if m[i].ItemStatus != Filled {
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
//	i := m.findItem(id)
//	if i < 0 {
//		log.Println(ErrItemNotFound)
//		return
//	}
//
//	m[i].QuantityRequested = qr
//}
