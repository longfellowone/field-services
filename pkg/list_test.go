package supply

import (
	"reflect"
	"testing"
)

func TestMaterialList_UpdateQuantityRequested(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		uuid     ProductUUID
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
			m := MaterialList{
				Items: tt.fields.Items,
			}
			m.UpdateQuantityRequested(tt.args.uuid, tt.args.quantity)
		})
	}
}

func TestMaterialList_UpdatePO(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		uuid ProductUUID
		po   string
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
			m := MaterialList{
				Items: tt.fields.Items,
			}
			m.UpdatePO(tt.args.uuid, tt.args.po)
		})
	}
}

func TestMaterialList_receiveItem(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		uuid     ProductUUID
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
			m := MaterialList{
				Items: tt.fields.Items,
			}
			m.receiveItem(tt.args.uuid, tt.args.quantity)
		})
	}
}

func TestMaterialList_removeItem(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		uuid ProductUUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MaterialList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MaterialList{
				Items: tt.fields.Items,
			}
			if got := m.removeItem(tt.args.uuid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaterialList.removeItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaterialList_findItem(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		uuid ProductUUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MaterialList{
				Items: tt.fields.Items,
			}
			got, got1 := m.findItem(tt.args.uuid)
			if got != tt.want {
				t.Errorf("MaterialList.findItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MaterialList.findItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaterialList_receivedAll(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MaterialList{
				Items: tt.fields.Items,
			}
			if got := m.receivedAll(); got != tt.want {
				t.Errorf("MaterialList.receivedAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
