package day_2

import "testing"

func TestPerformTaskOneShort(t *testing.T) {
	expected := 2
	got := PerformTaskOne("data-short.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskOneFull(t *testing.T) {
	expected := 411
	got := PerformTaskOne("data-full.txt")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskTwoShort(t *testing.T) {
	expected := 4
	got := PerformTaskTwo("data-short.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}

func TestPerformTaskTwoFull(t *testing.T) {
	expected := 465 // current 458, lost 7 in attempted refactor
	got := PerformTaskTwo("data-full.txt")

	if got != expected {
		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
	}
}
