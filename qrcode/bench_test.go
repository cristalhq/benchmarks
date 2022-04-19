package qrcode

import (
	"math/rand"
	"testing"

	boombuler "github.com/boombuler/barcode/qr"
	cristalhq "github.com/cristalhq/qrcode"
	skip2_go "github.com/skip2/go-qrcode"
	yeqown "github.com/yeqown/go-qrcode"
	rsc "rsc.io/qr"
)

var benchURL = "otpauth://totp/Example:alice@bob.com?secret=JBSWY3DPEHPK3PXP&issuer=Example"

func Benchmark_cristalhq(b *testing.B) {
	content := benchURL
	bitmap := make([]byte, 2500)

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := cristalhq.EncodeInto(bitmap, content, cristalhq.M)
		checkErr(b, err)

		c := qr.Image().At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_rsc(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := rsc.Encode(content, rsc.M)
		checkErr(b, err)

		c := qr.Image().At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_yeqown(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		code, err := yeqown.New(content)
		checkErr(b, err)

		_ = code
		// c := img.At(10, 10)
		// r, g, b, a := c.RGBA()
		// sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_boombuler(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		img, err := boombuler.Encode(content, boombuler.M, boombuler.Auto)
		checkErr(b, err)

		c := img.At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_skip2_go(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := skip2_go.New(content, skip2_go.Medium)
		checkErr(b, err)

		c := qr.Image(256).At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
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
