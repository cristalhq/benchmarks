package natsort_test

import (
	"bytes"
	_ "embed"
	"math/rand"
	"sort"
	"testing"

	cristalhq "github.com/cristalhq/natsort"
	dangogh "github.com/dangogh/naturally"
	evie404 "github.com/evie404/natsort"
	facette "github.com/facette/natsort"
	immortal "github.com/immortal/natcasesort"
	maruel "github.com/maruel/natural"
	miraclesu "github.com/miraclesu/natsort"
)

func BenchmarkStdlib_Strings(b *testing.B) {
	bench(b, "small", smallList, func(list []string) {
		sort.Strings(list)
	})

	bench(b, "huge", tagsList, func(list []string) {
		sort.Strings(list)
	})

	// no need to checkResult(..) because stdlib sort is correct
	// and result is sorted in lexicographical order, not natural
}

func BenchmarkCristalhq_Sort(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		cristalhq.Sort(list)
	})
	checkResult(b, data)

	data = bench(b, "huge", tagsList, func(list []string) {
		cristalhq.Sort(list)
	})
	checkResult(b, data)
}

func BenchmarkCristalhq_Slice(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		sort.Sort(cristalhq.Slice[string](list))
	})
	checkResult(b, data)

	data = bench(b, "huge", tagsList, func(list []string) {
		sort.Sort(cristalhq.Slice[string](list))
	})
	checkResult(b, data)
}

func BenchmarkFacette_Sort(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		facette.Sort(list)
	})
	checkResult(b, data)

	data = bench(b, "huge", tagsList, func(list []string) {
		facette.Sort(list)
	})
	// no check, small variance in the result
}

func BenchmarkMaruel_Sort(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		sort.Sort(maruel.StringSlice(list))
	})
	checkResult(b, data)

	data = bench(b, "huge", tagsList, func(list []string) {
		sort.Sort(maruel.StringSlice(list))
	})
	checkResult(b, data)
}

func BenchmarkDangogh_Slice(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		sort.Sort(dangogh.StringSlice(list))
	})
	// no check, incorrect result
	_ = data

	data = bench(b, "huge", tagsList, func(list []string) {
		sort.Sort(dangogh.StringSlice(list))
	})
	// no check, incorrect result
}

func BenchmarkEvie404_Slice(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		evie404.Strings(list)
	})
	checkResult(b, data)

	data = bench(b, "huge", tagsList, func(list []string) {
		evie404.Strings(list)
	})
	checkResult(b, data)
}

func BenchmarkMiraclesu_Slice(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		miraclesu.Sort(list)
	})
	// no check, incorrect result
	_ = data

	data = bench(b, "huge", tagsList, func(list []string) {
		miraclesu.Sort(list)
	})
	// no check, incorrect result
}

func BenchmarkImmortal_Slice(b *testing.B) {
	data := bench(b, "small", smallList, func(list []string) {
		sort.Sort(immortal.Sort(list))
	})
	// no check, incorrect result
	_ = data

	data = bench(b, "huge", tagsList, func(list []string) {
		sort.Sort(immortal.Sort(list))
	})
	// no check, incorrect result
}

func bench(b *testing.B, name string, input []string, fn func(list []string)) []string {
	b.Helper()

	data := make([]string, len(input))
	r := rand.New(rand.NewSource(69420))
	copy(data, input)

	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r.Shuffle(len(data), func(i, j int) {
				data[i], data[j] = data[j], data[i]
			})
			fn(data)
		}
	})
	return data
}

func checkResult(b *testing.B, list []string) {
	b.Helper()

	ok := sort.SliceIsSorted(list, func(i, j int) bool {
		return cristalhq.Less(list[i], list[j])
	})
	if !ok {
		b.Errorf("not sorted %+v", list)
	}
}

var (
	//go:embed testdata/small.txt
	smallFile []byte

	//go:embed testdata/aws-sdk-go-tags.txt
	tagsFile []byte
)

var (
	smallList []string
	tagsList  []string
)

func init() {
	smallList = readFile(smallFile)
	tagsList = readFile(tagsFile)
}

func readFile(raw []byte) []string {
	lines := bytes.Split(raw, []byte{'\n'})
	res := make([]string, len(lines))

	for i, line := range lines {
		res[i] = string(line)
	}
	return res
}
