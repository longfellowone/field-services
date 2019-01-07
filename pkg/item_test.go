package supply

import (
	"reflect"
	"testing"
)

func Test_newItem(t *testing.T) {
	type args struct {
		uuid ProductUUID
		name string
		uom  UOM
	}
	tests := []struct {
		name string
		args args
		want Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newItem(tt.args.uuid, tt.args.name, tt.args.uom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_receive(t *testing.T) {
	type fields struct {
		ProductUUID       ProductUUID
		Name              string
		UOM               UOM
		QuantityRequested uint
		QuantityReceived  uint
		QuantityRemaining uint
		ItemStatus        ItemStatus
		PONumber          string
	}
	type args struct {
		quantity uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Item{
				ProductUUID:       tt.fields.ProductUUID,
				Name:              tt.fields.Name,
				UOM:               tt.fields.UOM,
				QuantityRequested: tt.fields.QuantityRequested,
				QuantityReceived:  tt.fields.QuantityReceived,
				QuantityRemaining: tt.fields.QuantityRemaining,
				ItemStatus:        tt.fields.ItemStatus,
				PONumber:          tt.fields.PONumber,
			}
			if got := i.receive(tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item.receive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    ItemStatus
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("ItemStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
