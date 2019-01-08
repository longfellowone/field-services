package supply_test

import (
	"supply/pkg"
	"testing"
)

func TestItemStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    supply.ItemStatus
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
