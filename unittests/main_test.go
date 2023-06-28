package add_test

import (
	add "github.com/tomiok/course-23/unittests"
	"strconv"
	"testing"
)

// package name
// 1) should end with "_test"
// 2) use the test function func TestAddition(t *testing.T) {}
// there are couple of testing types. I.e. T (unit testing), B (for benchmarking), M main tests.

func TestAddition(t *testing.T) {
	value1 := 1
	value2 := 5
	expected := 6
	actual := add.Adding(value1, value2)

	if actual != expected {
		t.Fatal("the result should be " + strconv.Itoa(expected))
	}
}

func TestAdditionNotOK(t *testing.T) {
	value1 := 1
	value2 := 5
	expected := 100
	actual := add.Adding(value1, value2)

	if actual != expected {
		t.Fatal("the result should be " + strconv.Itoa(expected))
	}
}
