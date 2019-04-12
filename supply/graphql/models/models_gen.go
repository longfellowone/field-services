// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AddOrderItem struct {
	ID        string `json:"id"`
	ProductID string `json:"productID"`
	Name      string `json:"name"`
	Uom       string `json:"uom"`
}

type CloseProject struct {
	ID string `json:"id"`
}

type CreateOrder struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectID"`
	Name      string `json:"name"`
}

type CreateProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ModifyQuantity struct {
	ID        string `json:"id"`
	ProductID string `json:"productID"`
	Quantity  int    `json:"quantity"`
}

type RemoveOrderItem struct {
	ID        string `json:"id"`
	ProductID string `json:"productID"`
}

type SendOrder struct {
	ID       string `json:"id"`
	Comments string `json:"comments"`
}
