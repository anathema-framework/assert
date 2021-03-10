package assert

import (
	"errors"
	"reflect"
	"testing"
)

func Equal(t *testing.T, got, expecting interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expecting) {
		t.Fatalf("got %v, expecting %v", got, expecting)
	}
}

func ErrorIs(t *testing.T, got, expecting error) {
	t.Helper()
	if !errors.Is(got, expecting) {
		t.Fatalf("got %v, expecting %v", got, expecting)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got error %v", err)
	}
}
