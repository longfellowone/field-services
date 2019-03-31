//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"field/supply"
	"field/supply/auth"
	"field/supply/graphql/models"
	"field/supply/ordering"
	"field/supply/search"
	"github.com/99designs/gqlgen/handler"
	"net/http"
)

//if user := auth.ForContext(ctx); user == nil || !user.IsPurchaser {
//	return []ordering.ProjectOrder{}, fmt.Errorf("Access denied")

type Resolver struct {
	osvc ordering.Service
	ssvc search.Service
}

func Initialize(ssvc search.Service, osvc ordering.Service) http.HandlerFunc {
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{
		osvc: osvc,
		ssvc: ssvc,
	}}))
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrder) (*supply.Order, error) {
	user, err := auth.ForContext(ctx)
	if err != nil {
		return &supply.Order{}, err
	}
	order, err := r.osvc.CreateOrder(input.ID, input.ProjectID, input.Name, user.ID, user.Email)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) SendOrder(ctx context.Context, input models.SendOrder) (*supply.Order, error) {
	order, err := r.osvc.SendOrder(input.ID)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) AddOrderItem(ctx context.Context, input models.AddOrderItem) (*supply.Order, error) {
	order, err := r.osvc.AddOrderItem(input.ID, input.ProductID, input.Name, input.Uom)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) RemoveOrderItem(ctx context.Context, input models.RemoveOrderItem) (*supply.Order, error) {
	order, err := r.osvc.RemoveOrderItem(input.ID, input.ProductID)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) ReceiveOrderItem(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	order, err := r.osvc.ReceiveOrderItem(input.ID, input.ProductID, input.Quantity)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) ModifyRequestedQuantity(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	order, err := r.osvc.ModifyRequestedQuantity(input.ID, input.ProductID, input.Quantity)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) CreateProject(ctx context.Context, input models.CreateProject) (*supply.Project, error) {
	user, err := auth.ForContext(ctx)
	if err != nil {
		return &supply.Project{}, err
	}
	project, err := r.osvc.CreateProject(input.ID, input.Name, user.ID, user.Email)
	if err != nil {
		return &supply.Project{}, err
	}
	return project, nil
}
func (r *mutationResolver) CloseProject(ctx context.Context, input models.CloseProject) (*supply.Project, error) {
	project, err := r.osvc.CloseProject(input.ID)
	if err != nil {
		return &supply.Project{}, err
	}
	return project, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Order(ctx context.Context, id string) (*supply.Order, error) {
	order, err := r.osvc.FindOrder(id)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *queryResolver) ProjectOrders(ctx context.Context, id string) ([]ordering.ProjectOrder, error) {
	orders, err := r.osvc.FindProjectOrderDates(id)
	if err != nil {
		return []ordering.ProjectOrder{}, err
	}
	return orders, nil
}
func (r *queryResolver) Products(ctx context.Context, name string) ([]search.Result, error) {
	results, err := r.ssvc.ProductSearch(name)
	if err != nil {
		return []search.Result{}, nil // TODO
	}
	return results, nil
}
func (r *queryResolver) Projects(ctx context.Context, foremanID string) ([]supply.Project, error) {
	projects, err := r.osvc.FindProjectsByForeman(foremanID)
	if err != nil {
		return []supply.Project{}, err
	}
	return projects, nil
}
