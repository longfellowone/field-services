package supply

type ProductRepository interface {
	Save(p *Product) error
	Find(id string) (*Product, error)
}

type Product struct {
	ProductID string
	Name      string
	UOM       string
}

// Modifies a product name and uom
func (p *Product) ModifyProduct(name, uom string) {
	p.Name = name
	p.UOM = uom
}

// Returns a new *Product
func NewProduct(id, name, uom string) *Product {
	return &Product{
		ProductID: id,
		Name:      name,
		UOM:       uom,
	}
}
