package supply

type ProjectRepository interface {
	Save(p *Project) error
	Find(id string) (*Project, error)
}

type Project struct {
	ID      string
	Name    string
	Foreman string
}

//func NewProduct(id, category, name, uom string) *Product {
//	return
//}
