package supply_test

import (
	"supply/pkg"
	"testing"
)

func TestProduct(t *testing.T) {

	t.Run("Create new product", func(t *testing.T) {
		have := &supply.Product{
			ProductUUID: "34e02cba-aaf1-4f53-9e8a-94d2e12abdef",
			Name:        "Black Electrical Tape",
			UOM:         supply.FT,
		}
		want := supply.NewProduct(supply.ProductUUID("34e02cba-aaf1-4f53-9e8a-94d2e12abdef"), "Black Electrical Tape", supply.FT)

		if have.ProductUUID != want.ProductUUID {
			t.Errorf("have: [%v] want: [%v]", have.ProductUUID, want.ProductUUID)
		}
		if have.Name != want.Name {
			t.Errorf("have: [%v] want: [%v]", have.Name, want.Name)
		}
		if have.UOM != want.UOM {
			t.Errorf("have: [%v] want: [%v]", have.UOM, want.UOM)
		}
	})

	t.Run("Modify Product name and UOM", func(t *testing.T) {
		have := supply.NewProduct(supply.ProductUUID("34e02cba-aaf1-4f53-9e8a-94d2e12abdef"), "Black Electrical Tape", supply.FT)
		have.ModifyProduct(supply.ProductUUID("34e02cba-aaf1-4f53-9e8a-94d2e12abdef"), "3/4\" EMT Connector", supply.EA)
		want := &supply.Product{
			ProductUUID: "34e02cba-aaf1-4f53-9e8a-94d2e12abdef",
			Name:        "3/4\" EMT Connector",
			UOM:         supply.EA,
		}

		if have.ProductUUID != want.ProductUUID {
			t.Errorf("have: [%v] want: [%v] error: ProductUUID should not change", have.ProductUUID, want.ProductUUID)
		}
		if have.Name != want.Name {
			t.Errorf("have: [%v] want: [%v]", have.Name, want.Name)
		}
		if have.UOM != want.UOM {
			t.Errorf("have: [%v] want: [%v]", have.UOM, want.UOM)
		}
	})

	t.Run("Name", func(t *testing.T) {

	})
}

func Equal(p1 *supply.Product, p2 *supply.Product) bool {
	return p1 == p2
}
