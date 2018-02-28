package testtools

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expected %+v. Got: %+v", expected, actual))
	}
}

func AssertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Error(fmt.Sprintf("Expected True. Got False."))
	}
}
