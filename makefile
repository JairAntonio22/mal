all: lisp

lisp: *.go
	go build .

fmt:
	go fmt .

test:
	go test .

clean:
	rm -f lisp
