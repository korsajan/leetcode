
.PHONY: array binary

all: array binary

array:
	go test -cover ./array -v

binary:
	go test -cover ./binary -v