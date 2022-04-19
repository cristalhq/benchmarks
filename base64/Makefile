build:
	go test -c -o base64-bench.exec

bench:
	time ./base64-bench.exec -test.v -test.benchmem -test.bench ^Benchmark -test.count 5 -test.run ^$ > bench.txt

run: build bench
