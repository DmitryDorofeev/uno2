all: get build

run: build
	./bin/main

build:
	gb build --tags "libsqlite3 darwin"

get:
	gb vendor restore

install:
	cp build/uno /usr/local/bin/uno
	cp build/init.sh /etc/init.d/uno