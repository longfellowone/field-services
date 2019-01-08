package supply_test

import (
	"reflect"
	"supply/pkg"
	"testing"
)

func TestProduct_ModifyProduct(t *testing.T) {
	type fields struct {
		ProductUUID string
		Name        string
		UOM         string
	}
	type args struct {
		uuid string
		name string
		uom  string
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
		uuid string
		name string
		uom  string
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
