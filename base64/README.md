# base64

## Who

```
github.com/cristalhq/base64 v0.1.2
github.com/segmentio/asm v1.1.4
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
$ go test -c -o base64-bench.exec

# run executable
$ time ./base64-bench.exec -v -benchmem -bench ^Benchmark -count 5 -run ^$ > bench.txt
```

## Results

See [bench.txt](https://github.com/cristalhq/benchmarks/blob/main/base64/bench.txt)
