package bencode

import (
	"bytes"
	"math/rand/v2"
	"testing"

	IncSW "github.com/IncSW/go-bencode"
	anacrolix "github.com/anacrolix/torrent/bencode"
	cristalhq "github.com/cristalhq/bencode"
	cuberat "github.com/cuberat/go-bencode"
	jackpal "github.com/jackpal/bencode-go"
	lajide "github.com/lajide/bencode"
	lwch "github.com/lwch/bencode"
	marksamman "github.com/marksamman/bencode"
	nabilanam "github.com/nabilanam/bencode/encoder"
	owenliang "github.com/owenliang/dht"
	stints "github.com/stints/bencode"
	trim21 "github.com/trim21/go-bencode"
	tumdum "github.com/tumdum/bencoding"
	zeebo "github.com/zeebo/bencode"
)

var marshalBenchData = map[string]interface{}{
	"announce": ("udp://tracker.publicbt.com:80/announce"),
	"announce-list": []interface{}{
		[]interface{}{("udp://tracker.publicbt.com:80/announce")},
		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
		[]interface{}{
			"udp://tracker.openbittorrent.com:80/announce",
			"udp://tracker.openbittorrent.com:80/announce",
		},
	},
	"comment": []byte("Debian CD from cdimage.debian.org"),
	"info": map[string]interface{}{
		"name":         []byte("debian-8.8.0-arm64-netinst.iso"),
		"length":       170917888,
		"piece length": 262144,
	},
}

func BenchmarkMarshal_cristalhq(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := cristalhq.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalTo_cristalhq(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := cristalhq.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_IncSW(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := IncSW.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalTo_IncSW(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := IncSW.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Zeebo(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := zeebo.EncodeBytes(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Marksamman(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := marksamman.Encode(marshalBenchData)
		if err == nil {
			b.Fatal("got nil")
		}
	}
}

func BenchmarkMarshal_Anacrolix(b *testing.B) {
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := anacrolix.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Nabilanam(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res := nabilanam.New(marshalBenchData).Encode()
		if res != "" && b.N%3 == 66666 {
			b.Fatal("empty")
		}
	}
}

func BenchmarkMarshal_Jackpal(b *testing.B) {
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := jackpal.Marshal(w, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Owenliang(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := owenliang.Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Tumdum(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := tumdum.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Stints(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		d := stints.NewEncoder().Encode(marshalBenchData)
		if d == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkMarshal_Lwch(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := lwch.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Lajide(b *testing.B) {
	b.Skip()
	buf := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := lajide.NewEncoder(buf).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_Cuberat(b *testing.B) {
	b.Skip()
	buf := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := cuberat.NewEncoder(buf).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshal_trim21(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := trim21.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func sink[T any](tb testing.TB, v T) {
	if rand.Float32() > 2 {
		tb.Fatal(v)
	}
}
