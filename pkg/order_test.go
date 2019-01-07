package supply_test

import (
	"supply/pkg"
	"testing"
)

//func TestCreate(t *testing.T) {
//	type args struct {
//		id  supply.OrderUUID
//		pid supply.ProjectUUID
//	}
//	tests := []struct {
//		name string
//		args args
//		want *supply.Order
//	}{
//		{
//			name: "check create order",
//			args: args{
//				id:  "ada612cb-7663-4c64-8fcd-5f3701daeace",
//				pid: "f4e06842-311f-4f21-a10a-06a24e1221de",
//			},
//			want: &supply.Order{
//				OrderUUID:   "ada612cb-7663-4c64-8fcd-5f3701daeace",
//				ProjectUUID: "f4e06842-311f-4f21-a10a-06a24e1221de",
//				MaterialList: supply.MaterialList{
//					Items: nil,
//				},
//				OrderHistory: []supply.Event{{
//					Date:        time.Time{},
//					OrderStatus: supply.Created,
//				}},
//			},
//		},
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := supply.Create(tt.args.id, tt.args.pid); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Create() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestOrder_Send(t *testing.T) {
	type fields struct {
		OrderUUID    supply.OrderUUID
		ProjectUUID  supply.ProjectUUID
		MaterialList supply.MaterialList
		OrderHistory []supply.Event
	}
	tests := []struct {
		name       string
		shouldFail bool
		fields     fields
	}{
		{
			name:       "order fails to update status to send having an item with 0 quantity",
			shouldFail: true,
			fields: fields{
				MaterialList: supply.MaterialList{
					Items: []supply.Item{{
						QuantityRequested: 0,
					}},
				},
				OrderHistory: []supply.Event{{}},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderUUID:    tt.fields.OrderUUID,
				ProjectUUID:  tt.fields.ProjectUUID,
				MaterialList: tt.fields.MaterialList,
				OrderHistory: tt.fields.OrderHistory,
			}
			o.Send()

			if o.OrderHistory[len(o.OrderHistory)-1].OrderStatus == supply.Sent {
				t.Errorf("have: [%v] want: [%v]", o.OrderHistory[len(o.OrderHistory)-1].OrderStatus, supply.Sent)
			}
		})
	}
}

func TestOrder_AddItem(t *testing.T) {
	type fields struct {
		OrderUUID    supply.OrderUUID
		ProjectUUID  supply.ProjectUUID
		MaterialList supply.MaterialList
		OrderHistory []supply.Event
	}
	type args struct {
		uuid supply.ProductUUID
		name string
		uom  supply.UOM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderUUID:    tt.fields.OrderUUID,
				ProjectUUID:  tt.fields.ProjectUUID,
				MaterialList: tt.fields.MaterialList,
				OrderHistory: tt.fields.OrderHistory,
			}
			o.AddItem(tt.args.uuid, tt.args.name, tt.args.uom)
		})
	}
}

//func TestOrder_updateList(t *testing.T) {
//	type fields struct {
//		OrderUUID    supply.OrderUUID
//		ProjectUUID  supply.ProjectUUID
//		MaterialList supply.MaterialList
//		OrderHistory []supply.Event
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &supply.Order{
//				OrderUUID:    tt.fields.OrderUUID,
//				ProjectUUID:  tt.fields.ProjectUUID,
//				MaterialList: tt.fields.MaterialList,
//				OrderHistory: tt.fields.OrderHistory,
//			}
//			o.updateList()
//		})
//	}
//}

func TestOrder_RemoveItem(t *testing.T) {
	type fields struct {
		OrderUUID    supply.OrderUUID
		ProjectUUID  supply.ProjectUUID
		MaterialList supply.MaterialList
		OrderHistory []supply.Event
	}
	type args struct {
		uuid supply.ProductUUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderUUID:    tt.fields.OrderUUID,
				ProjectUUID:  tt.fields.ProjectUUID,
				MaterialList: tt.fields.MaterialList,
				OrderHistory: tt.fields.OrderHistory,
			}
			o.RemoveItem(tt.args.uuid)
		})
	}
}

func TestOrder_ReceiveItem(t *testing.T) {
	type fields struct {
		OrderUUID    supply.OrderUUID
		ProjectUUID  supply.ProjectUUID
		MaterialList supply.MaterialList
		OrderHistory []supply.Event
	}
	type args struct {
		uuid     supply.ProductUUID
		quantity uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &supply.Order{
				OrderUUID:    tt.fields.OrderUUID,
				ProjectUUID:  tt.fields.ProjectUUID,
				MaterialList: tt.fields.MaterialList,
				OrderHistory: tt.fields.OrderHistory,
			}
			o.ReceiveItem(tt.args.uuid, tt.args.quantity)
		})
	}
}

//func TestOrder_missingQuantities(t *testing.T) {
//	type fields struct {
//		OrderUUID    supply.OrderUUID
//		ProjectUUID  supply.ProjectUUID
//		MaterialList supply.MaterialList
//		OrderHistory []supply.Event
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &supply.Order{
//				OrderUUID:    tt.fields.OrderUUID,
//				ProjectUUID:  tt.fields.ProjectUUID,
//				MaterialList: tt.fields.MaterialList,
//				OrderHistory: tt.fields.OrderHistory,
//			}
//			if got := o.missingQuantities(); got != tt.want {
//				t.Errorf("Order.missingQuantities() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestOrder_newEvent(t *testing.T) {
//	type fields struct {
//		OrderUUID    supply.OrderUUID
//		ProjectUUID  supply.ProjectUUID
//		MaterialList supply.MaterialList
//		OrderHistory []supply.Event
//	}
//	type args struct {
//		event OrderStatus
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &supply.Order{
//				OrderUUID:    tt.fields.OrderUUID,
//				ProjectUUID:  tt.fields.ProjectUUID,
//				MaterialList: tt.fields.MaterialList,
//				OrderHistory: tt.fields.OrderHistory,
//			}
//			o.newEvent(tt.args.event)
//		})
//	}
//}

//func TestOrder_lastEvent(t *testing.T) {
//	type fields struct {
//		OrderUUID    OrderUUID
//		ProjectUUID  ProjectUUID
//		MaterialList MaterialList
//		OrderHistory []Event
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   OrderStatus
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &Order{
//				OrderUUID:    tt.fields.OrderUUID,
//				ProjectUUID:  tt.fields.ProjectUUID,
//				MaterialList: tt.fields.MaterialList,
//				OrderHistory: tt.fields.OrderHistory,
//			}
//			if got := o.lastEvent(); got != tt.want {
//				t.Errorf("Order.lastEvent() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestOrder_alreadySent(t *testing.T) {
//	type fields struct {
//		OrderUUID    OrderUUID
//		ProjectUUID  ProjectUUID
//		MaterialList MaterialList
//		OrderHistory []Event
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &Order{
//				OrderUUID:    tt.fields.OrderUUID,
//				ProjectUUID:  tt.fields.ProjectUUID,
//				MaterialList: tt.fields.MaterialList,
//				OrderHistory: tt.fields.OrderHistory,
//			}
//			if got := o.alreadySent(); got != tt.want {
//				t.Errorf("Order.alreadySent() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_createEvent(t *testing.T) {
//	type args struct {
//		status OrderStatus
//	}
//	tests := []struct {
//		name string
//		args args
//		want Event
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := createEvent(tt.args.status); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("createEvent() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestOrderStatus_String(t *testing.T) {
//	tests := []struct {
//		name string
//		s    OrderStatus
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.s.String(); got != tt.want {
//				t.Errorf("OrderStatus.String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
