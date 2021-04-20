package rw

import (
	"testing"
)

var targetFile string = "../test.txt"

func TestReadFromFile(t *testing.T) {
	want := "This is a test."
	got := ReadFromFile(targetFile)
	if got != want {
		t.Errorf("Expected %s got %s", want, got)
	}
}
