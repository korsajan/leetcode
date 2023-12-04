
.PHONY: array binary string

all: array binary string

array:
	go test -cover ./array -v

binary:
	go test -cover ./binary -v

string:
	go test -cover ./string -v