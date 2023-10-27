package bson_test

import (
	"math/rand"
	"testing"
)

func mustOk(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Fatal(err)
	}
}

func sink[T any](tb testing.TB, v T) {
	if rand.Float32() > 2 {
		tb.Fatal(v)
	}
}
