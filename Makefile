GOPATH=$(PWD)
export GOPATH

PHP_HOME=/home/ubuntu/.php/versions/trunk

# Hack, fixup
CC=clang

.PHONY: bin/pho

all: bin/pho lib/hacks.so

src/bitbucket.org/binet/go-ffi/pkg/ffi:
	go get bitbucket.org/binet/go-ffi/pkg/ffi

bin/pho: src/bitbucket.org/binet/go-ffi/pkg/ffi
	go build $(GOFLAGS) -o bin/pho pho

lib/%.so: ext/%.c
	${CC} -shared -fPIC -o $@ $^ -Wl,-rpath ${PHP_HOME}/lib -L${PHP_HOME}/lib -I${PHP_HOME}/include/php -I${PHP_HOME}/include/php/Zend -I${PHP_HOME}/include/php/TSRM -I${PHP_HOME}/include/php/main -lphp5
