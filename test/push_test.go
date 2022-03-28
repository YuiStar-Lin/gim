package test

import (
	"net/http"
	"testing"
)

///	HTTP2 push demo
func TestPush(t *testing.T) {
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		if push, ok := writer.(http.Pusher); ok {
			t.Log("call push")
			push.Push("/link",nil)
		}


	})

	http.HandleFunc("/link", func(writer http.ResponseWriter, request *http.Request) {
		t.Log("call link")
		writer.Write([]byte("hello world"))
	})
}
