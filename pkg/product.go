package supply

type ProductUUID string

type Product struct {
	ProductUUID
	Name string
	UOM  UOM
}

func (p *Product) ModifyProduct(uuid ProductUUID, name string, uom UOM) {
	p.Name = name
	p.UOM = uom
}

func NewProduct(uuid ProductUUID, name string, uom UOM) *Product {
	return &Product{
		ProductUUID: uuid,
		Name:        name,
		UOM:         uom,
	}
}

type UOM int

const (
	EA UOM = iota
	FT
	M
)

func (s UOM) String() string {
	switch s {
	case EA:
		return "ea"
	case FT:
		return "ft"
	case M:
		return "m"
	}
	return ""
}
