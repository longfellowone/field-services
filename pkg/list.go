package field

type MaterialList []Item

func (m MaterialList) NewItem(uuid ProductUUID) Item {
	return newItem(uuid)
}

func (m MaterialList) UpdateQuantityRequested(uuid ProductUUID, quantity int) Item {
	return m.findItem(uuid).updateQuantityRequested(quantity)
}

func (m MaterialList) ReceiveItem(uuid ProductUUID, quantity int) Item {
	return Item{}
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
