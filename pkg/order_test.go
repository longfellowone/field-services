package supply_test

import (
	"reflect"
	"supply/pkg"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	type args struct {
		id  string
		pid string
	}
	tests := []struct {
		name string
		args args
		want *supply.Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supply.Create(tt.args.id, tt.args.pid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test for len() then test [len()-1}
func TestOrder_AddItem(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	type args struct {
		id   string
		name string
		uom  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.AddItem(tt.args.id, tt.args.name, tt.args.uom); (err != nil) != tt.wantErr {
				t.Errorf("Order.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_RemoveItem(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.RemoveItem(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Order.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_UpdateQuantityRequested(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	type args struct {
		id       string
		quantity uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.UpdateQuantityRequested(tt.args.id, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Order.UpdateQuantityRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_Send(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_UpdatePO(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	type args struct {
		id string
		po string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.UpdatePO(tt.args.id, tt.args.po); (err != nil) != tt.wantErr {
				t.Errorf("Order.UpdatePO() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_ReceiveItem(t *testing.T) {
	type fields struct {
		OrderID   string
		ProjectID string
		Items     []supply.Item
		OrderDate time.Time
		Status    supply.OrderStatus
	}
	type args struct {
		id       string
		quantity uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderID:   tt.fields.OrderID,
				ProjectID: tt.fields.ProjectID,
				Items:     tt.fields.Items,
				OrderDate: tt.fields.OrderDate,
				Status:    tt.fields.Status,
			}
			if err := o.ReceiveItem(tt.args.id, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Order.ReceiveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
