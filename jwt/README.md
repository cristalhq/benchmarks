# jwt

## Who

```
github.com/cristalhq/jwt/v4 v4.0.0
github.com/golang-jwt/jwt v3.2.2+incompatible
github.com/lestrrat-go/jwx v1.2.23
github.com/pascaldekloe/jwt v1.10.0
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
$ go test -c -o jwt-bench.exec

# run executable
$ time ./jwt-bench.exec -v -benchmem -bench ^Benchmark -count 5 -run ^$ > bench.txt
```

## Results

See [bench.txt](https://github.com/cristalhq/benchmarks/blob/main/jwt/bench.txt)
