
.PHONY: array binary string linkedlist calculator

all: array binary string linkedlist calculator

array:
	go test -cover ./array -v

calculator:
	go test -cover ./calculator -v

binary:
	go test -cover ./binary -v

linkedlist:
	go test -cover ./linkedlist -v

string:
	go test -cover ./string -v