package supply

type Product struct {
	ProductUUID string
	Name        string
	UOM         string
}

func (p *Product) ModifyProduct(name, uom string) {
	p.Name = name
	p.UOM = uom
}

func NewProduct(uuid, name, uom string) *Product {
	return &Product{
		ProductUUID: uuid,
		Name:        name,
		UOM:         uom,
	}
}
