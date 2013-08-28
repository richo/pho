GOPATH=$(PWD)
export GOPATH

# Hack, fixup
CC=gcc

.PHONY: bin/pho

all: bin/pho lib/hacks.so

src/bitbucket.org/binet/go-ffi/pkg/ffi:
	go get bitbucket.org/binet/go-ffi/pkg/ffi

bin/pho: src/bitbucket.org/binet/go-ffi/pkg/ffi
	go build $(GOFLAGS) -o bin/pho pho

lib/%.so: ext/%.c
	${CC} -shared -fPIC -o $@ $^ -Wl,-rpath /home/ubuntu/.php/versions/trunk/lib -L/home/ubuntu/.php/versions/trunk/lib -I/home/ubuntu/.php/versions/trunk/include -lphp5
