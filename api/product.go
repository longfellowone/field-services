package supply

type ProductRepository interface {
	Save(p *Product) error
	Find(id string) (*Product, error)
}

type Product struct {
	ProductID string
	Category  string
	Name      string
	UOM       string
}

// Modifies a product name and uom
func (p *Product) ModifyProduct(category, name, uom string) {
	p.Category = category
	p.Name = name
	p.UOM = uom
}

// Returns a new *Product
func NewProduct(id, category, name, uom string) *Product {
	return &Product{
		ProductID: id,
		Category:  category,
		Name:      name,
		UOM:       uom,
	}
}
