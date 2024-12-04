package day_4

import "testing"

func TestPerformTaskOneShort(t *testing.T) {
	expected := 18
	got := PerformTaskOne("data-short.txt", "XMAS")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

func TestPerformTaskOneFull(t *testing.T) {
	expected := 2536
	got := PerformTaskOne("data-full.txt", "XMAS")

	if got != expected {
		t.Errorf("PerformTaskOne() = %v; want %v", got, expected)
	}
}

//func TestPerformTaskTwoShort(t *testing.T) {
//	expected := 9
//	got := PerformTaskTwo("data-short.txt")
//
//	if got != expected {
//		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
//	}
//}
//
//func TestPerformTaskTwoFull(t *testing.T) {
//	expected := 123
//	got := PerformTaskTwo("data-full.txt")
//
//	if got != expected {
//		t.Errorf("PerformTaskTwo() = %v; want %v", got, expected)
//	}
//}
