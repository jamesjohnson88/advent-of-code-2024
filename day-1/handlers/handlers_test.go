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

func TestPerformTaskTwoShort(t *testing.T) {
	expected := 31
	got := PerformTaskTwo("../data-1.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}

func TestPerformTaskTwoFull(t *testing.T) {
	expected := 27732508
	got := PerformTaskTwo("../data.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}
