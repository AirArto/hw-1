package printtime

import "testing"

func TestPrint(t *testing.T) {
	if _, err := Do(); err != nil {
		t.Fatalf("Fatal error in displaying datetime")
	}
}
