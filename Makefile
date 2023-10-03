
.PHONY: array calculator


array:
	go test -cover ./array -v

calculator:
	go test -cover ./calculator -v