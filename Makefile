
.PHONY: array binary string linkedlist

all: array binary string linkedlist

array:
	go test -cover ./array -v

binary:
	go test -cover ./binary -v

linkedlist:
	go test -cover ./linkedlist -v

string:
	go test -cover ./string -v