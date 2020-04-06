package main

import "testing"

func TestNATSConnect(t *testing.T) {
	_, err := connect()
	if err != nil {
		t.Error(err)
	}
}
