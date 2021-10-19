package gee

import "testing"

func TestHelloEmpty(t *testing.T) {
	msg := "123"
	if msg != "" {
		t.Fatalf(`Hello("") = %q , want "", error`, msg)
	}
}
