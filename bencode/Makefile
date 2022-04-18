build:
	go test -c -o bencode-bench.exec

bench:
	time ./bencode-bench.exec -test.v -test.benchmem -test.bench ^Benchmark -test.count 5 -test.run ^$ > bench.txt

run: build bench
