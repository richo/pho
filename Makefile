GOPATH=$(PWD)
export GOPATH

SLASH_HOME=$(HOME)/.slash/versions/trunk
SLASH_OPTS=-lgmp -liconv -lpcre
CGO_LDFLAGS=-L$(PWD)/lib -lhacks
CGO_CFLAGS=-I$(SLASH_HOME)/include
export CGO_LDFLAGS
export CGO_CFLAGS

# Hack, fixup
CC=clang
export CC

.PHONY: bin/pho test

all: bin/pho

bin/pho: lib/libhacks.so
	go build -work $(GOFLAGS) -o bin/pho pho > .lastbuild

lib/lib%.so: ext/%.c
	${CC} -shared -fPIC -g -o $@ $^ -I$(SLASH_HOME)/include $(SLASH_HOME)/lib/libslash.a $(SLASH_OPTS)

.test/%: test/%.c
	${CC} -o $@ $^ ${PHPFLAGS}

test: .test/hacks
	./.test/hacks

clean:
	rm -rf `cat .lastbuild`
	rm .lastbuild
