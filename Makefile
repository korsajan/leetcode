
.PHONY: array binary linkedlist

all: array binary linkedlist

array:
	go test -cover ./array -v

binary:
	go test -cover ./binary -v

linkedlist:
	go test -cover ./linkedlist -v
