package postgres

import (
	"database/sql"
	"errors"
	"github.com/longfellowone/field-services/supply"
	"log"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

const findProject = `
	SELECT projectid, name, foreman, foreman_email, active
	FROM projects
	WHERE projectid=$1`

func (r *ProjectRepository) Find(id string) (*supply.Project, error) {
	p := &supply.Project{}
	err := r.db.QueryRow(findProject, id).Scan(&p.ID, &p.Name, &p.Foreman, &p.ForemanEmail, &p.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return &supply.Project{}, errors.New("order not found")
		} else {
			return &supply.Project{}, err
		}
	}
	return p, nil
}

const saveProject = `
	INSERT INTO projects (projectid, name, foreman, foreman_email, active)
	VALUES ($1,$2,$3,$4,$5)
	ON CONFLICT ON CONSTRAINT projects_pk
	DO UPDATE SET
		projectid=EXCLUDED.projectid,
		name=EXCLUDED.name,
		foreman=EXCLUDED.foreman,
		foreman_email=EXCLUDED.foreman_email,
		active=EXCLUDED.active`

func (r *ProjectRepository) Save(p *supply.Project) error {
	_, err := r.db.Exec(saveProject, p.ID, p.Name, p.Foreman, p.ForemanEmail, p.Active)
	if err != nil {
		return err
	}
	return nil
}

const findProjectsByForeman = `
	SELECT projectid, name, foreman, foreman_email, active
	FROM projects
	WHERE foreman=$1
	AND active=TRUE`

func (r *ProjectRepository) FindAllByForeman(foremanid string) ([]supply.Project, error) {
	orders := make([]supply.Project, 0)

	rows, err := r.db.Query(findProjectsByForeman, foremanid)
	if err != nil {
		return []supply.Project{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var o supply.Project
		err := rows.Scan(&o.ID, &o.Name, &o.Foreman, &o.ForemanEmail, &o.Active)
		if err != nil {
			return []supply.Project{}, nil
		}
		orders = append(orders, o)
	}
	err = rows.Err()
	if err != nil {
		return []supply.Project{}, nil
	}

	return orders, nil
}
