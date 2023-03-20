package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, universe!"

	wg.Add(2)
	go updateMessage("Hello, x!")
	go updateMessage("Hello, Go!")
	wg.Wait()

	if msg != "Hello, Go!" {
		t.Error("Expected to find Hello, Go!, but it is not there")
	}
}
