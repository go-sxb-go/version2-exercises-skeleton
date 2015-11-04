package slices

import "reflect"
import "testing"

func TestInsert(t *testing.T) {
	a1 := []int{}
	a1 = Insert(a1, 0, 10)
	expected := []int{10}
	if !reflect.DeepEqual(expected, a1) {
		t.Fatal("expected", expected, "got", a1)
	}

	a2 := []int{1, 2, 3}
	a2 = Insert(a2, 2, 10)
	expected = []int{1, 2, 10, 3}
	if !reflect.DeepEqual(expected, a2) {
		t.Fatal("expected", expected, "got", a2)
	}

	a3 := []int{1, 2, 3, 4}
	a3 = Insert(a3, 4, 10)
	expected = []int{1, 2, 3, 4, 10}
	if !reflect.DeepEqual(expected, a3) {
		t.Fatal("expected", expected, "got", a3)
	}
}

func TestDelete(t *testing.T) {
	a1 := []int{1, 2, 3}
	a1 = Delete(a1, 0)
	expected := []int{2, 3}
	if !reflect.DeepEqual(expected, a1) {
		t.Fatal("expected", expected, "got", a1)
	}

	a2 := []int{1, 2, 3}
	a2 = Delete(a2, 1)
	expected = []int{1, 3}
	if !reflect.DeepEqual(expected, a2) {
		t.Fatal("expected", expected, "got", a2)
	}

	a3 := []int{1, 2, 3, 4}
	a3 = Delete(a3, 3)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(expected, a3) {
		t.Fatal("expected", expected, "got", a3)
	}
}
