run:
	go test -c -o jwt-bench.exec
	time ./jwt-bench.exec -test.v -test.benchmem -test.bench ^Benchmark -test.count 5 -test.run ^$ > bench.txt
