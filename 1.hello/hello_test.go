package hello

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	out, _ := ioutil.TempFile("", "go-hello")

	oldStdout := os.Stdout
	os.Stdout = out
	Hello()
	os.Stdout = oldStdout

	gotBytes, _ := ioutil.ReadFile(out.Name())
	got := string(gotBytes)
	expect := "Hello World!\n"
	if got != expect {
		t.Errorf("Expected '%s', got '%s'", expect, got)
	}

	out.Close()
	os.Remove(out.Name())
}
