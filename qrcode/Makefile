build:
	go test -c -o qrcode-bench.exec

bench:
	time ./qrcode-bench.exec -test.v -test.benchmem -test.bench ^Benchmark -test.count 5 -test.run ^$ > bench.txt

run: build bench
