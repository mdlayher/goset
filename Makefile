make:
	go build github.com/mdlayher/goset

fmt:
	go fmt
	golint .
	errcheck github.com/mdlayher/goset

bench:
	go test -run=NONE -bench=. | column -t

test:
	go test
