package channelbis

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	stop := make(chan bool)
	c := make(chan bool)
	go Ping(c, stop)
	timeout := time.NewTimer(time.Second)
	select {
	case <-timeout.C:
		t.Fatal("not received ping after 1 second")
	case <-c:
		close(stop)
	}

	_, ok := <-c
	if ok {
		t.Fatal("channel has not been closed correctly when stop has been triggered")
	}
}

func TestPong(t *testing.T) {
	c := make(chan bool)
	out, _ := ioutil.TempFile("", "go-chan")
	oldStdout := os.Stdout
	os.Stdout = out

	go Pong(c)
	timeout := time.NewTimer(2 * time.Second)
	select {
	case <-timeout.C:
		t.Fatal("pong not sent after 2 seconds")
	case c <- true:
		close(c)
		time.Sleep(500 * time.Millisecond)

		gotBytes, _ := ioutil.ReadFile(out.Name())
		got := string(gotBytes)
		expect := "PONG\nEND\n"
		if got != expect {
			t.Errorf("Expected '%s', got '%s'", expect, got)
		}

		os.Stdout = oldStdout
		out.Close()
		os.Remove(out.Name())
	}
}
