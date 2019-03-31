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

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

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
func (r *Resolver) Order() OrderResolver {
	return &orderResolver{}
}
func (r *Resolver) Project() ProjectResolver {
	return &projectResolver{}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type orderResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrder) (*supply.Order, error) {
	user := auth.ForContext(ctx)
	order, err := r.osvc.CreateOrder(input.OrderID, input.ProjectID, input.Name, user.ID, user.Email)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) SendOrder(ctx context.Context, input models.SendOrder) (*supply.Order, error) {
	order, err := r.osvc.SendOrder(input.OrderID)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) AddOrderItem(ctx context.Context, input models.AddOrderItem) (*supply.Order, error) {
	order, err := r.osvc.AddOrderItem(input.OrderID, input.ProductID, input.Name, input.Uom)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) RemoveOrderItem(ctx context.Context, input models.RemoveOrderItem) (*supply.Order, error) {
	order, err := r.osvc.RemoveOrderItem(input.OrderID, input.ProductID)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) ReceiveOrderItem(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	order, err := r.osvc.ReceiveOrderItem(input.OrderID, input.ProductID, input.Quantity)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) ModifyRequestedQuantity(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	order, err := r.osvc.ModifyRequestedQuantity(input.OrderID, input.ProductID, input.Quantity)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *mutationResolver) CreateProject(ctx context.Context, input models.CreateProject) (*supply.Project, error) {
	user := auth.ForContext(ctx)
	project, err := r.osvc.CreateProject(input.ProjectID, input.Name, user.ID, user.Email)
	if err != nil {
		return &supply.Project{}, err
	}
	return project, nil
}
func (r *mutationResolver) CloseProject(ctx context.Context, input models.CloseProject) (*supply.Project, error) {
	project, err := r.osvc.CloseProject(input.ProjectID)
	if err != nil {
		return &supply.Project{}, err
	}
	return project, nil
}

//type orderResolver struct{ *Resolver }
//
//func (r *orderResolver) ProjectID(ctx context.Context, obj *supply.Order) (string, error) {
//	panic("not implemented")
//}
//
//type projectResolver struct{ *Resolver }
//
//func (r *projectResolver) ProjectID(ctx context.Context, obj *supply.Project) (string, error) {
//	panic("not implemented")
//}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Order(ctx context.Context, orderID string) (*supply.Order, error) {
	order, err := r.osvc.FindOrder(orderID)
	if err != nil {
		return &supply.Order{}, err
	}
	return order, nil
}
func (r *queryResolver) ProjectOrders(ctx context.Context, projectID string) ([]ordering.ProjectOrder, error) {
	//user := auth.ForContext(ctx)

	//if user := auth.ForContext(ctx); user == nil || !user.IsPurchaser {
	//	return []ordering.ProjectOrder{}, fmt.Errorf("Access denied")
	//}
	orders, err := r.osvc.FindProjectOrderDates(projectID)
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
	panic("not implemented")
}
