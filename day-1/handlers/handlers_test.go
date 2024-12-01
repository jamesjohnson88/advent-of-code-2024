package handlers

import (
	"testing"
)

func TestPerformTaskOneShort(t *testing.T) {
	expected := 11
	got := PerformTaskOne("../data-1.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskOneFull(t *testing.T) {
	expected := 765748
	got := PerformTaskOne("../data.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}
