run: build bench stat

build:
	go test -c -o natsort-bench.exec

bench:
	time ./natsort-bench.exec \
		-test.v -test.benchmem \
		-test.bench ^Benchmark \
		-test.count 3 \
		-test.run ^$ | tee bench.txt

stat:
	benchstat bench.txt
