package day_3

import "testing"

func TestPerformTaskOneShort(t *testing.T) {
	expected := 161
	got := PerformTaskOne("data-short.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskOneFull(t *testing.T) {
	expected := 171183089
	got := PerformTaskOne("data-full.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskTwoShort(t *testing.T) {
	expected := 48
	got := PerformTaskTwo("data-short-2.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}

func TestPerformTaskTwoFull(t *testing.T) {
	expected := 63866497
	got := PerformTaskTwo("data-full.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}
