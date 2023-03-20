package main

import "testing"

// That will work fine, but if you use -race inside the command, a failure because we don't know which one will be the last one to be executed.
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
