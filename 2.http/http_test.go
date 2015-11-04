package http

import "testing"

func TestGetGolang(t *testing.T) {
	md5, err := GetGolang()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := "d9f531d1a9dae6c644b14482fe71bcb3"
	if md5 != expect {
		t.Errorf("Expected '%v', got '%v'", expect, md5)
	}
}
