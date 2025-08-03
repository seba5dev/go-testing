package unit_test

import (
	"testing"

	"github.com/seba5dev/go-testing/unit"
)

func TestFooer_Foo(t *testing.T) {
	result := unit.Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooer_Number(t *testing.T) {
	result := unit.Fooer(1)
	if result != "1" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
