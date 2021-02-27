package main

import (
	"io/ioutil"
	"testing"
)

func TestClientEmpty(t *testing.T) {
	directory, err := ioutil.TempDir(".", "test*")
	if err != nil {
		t.Fail()
	}
	err = listDir(directory)
	if err != nil {
		t.Fail()
	}
}
