package main

import (
	"testing"
)

func Test_dine(t *testing.T) {

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("Expected 5, got %d", len(orderFinished))
		}
	}
}
