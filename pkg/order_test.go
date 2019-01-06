package supply_test

//
//import (
//	"fmt"
//	"supply/pkg"
//	"testing"
//	"time"
//)
//
//var timeNow = func() time.Time {
//	return time.Now() // Some time that you need
//}
//
//func TestCreate(t *testing.T) {
//	have := supply.Create("95f12351-6121-4dda-adba-41f665af7b96", "2e2b84ff-0909-4244-af0d-854c88412262")
//	want := &supply.Order{
//		OrderUUID:   "95f12351-6121-4dda-adba-41f665af7b96",
//		ProjectUUID: "2e2b84ff-0909-4244-af0d-854c88412262",
//		MaterialList: supply.MaterialList{
//			Items: nil,
//		},
//		OrderHistory: []supply.Event{{
//			Date:        timeNow(),
//			OrderStatus: supply.Created,
//		}},
//	}
//
//	if have != want {
//		t.Errorf("Create() = %v want %v", have, want)
//	}
//}
//
//func TestOrder_Send(t *testing.T) {
//
//	have := &supply.Order{
//		OrderHistory: []supply.Event{{
//			Date:        timeNow(),
//			OrderStatus: supply.Created,
//		},
//		},
//	}
//	have.Send()
//
//	want := &supply.Order{
//		OrderHistory: []supply.Event{{
//			Date:        timeNow(),
//			OrderStatus: supply.Sent,
//		}, {
//			Date:        timeNow(),
//			OrderStatus: supply.Created,
//		},
//		},
//	}
//
//	if have != want {
//		fmt.Println(have, want)
//	}
//}
