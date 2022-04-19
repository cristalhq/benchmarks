# qrcode

## Who

```
github.com/boombuler/barcode v1.0.1
github.com/cristalhq/qrcode v0.0.0-00010101000000-000000000000
github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
github.com/yeqown/go-qrcode v1.5.10
rsc.io/qr v0.2.0
```

## Where

```
MacBook Pro (15-inch, 2017)
2,8 GHz Quad-Core Intel Core i7
16 GB 2133 MHz LPDDR3
```

## How

```shell script
# build & run benchmark
$ make run

# or in steps

# build test executable
$ go test -c -o qrcode-bench.exec

# run executable
$ time ./qrcode-bench.exec -v -benchmem -bench ^Benchmark -count 5 -run ^$ > bench.txt
```

## Results

See [bench.txt](https://github.com/cristalhq/benchmarks/blob/main/qrcode/bench.txt)
