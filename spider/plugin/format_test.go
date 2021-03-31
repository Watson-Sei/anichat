package plugin

import (
	"fmt"
	"testing"
)

func TestShapingTime(t *testing.T) {
	got := ShapingTime("23:00")
	want := fmt.Sprintf("23:00 03/30/21")
	if got != want {
		t.Error("Expected values do not match.")
	}
}