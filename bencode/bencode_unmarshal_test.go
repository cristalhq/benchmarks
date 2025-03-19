package bencode

import (
	"bytes"
	"io"
	"testing"

	IncSW "github.com/IncSW/go-bencode"
	anacrolix "github.com/anacrolix/torrent/bencode"
	cristalhq "github.com/cristalhq/bencode"
	cuberat "github.com/cuberat/go-bencode"
	jackpal "github.com/jackpal/bencode-go"
	lajide "github.com/lajide/bencode"
	lwch "github.com/lwch/bencode"
	marksamman "github.com/marksamman/bencode"
	nabilanam "github.com/nabilanam/bencode/decoder"
	owenliang "github.com/owenliang/dht"
	trim21 "github.com/trim21/go-bencode"
	tumdum "github.com/tumdum/bencoding"
	zeebo "github.com/zeebo/bencode"
)

var unmarshalBenchData = []byte("d4:infod6:lengthi170917888e12:piece lengthi262144e4:name30:debian-8.8.0-arm64-netinst.isoe8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.orge")

func BenchmarkUnmarshal_cristalhq(b *testing.B) {
	var res interface{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := cristalhq.NewDecodeBytes(unmarshalBenchData).Decode(&res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshalReader_cristalhq(b *testing.B) {
	var res interface{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		err := cristalhq.NewDecoder(r).Decode(&res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_IncSW(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := IncSW.Unmarshal(unmarshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal_ZeeboBencode(b *testing.B) {
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := zeebo.DecodeBytes(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_NabilanamBencode(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res := nabilanam.New(unmarshalBenchData).Decode()
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_MarksammanBencode(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		_, err := marksamman.Decode(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal_JackpalBencode(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		_, err := jackpal.Decode(r)
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal_AnacrolixTorrent(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := anacrolix.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_OwenliangDht(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res, err := owenliang.Decode(unmarshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_TumdumBencoding(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := tumdum.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_trim21Bencoding(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := trim21.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_StintsBencode(b *testing.B) {
	b.Skip()
}

func BenchmarkUnmarshal_LwchBencode(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := lwch.Decode(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_ClearcodecnBencode(b *testing.B) {
	b.Skip()
}

func BenchmarkUnmarshal_LajideBencode(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(unmarshalBenchData)
		res, err := lajide.NewDecoder(buf).Decode()
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
		if res == nil {
			// b.Fatal("is nil")
		}
	}
}

func BenchmarkUnmarshal_CuberatGoBencode(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(unmarshalBenchData)
		res, err := cuberat.NewDecoder(buf).Decode()
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
		if res == nil {
			// b.Fatal("is nil")
		}
	}
}
