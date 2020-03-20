package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "testHello World!!" {
		t.Fatal("Test fail")
	}
}
