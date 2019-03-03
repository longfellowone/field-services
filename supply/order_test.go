package supply_test

import (
	"field/supply"
	"reflect"
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
	}{{
		name: "must return new *Order",
		args: args{
			id:  "oid1",
			pid: "pid1",
		},
		want: &supply.Order{
			OrderID:   "oid1",
			ProjectID: "pid1",
			Items:     []supply.Item{},
			SentDate:  time.Now().Unix(),
			Status:    supply.New,
		},
	}}
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
	type args struct {
		id   string
		name string
		uom  string
	}
	tests := []struct {
		name    string
		got     *supply.Order
		args    args
		want    *supply.Order
		wantErr bool
	}{{
		name: "must add item to order",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				Name:     "EMT Conduit",
				UOM:      "ft",
				PONumber: "N/A",
			}},
		},
		args: args{
			id:   "pid2",
			name: "Connector",
			uom:  "ea",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				Name:     "EMT Conduit",
				UOM:      "ft",
				PONumber: "N/A",
			}, {
				ItemID:   "pid2",
				Name:     "Connector",
				UOM:      "ea",
				PONumber: "N/A",
			}},
		},
		wantErr: false,
	}, {
		name: "cannot add same item more than once",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID: "pid1",
			}},
		},
		args: args{
			id: "pid1",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID: "pid1",
			}},
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.AddItem(tt.args.id, tt.args.name, tt.args.uom); (err != nil) != tt.wantErr {
				t.Errorf("Order.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got.Items, tt.want.Items) {
				t.Errorf("Order.AddItem() = %v, want %v", got.Items, tt.want.Items)
			}
		})
	}
}

func TestOrder_RemoveItem(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		got     *supply.Order
		args    args
		want    *supply.Order
		wantErr bool
	}{{
		name: "must remove item from order",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID: "pid1",
			}},
		},
		args: args{
			id: "pid1",
		},
		want: &supply.Order{
			Items: []supply.Item{},
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.RemoveItem(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Order.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got.Items, tt.want.Items) {
				t.Errorf("Order.RemoveItem() = %v, want %v", got.Items, tt.want.Items)
			}
		})
	}
}

func TestOrder_UpdateQuantityRequested(t *testing.T) {
	type args struct {
		id       string
		quantity uint
	}
	tests := []struct {
		name    string
		got     *supply.Order
		args    args
		want    *supply.Order
		wantErr bool
	}{{
		name: "must update an items quantity requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 100,
			}},
		},
		args: args{
			id:       "pid1",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
			}},
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.UpdateQuantityRequested(tt.args.id, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Order.UpdateQuantityRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got.Items, tt.want.Items) {
				t.Errorf("Order.UpdateQuantityRequested() = %v, want %v", got.Items[len(got.Items)-1], tt.want.Items[len(got.Items)-1])
			}
		})
	}
}

func TestOrder_Send(t *testing.T) {
	tests := []struct {
		name    string
		got     *supply.Order
		want    *supply.Order
		wantErr bool
	}{{
		name: "order must have at least one item",
		got: &supply.Order{
			Items: []supply.Item{},
		},
		want: &supply.Order{
			Items: []supply.Item{},
		},
		wantErr: true,
	}, {
		name: "quantity requested for items must be great than 0",
		got: &supply.Order{
			Items: []supply.Item{{
				QuantityRequested: 0,
			}},
		},
		want: &supply.Order{
			Items: []supply.Item{{
				QuantityRequested: 0,
			}},
		},
		wantErr: true,
	}, {
		name: "can update order status to sent",
		got: &supply.Order{
			Items: []supply.Item{{
				QuantityRequested: 50,
			}},
			Status: supply.New,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				QuantityRequested: 50,
			}},
			Status: supply.Sent,
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got.Status, tt.want.Status) {
				t.Errorf("Order.Send() = %v, want %v", got.Status, tt.want.Status)
			}
		})
	}
}

func TestOrder_Process(t *testing.T) {
	tests := []struct {
		name string
		got  *supply.Order
		want *supply.Order
	}{{
		name: "must mark item processed",
		got: &supply.Order{
			Status: supply.Sent,
		},
		want: &supply.Order{
			Status: supply.Processed,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			o.Process()
			if got := o; !reflect.DeepEqual(got.Status, tt.want.Status) {
				t.Errorf("Order.Process() = %v, want %v", got.Status, tt.want.Status)
			}
		})
	}
}

func TestOrder_UpdatePO(t *testing.T) {
	type args struct {
		id string
		po string
	}
	tests := []struct {
		name    string
		got     *supply.Order
		args    args
		want    *supply.Order
		wantErr bool
	}{{
		name: "po number must update",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				PONumber: "N/A",
			}},
		},
		args: args{
			id: "pid1",
			po: "HE748563",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				PONumber: "HE748563",
			}},
		},
		wantErr: false,
	}, {
		name: "po must be equal to N/A when empty",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				PONumber: "RX234738",
			}},
		},
		args: args{
			id: "pid1",
			po: "",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:   "pid1",
				PONumber: "N/A",
			}},
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.UpdatePO(tt.args.id, tt.args.po); (err != nil) != tt.wantErr {
				t.Errorf("Order.UpdatePO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got.Items, tt.want.Items) {
				t.Errorf("Order.UpdatePO() = %v, want %v", got.Items[len(got.Items)-1], tt.want.Items[len(got.Items)-1])
			}
		})
	}
}

func TestOrder_ReceiveItem(t *testing.T) {
	type args struct {
		id       string
		quantity uint
	}
	tests := []struct {
		name    string
		got     *supply.Order
		args    args
		want    *supply.Order
		wantErr bool
	}{{
		name: "order must be marked complete when all items received",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  0,
				QuantityRemaining: 0,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Processed,
		},
		args: args{
			id:       "pid1",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  50,
				QuantityRemaining: 0,
				ItemStatus:        supply.Filled,
			}},
			Status: supply.Complete,
		},
		wantErr: false,
	}, {
		name: "order must be marked back to processed if received was modified to < requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  50,
				QuantityRemaining: 0,
				ItemStatus:        supply.Filled,
			}},
			Status: supply.Complete,
		},
		args: args{
			id:       "pid1",
			quantity: 49,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  49,
				QuantityRemaining: 1,
				ItemStatus:        supply.BackOrdered,
			}},
			Status: supply.Processed,
		},
		wantErr: false,
	}, {
		name: "item status must update to filled when requested quantity received",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  0,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Processed,
		},
		args: args{
			id:       "pid1",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  50,
				ItemStatus:        supply.Filled,
			}},
			Status: supply.Complete,
		},
		wantErr: false,
	}, {
		name: "item must be marked back ordered when received > 0 && < requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  0,
				QuantityRemaining: 50,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Processed,
		},
		args: args{
			id:       "pid1",
			quantity: 49,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  49,
				QuantityRemaining: 1,
				ItemStatus:        supply.BackOrdered,
			}},
			Status: supply.Processed,
		},
		wantErr: false,
	}, {
		name: "item must be marked exceeded when received more than requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  51,
				QuantityRemaining: 0,
				ItemStatus:        supply.BackOrdered,
			}},
			Status: supply.Processed,
		},
		args: args{
			id:       "pid1",
			quantity: 51,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ItemID:            "pid1",
				QuantityRequested: 50,
				QuantityReceived:  51,
				QuantityRemaining: 0,
				ItemStatus:        supply.OrderExceeded,
			}},
			Status: supply.Complete,
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.ReceiveItem(tt.args.id, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Order.ReceiveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.ReceiveItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
