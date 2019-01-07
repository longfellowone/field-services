package supply_test

import (
	"reflect"
	"supply/pkg"
	"testing"
)

//noinspection ALL
func TestProduct_ModifyProduct(t *testing.T) {
	type fields struct {
		ProductUUID supply.ProductUUID
		Name        string
		UOM         supply.UOM
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
			p := &supply.Product{
				ProductUUID: tt.fields.ProductUUID,
				Name:        tt.fields.Name,
				UOM:         tt.fields.UOM,
			}
			p.ModifyProduct(tt.args.uuid, tt.args.name, tt.args.uom)
		})
	}
}

func TestNewProduct(t *testing.T) {
	type args struct {
		uuid supply.ProductUUID
		name string
		uom  supply.UOM
	}
	tests := []struct {
		name string
		args args
		want *supply.Product
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supply.NewProduct(tt.args.uuid, tt.args.name, tt.args.uom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUOM_String(t *testing.T) {
	tests := []struct {
		name string
		s    supply.UOM
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("UOM.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
