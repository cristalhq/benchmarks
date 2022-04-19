package base64_test

import (
	"math/rand"
	"testing"

	golang "encoding/base64"

	cristalhq "github.com/cristalhq/base64"
	segmentio "github.com/segmentio/asm/base64"
)

func Benchmark_cristalhq(b *testing.B) {
	data := make([]byte, 8192)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	var v int64
	for i := 0; i < b.N; i++ {
		res := cristalhq.StdEncoding.EncodeToString(data)
		v += int64(len(res))
	}

	sinkValue(v)
}

func Benchmark_golang(b *testing.B) {
	data := make([]byte, 8192)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	var v int64
	for i := 0; i < b.N; i++ {
		res := golang.StdEncoding.EncodeToString(data)
		v += int64(len(res))
	}

	sinkValue(v)
}

func Benchmark_segmentio(b *testing.B) {
	data := make([]byte, 8192)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	var v int64

	for i := 0; i < b.N; i++ {
		res := segmentio.StdEncoding.EncodeToString(data)
		v += int64(len(res))
	}

	sinkValue(v)
}

func checkErr(tb testing.TB, err error) {
	if err != nil {
		tb.Fatal(err)
	}
}

func sinkValue(v int64) {
	if rand.Float64() > 2 {
		panic(v)
	}
}
