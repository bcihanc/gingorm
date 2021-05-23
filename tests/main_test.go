package tests

import (
	"errors"
	"testing"
)

func Foo(value bool) (bool, error) {
	if value {
		return true, nil
	} else {
		return false, errors.New("some error")
	}
}

func TestFoo(t *testing.T) {
	if bar, err := Foo(true); err != nil {
		panic(err)
	} else {
		print(bar)
	}
}
