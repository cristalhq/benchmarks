# bencode

## Who

```
github.com/IncSW/go-bencode v0.2.2
github.com/anacrolix/torrent v1.42.0
github.com/chihaya/chihaya v2.0.0-rc.2+incompatible
github.com/cristalhq/bencode v0.2.1
github.com/cuberat/go-bencode v1.0.0
github.com/ehmry/go-bencode v1.1.1
github.com/jackpal/bencode-go v1.0.0
github.com/lajide/bencode v0.0.0-20190629152703-fb09cf4e9a4a
github.com/lwch/bencode v1.0.0
github.com/marksamman/bencode v0.0.0-20150821143521-dc84f26e086e
github.com/nabilanam/bencode v0.0.0-20190329141800-cde546b1530b
github.com/owenliang/dht v0.0.0-20180118074908-44baeeba7b13
github.com/stints/bencode v0.0.0-20160226205624-0ba65bd80165
github.com/tumdum/bencoding v0.0.0-20160911135503-fcfcb8bd55e9
github.com/zeebo/bencode v1.0.0
```

## Where

```
MacBook Pro (16-inch, 2019)
2,6 GHz 6-Core Intel Core i7
16 GB 2667 MHz DDR4
```

## How

```shell script
# build & run benchmark
$ make run

# or in steps

# build test executable
$ go test -c -o bencode-bench.exec

# run executable
$ time ./bencode-bench.exec -v -benchmem -bench ^Benchmark -count 5 -run ^$ > bench.txt
```

## Results

See [bench.txt](https://github.com/cristalhq/benchmarks/blob/main/bencode/bench.txt)
