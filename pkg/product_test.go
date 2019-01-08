package supply_test

import (
	"reflect"
	"supply/pkg"
	"testing"
)

func TestProduct_ModifyProduct(t *testing.T) {
	type args struct {
		name string
		uom  string
	}
	tests := []struct {
		name string
		got  *supply.Product
		args args
		want *supply.Product
	}{
		{
			name: "uuid must not change",
			got: &supply.Product{
				ProductUUID: "d5820c15-7295-420b-838c-33d04209e882",
			},
			args: args{},
			want: &supply.Product{
				ProductUUID: "d5820c15-7295-420b-838c-33d04209e882",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.got
			p.ModifyProduct(tt.args.name, tt.args.uom)
			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModifyProduct() = %v, want %v", got, tt.want)
			}
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
		{
			name: "must return a *Product",
			args: args{
				uuid: "649739bf-66ee-4023-90bf-2e931c94e024",
				name: "Pencil",
				uom:  "ea",
			},
			want: &supply.Product{
				ProductUUID: "649739bf-66ee-4023-90bf-2e931c94e024",
				Name:        "Pencil",
				UOM:         "ea",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supply.NewProduct(tt.args.uuid, tt.args.name, tt.args.uom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
