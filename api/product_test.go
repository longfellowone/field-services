package supply_test

import (
	"reflect"
	"supply/api"
	"testing"
)

func TestProduct_ModifyProduct(t *testing.T) {
	type args struct {
		category string
		name     string
		uom      string
	}
	tests := []struct {
		name string
		got  *supply.Product
		args args
		want *supply.Product
	}{{
		name: "category must change",
		got: &supply.Product{
			Category: "Misc",
		},
		args: args{
			category: "Consumables",
		},
		want: &supply.Product{
			Category: "Consumables",
		},
	}, {
		name: "name must change",
		got: &supply.Product{
			Name: "Pencil",
		},
		args: args{
			name: "Marker",
		},
		want: &supply.Product{
			Name: "Marker",
		},
	}, {
		name: "uom must change",
		got: &supply.Product{
			UOM: "ft",
		},
		args: args{
			uom: "ea",
		},
		want: &supply.Product{
			UOM: "ea",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.got
			p.ModifyProduct(tt.args.category, tt.args.name, tt.args.uom)
			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModifyProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProduct(t *testing.T) {
	type args struct {
		id       string
		category string
		name     string
		uom      string
	}
	tests := []struct {
		name string
		args args
		want *supply.Product
	}{{
		name: "must return a new *Product",
		args: args{
			id:       "pid1",
			category: "Consumables",
			name:     "Pencil",
			uom:      "ea",
		},
		want: &supply.Product{
			ProductID: "pid1",
			Category:  "Consumables",
			Name:      "Pencil",
			UOM:       "ea",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supply.NewProduct(tt.args.id, tt.args.category, tt.args.name, tt.args.uom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
