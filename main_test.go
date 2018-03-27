package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendHelloWorld(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/example", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(ExampleHandler)
	handler.ServeHTTP()

	for name, test := range tests {
	}
}
