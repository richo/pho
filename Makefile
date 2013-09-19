GOPATH=$(PWD)
export GOPATH

PHP_HOME=$(HOME)/.php/versions/trunk
CGO_LDFLAGS=-L$(PWD)/lib -lhacks -Wl,-rpath $(PWD)/lib
export CGO_LDFLAGS

# Hack, fixup
CC=clang
export CC
PHPFLAGS=-Wl,-rpath ${PHP_HOME}/lib -L${PHP_HOME}/lib -I${PHP_HOME}/include/php -I${PHP_HOME}/include/php/Zend -I${PHP_HOME}/include/php/TSRM -I${PHP_HOME}/include/php/main -lphp5

.PHONY: bin/pho test

all: bin/pho

bin/pho: lib/libhacks.so
	go build $(GOFLAGS) -o bin/pho pho

lib/lib%.so: ext/%.c
	${CC} -shared -fPIC -g -o $@ $^ ${PHPFLAGS}

.test/%: test/%.c
	${CC} -o $@ $^ ${PHPFLAGS}

test: .test/hacks
	./.test/hacks
