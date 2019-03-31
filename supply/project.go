package supply

type ProjectRepository interface {
	Save(p *Project) error
	Find(id string) (*Project, error)
}

type Project struct {
	ID           string
	Name         string
	Foreman      string
	ForemanEmail string
	Active       bool
}

func NewProject(id, name, foreman, email string) *Project {
	return &Project{
		ID:           id,
		Name:         name,
		Foreman:      foreman,
		ForemanEmail: email,
		Active:       true,
	}
}

func (p *Project) Close() {
	p.Active = false
}

//func NewProduct(id, category, name, uom string) *Product {
//	return
//}
