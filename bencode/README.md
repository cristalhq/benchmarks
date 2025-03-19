# bencode

## Who

```
github.com/IncSW/go-bencode v0.2.2
github.com/anacrolix/torrent v1.56.1
github.com/cristalhq/bencode v0.4.0
github.com/cuberat/go-bencode v1.0.0
github.com/jackpal/bencode-go v1.0.2
github.com/lajide/bencode v0.0.0-20190629152703-fb09cf4e9a4a
github.com/lwch/bencode v1.0.0
github.com/marksamman/bencode v0.0.0-20150821143521-dc84f26e086e
github.com/nabilanam/bencode v0.0.0-20190329141800-cde546b1530b
github.com/owenliang/dht v0.0.0-20180118074908-44baeeba7b13
github.com/stints/bencode v0.0.0-20160226205624-0ba65bd80165
github.com/trim21/go-bencode v0.0.7
github.com/tumdum/bencoding v0.0.0-20160911135503-fcfcb8bd55e9
github.com/zeebo/bencode v1.0.0
```

## Where

```
goos: darwin
goarch: arm64
pkg: github.com/cristalhq/benches/bencode
cpu: Apple M1 Pro
```

## How

```shell script
$ task run-bencode
```

## Results

See `bench.txt` and `bench-raw.txt`

```
args: [bench-raw.txt]name                            time/op
Marshal_cristalhq-10             588ns ± 1%
MarshalTo_cristalhq-10           431ns ± 0%
Marshal_IncSW-10                 429ns ± 1%
MarshalTo_IncSW-10               367ns ± 1%
Marshal_Zeebo-10                2.91µs ± 0%
Marshal_Marksamman-10            686ns ± 1%
Marshal_Anacrolix-10            1.69µs ± 0%
Marshal_Jackpal-10              2.61µs ± 1%
Marshal_trim21-10               1.03µs ± 0%
Unmarshal_cristalhq-10           656ns ± 2%
UnmarshalReader_cristalhq-10     749ns ± 1%
Unmarshal_IncSW-10               610ns ± 1%
Unmarshal_ZeeboBencode-10       3.04µs ± 2%
Unmarshal_NabilanamBencode-10    812ns ± 0%
Unmarshal_MarksammanBencode-10  1.63µs ± 2%
Unmarshal_JackpalBencode-10     1.45µs ± 1%
Unmarshal_OwenliangDht-10        979ns ± 1%

name                            alloc/op
Marshal_cristalhq-10              976B ± 0%
MarshalTo_cristalhq-10           0.00B     
Marshal_IncSW-10                  512B ± 0%
MarshalTo_IncSW-10               0.00B     
Marshal_Zeebo-10                1.65kB ± 0%
Marshal_Marksamman-10           1.14kB ± 0%
Marshal_Anacrolix-10            2.03kB ± 2%
Marshal_Jackpal-10              2.31kB ± 1%
Marshal_trim21-10                 816B ± 0%
Unmarshal_cristalhq-10            960B ± 0%
UnmarshalReader_cristalhq-10    1.52kB ± 0%
Unmarshal_IncSW-10                960B ± 0%
Unmarshal_ZeeboBencode-10       5.98kB ± 0%
Unmarshal_NabilanamBencode-10   1.22kB ± 0%
Unmarshal_MarksammanBencode-10  5.74kB ± 0%
Unmarshal_JackpalBencode-10     1.66kB ± 0%
Unmarshal_OwenliangDht-10       1.33kB ± 0%

name                            allocs/op
Marshal_cristalhq-10              3.00 ± 0%
MarshalTo_cristalhq-10            0.00     
Marshal_IncSW-10                  1.00 ± 0%
MarshalTo_IncSW-10                0.00     
Marshal_Zeebo-10                  36.0 ± 0%
Marshal_Marksamman-10             10.0 ± 0%
Marshal_Anacrolix-10              19.0 ± 0%
Marshal_Jackpal-10                39.0 ± 0%
Marshal_trim21-10                 17.0 ± 0%
Unmarshal_cristalhq-10            18.0 ± 0%
UnmarshalReader_cristalhq-10      20.0 ± 0%
Unmarshal_IncSW-10                18.0 ± 0%
Unmarshal_ZeeboBencode-10         79.0 ± 0%
Unmarshal_NabilanamBencode-10     30.0 ± 0%
Unmarshal_MarksammanBencode-10    45.0 ± 0%
Unmarshal_JackpalBencode-10       50.0 ± 0%
Unmarshal_OwenliangDht-10         37.0 ± 0%

```
