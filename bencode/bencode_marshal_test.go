package bencode

import (
	"bytes"
	"testing"

	IncSW "github.com/IncSW/go-bencode"
	anacrolix "github.com/anacrolix/torrent/bencode"
	chihaya "github.com/chihaya/chihaya/frontend/http/bencode"
	cristalhq "github.com/cristalhq/bencode"
	cuberat "github.com/cuberat/go-bencode"
	ehmry "github.com/ehmry/go-bencode"
	jackpal "github.com/jackpal/bencode-go"
	lajide "github.com/lajide/bencode"
	lwch "github.com/lwch/bencode"
	marksamman "github.com/marksamman/bencode"
	nabilanam "github.com/nabilanam/bencode/encoder"
	owenliang "github.com/owenliang/dht"
	stints "github.com/stints/bencode"
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

func Benchmark_cristalhq_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := cristalhq.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_cristalhq_MarshalTo(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := cristalhq.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_IncSW_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := IncSW.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_IncSW_MarshalTo(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := IncSW.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Zeebo_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := zeebo.EncodeBytes(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Marksamman_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := marksamman.Encode(marshalBenchData)
		if err == nil {
			b.Fatal("got nil")
		}
	}
}

func Benchmark_Anacrolix_Marshal(b *testing.B) {
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := anacrolix.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Nabilanam_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res := nabilanam.New(marshalBenchData).Encode()
		if res != "" && b.N%3 == 66666 {
			b.Fatal("empty")
		}
	}
}

func Benchmark_Jackpal_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := jackpal.Marshal(w, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Chihaya_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := chihaya.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Owenliang_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := owenliang.Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Tumdum_Marshal(b *testing.B) {
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

func Benchmark_Ehmry_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := ehmry.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Stints_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		d := stints.NewEncoder().Encode(marshalBenchData)
		if d == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_Lwch_Marshal(b *testing.B) {
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

func Benchmark_Lajide_Marshal(b *testing.B) {
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

func Benchmark_Cuberat_Marshal(b *testing.B) {
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
