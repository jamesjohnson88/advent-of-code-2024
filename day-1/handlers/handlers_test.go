package handlers

import (
	"testing"
)

func TestPerformTaskOne(t *testing.T) {
	expected := ":)"
	got := PerformTaskOne("../data-1.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %q; want %q", got, expected)
	}
}
