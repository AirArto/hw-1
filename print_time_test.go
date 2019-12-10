package print_time

import "testing"

func TestPrint(t *testing.T) {
	if tryVar := PrintDt(); tryVar == false {
		t.Fatalf("Fatal error in displaying datetime")
	}
}
