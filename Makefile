run-hr:
	go build -o heartrate/heartrate github.com/pieterclaerhout/go-sports/heartrate
	./heartrate/heartrate

run-pacer:
	go build -o pacer/pacer github.com/pieterclaerhout/go-sports/pacer
	@echo
	./pacer/pacer -distance 10 -duration 60m -endo
	@echo
	./pacer/pacer -distance 31.7 -duration 61m -endo