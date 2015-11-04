package httpserver

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	go HttpServer("12345")
	time.Sleep(time.Second)
	res, err := http.Get("localhost:12345")
	if err != nil {
		t.Fatal("expected no error, got", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	expected := "Hello World!"
	if expected != string(body) {
		t.Fatalf("expect '%v', got '%v'", expected, string(body))
	}
}
