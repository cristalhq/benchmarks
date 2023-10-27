package bson_test

import (
	"testing"

	cristalhq "github.com/cristalhq/bson"
	globalsign "github.com/globalsign/mgo/bson"
	gomicro "github.com/go-micro/go-bson"
	mongo_driver "go.mongodb.org/mongo-driver/bson"
	mgo "gopkg.in/mgo.v2/bson"
)

var marshalBenchData = testStruct{
	Announce: "udp://tracker.publicbt.com:80/announce",
	AnnounceList: []any{
		"udp://tracker.publicbt.com:80/announce",
		"udp://tracker.openbittorrent.com:80/announce",
		"udp://tracker.openbittorrent.com:80/announce",
		"udp://tracker.openbittorrent.com:80/announce",
	},
	Comment: "Debian CD from cdimage.debian.org",
	Info: map[string]any{
		"name":         "debian-8.8.0-arm64-netinst.iso",
		"length":       int64(170917888),
		"piece length": int64(262144),
	},
}

// var marshalBenchData = map[string]any{
// 	"announce": "udp://tracker.publicbt.com:80/announce",
// 	"announce-list": []any{
// 		"udp://tracker.publicbt.com:80/announce",
// 		"udp://tracker.openbittorrent.com:80/announce",
// 		"udp://tracker.openbittorrent.com:80/announce",
// 		"udp://tracker.openbittorrent.com:80/announce",
// 	},
// 	"comment": "Debian CD from cdimage.debian.org",
// 	"info": map[string]any{
// 		"name":         "debian-8.8.0-arm64-netinst.iso",
// 		"length":       int64(170917888),
// 		"piece length": int64(262144),
// 	},
// }

type testStruct struct {
	Announce     string
	AnnounceList []any
	Comment      string
	Info         map[string]any
}

func Benchmark_cristalhq_Marshal(b *testing.B) {
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := cristalhq.Marshal(marshalBenchData)
		if err != nil {
			panic(err)
		}
		count += int64(len(buf))
	}
	sink(b, count)
}

func Benchmark_cristalhq_MarshalTo(b *testing.B) {
	buf := make([]byte, 0, 512)
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := cristalhq.MarshalTo(buf, marshalBenchData)
		if err != nil {
			panic(err)
		}
		count += int64(len(buf))
	}
	sink(b, count)
}

func Benchmark_mongo_driver_Marshal(b *testing.B) {
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := mongo_driver.Marshal(marshalBenchData)
		mustOk(b, err)
		count += int64(len(buf))
	}
	sink(b, count)
}

func Benchmark_mgo_Marshal(b *testing.B) {
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := mgo.Marshal(marshalBenchData)
		mustOk(b, err)
		count += int64(len(buf))
	}
	sink(b, count)
}

func Benchmark_globalsign_Marshal(b *testing.B) {
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := globalsign.Marshal(marshalBenchData)
		mustOk(b, err)
		count += int64(len(buf))
	}
	sink(b, count)
}

func Benchmark_gomicro_Marshal(b *testing.B) {
	b.ReportAllocs()

	var count int64
	for n := 0; n < b.N; n++ {
		buf, err := gomicro.Marshal(marshalBenchData)
		mustOk(b, err)
		count += int64(len(buf))
	}
	sink(b, count)
}
