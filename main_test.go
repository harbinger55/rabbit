package main

import "testing"

func TestHelloWorld(t *testing.T) {
	hell := helloWorld()
	if hell != "Hello World" {
		t.Errorf("helloWorld did not return Hello World %s", hell)
	}
}