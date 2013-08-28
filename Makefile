GOPATH=$(PWD)
export GOPATH

.PHONY: bin/pho

all: bin/pho

src/bitbucket.org/binet/go-ffi/pkg/ffi:
	go get bitbucket.org/binet/go-ffi/pkg/ffi

bin/pho: src/bitbucket.org/binet/go-ffi/pkg/ffi
	go build $(GOFLAGS) -o bin/pho pho
