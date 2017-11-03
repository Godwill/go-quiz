package main

import (
	"testing"
)

func TestFile(t *testing.T) {
	v := readFile()
	if len(v) < 1 {
		t.Error(
			"For", v,
			"expected", 1,
			"got", len(v),
		)
	}
}
