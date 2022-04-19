package qrcode

import (
	"math/rand"
	"testing"

	boombuler_barcode "github.com/boombuler/barcode/qr"
	cristalhq_qrcode "github.com/cristalhq/qrcode"
	skip2_go_qrcode "github.com/skip2/go-qrcode"
	yeqown_go_qrcode "github.com/yeqown/go-qrcode"
	rsc_qr "rsc.io/qr"
)

var benchURL = "otpauth://totp/Example:alice@bob.com?secret=JBSWY3DPEHPK3PXP&issuer=Example"

func Benchmark_cristalhq_qrcode(b *testing.B) {
	content := benchURL
	bitmap := make([]byte, 2500)

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := cristalhq_qrcode.EncodeInto(bitmap, content, cristalhq_qrcode.M)
		checkErr(b, err)

		c := qr.Image().At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_rsc_qr(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := rsc_qr.Encode(content, rsc_qr.M)
		checkErr(b, err)

		c := qr.Image().At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_yeqown_go_qrcode(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		code, err := yeqown_go_qrcode.New(content)
		checkErr(b, err)

		_ = code
		// c := img.At(10, 10)
		// r, g, b, a := c.RGBA()
		// sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_boombuler_barcode(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		img, err := boombuler_barcode.Encode(content, boombuler_barcode.M, boombuler_barcode.Auto)
		checkErr(b, err)

		c := img.At(10, 10)
		r, g, b, a := c.RGBA()
		sink += int64(r + g + b + a)
	}
	sinkValue(sink)
}

func Benchmark_skip2_go_qrcode(b *testing.B) {
	content := benchURL

	b.ReportAllocs()
	b.ResetTimer()

	var sink int64

	for i := 0; i < b.N; i++ {
		qr, err := skip2_go_qrcode.New(content, skip2_go_qrcode.Medium)
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
