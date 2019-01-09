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
	}{{
		name: "must return new *Order",
		args: args{
			id:  "650aab76-6e98-4d4f-9380-beb76bb7cb9d",
			pid: "5341fe74-8dde-48e4-bb5b-fb30c473a51f",
		},
		want: &supply.Order{
			OrderID:   "650aab76-6e98-4d4f-9380-beb76bb7cb9d",
			ProjectID: "5341fe74-8dde-48e4-bb5b-fb30c473a51f",
			Items:     []supply.Item{},
			Status:    0,
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
				ProductID: "fcfaa888-529c-44ee-850c-2a5e02d0f7cd",
				Name:      "EMT Conduit",
				UOM:       "ft",
				PONumber:  "N/A",
			}},
		},
		args: args{
			id:   "a35b19b0-6c6a-493a-9739-75cf5addd3d1",
			name: "Connector",
			uom:  "ea",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID: "fcfaa888-529c-44ee-850c-2a5e02d0f7cd",
				Name:      "EMT Conduit",
				UOM:       "ft",
				PONumber:  "N/A",
			}, {
				ProductID: "a35b19b0-6c6a-493a-9739-75cf5addd3d1",
				Name:      "Connector",
				UOM:       "ea",
				PONumber:  "N/A",
			}},
		},
		wantErr: false,
	}, {
		name: "cannot add same item more than once",
		got: &supply.Order{
			Items: []supply.Item{{
				ProductID: "ff12335b-5e6a-4564-9109-a216465602e1",
			}},
		},
		args: args{
			id: "ff12335b-5e6a-4564-9109-a216465602e1",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID: "ff12335b-5e6a-4564-9109-a216465602e1",
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
			if got := o; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.AddItem() = %v, want %v", got, tt.want)
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
				ProductID: "fe84ae35-6e44-4d23-b5f4-1ed57f688af6",
			}},
		},
		args: args{
			id: "fe84ae35-6e44-4d23-b5f4-1ed57f688af6",
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
			if got := o; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.RemoveItem() = %v, want %v", got, tt.want)
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
				ProductID:         "ab54c26d-b8bd-4788-a621-4d419e8130c8",
				QuantityRequested: 100,
			}},
		},
		args: args{
			id:       "ab54c26d-b8bd-4788-a621-4d419e8130c8",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "ab54c26d-b8bd-4788-a621-4d419e8130c8",
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
			if got := o; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.UpdateQuantityRequested() = %v, want %v", got.Items[len(got.Items)-1], tt.want.Items[len(got.Items)-1])
			}
		})
	}
}

func TestOrder_Send(t *testing.T) {
	timeNow := time.Now()
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
				QuantityRequested: 100,
			}},
			OrderDate: timeNow,
			Status:    supply.New,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				QuantityRequested: 100,
			}},
			OrderDate: timeNow,
			Status:    supply.Sent,
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.got
			if err := o.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := o; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.Send() = %v, want %v", got, tt.want)
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
				ProductID: "8eabe046-baab-42ac-a02c-6e7d08d4dd98",
				PONumber:  "N/A",
			}},
		},
		args: args{
			id: "8eabe046-baab-42ac-a02c-6e7d08d4dd98",
			po: "HE748563",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID: "8eabe046-baab-42ac-a02c-6e7d08d4dd98",
				PONumber:  "HE748563",
			}},
		},
		wantErr: false,
	}, {
		name: "po must be equal to N/A when empty",
		got: &supply.Order{
			Items: []supply.Item{{
				ProductID: "8729fb37-6f11-423d-9b03-173b5608348a",
				PONumber:  "RX234738",
			}},
		},
		args: args{
			id: "8729fb37-6f11-423d-9b03-173b5608348a",
			po: "",
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID: "8729fb37-6f11-423d-9b03-173b5608348a",
				PONumber:  "N/A",
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
			if got := o; !reflect.DeepEqual(got, tt.want) {
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
				ProductID:         "c9bfa352-6395-4031-9936-9b317f9d5f21",
				QuantityRequested: 50,
				QuantityReceived:  0,
				QuantityRemaining: 0,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Sent,
		},
		args: args{
			id:       "c9bfa352-6395-4031-9936-9b317f9d5f21",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "c9bfa352-6395-4031-9936-9b317f9d5f21",
				QuantityRequested: 50,
				QuantityReceived:  50,
				QuantityRemaining: 0,
				ItemStatus:        supply.Filled,
			}},
			Status: supply.Complete,
		},
		wantErr: false,
	}, {
		name: "item status must update to filled when requested quantity received",
		got: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "57837a3d-691e-4bb1-8c55-2cdf9146ebd5",
				QuantityRequested: 50,
				QuantityReceived:  0,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Sent,
		},
		args: args{
			id:       "57837a3d-691e-4bb1-8c55-2cdf9146ebd5",
			quantity: 50,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "57837a3d-691e-4bb1-8c55-2cdf9146ebd5",
				QuantityRequested: 50,
				QuantityReceived:  50,
				ItemStatus:        supply.Filled,
			}},
			Status: supply.Complete,
		},
		wantErr: false,
	}, {
		name: "item must be marked backordered when > 0 received && < requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "47b2acf7-147a-47ff-ae16-18f02e668ecd",
				QuantityRequested: 50,
				QuantityReceived:  0,
				QuantityRemaining: 50,
				ItemStatus:        supply.Waiting,
			}},
			Status: supply.Sent,
		},
		args: args{
			id:       "47b2acf7-147a-47ff-ae16-18f02e668ecd",
			quantity: 49,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "47b2acf7-147a-47ff-ae16-18f02e668ecd",
				QuantityRequested: 50,
				QuantityReceived:  49,
				QuantityRemaining: 1,
				ItemStatus:        supply.BackOrdered,
			}},
			Status: supply.Sent,
		},
		wantErr: false,
	}, {
		name: "item must be marked exceeded when received more than requested",
		got: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "ea7d18f9-25ac-4f95-bfd0-c6c54fa5cced",
				QuantityRequested: 50,
				QuantityReceived:  51,
				QuantityRemaining: 0,
				ItemStatus:        supply.BackOrdered,
			}},
			Status: supply.Sent,
		},
		args: args{
			id:       "ea7d18f9-25ac-4f95-bfd0-c6c54fa5cced",
			quantity: 51,
		},
		want: &supply.Order{
			Items: []supply.Item{{
				ProductID:         "ea7d18f9-25ac-4f95-bfd0-c6c54fa5cced",
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
