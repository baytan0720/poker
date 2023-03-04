all: clean build install

build:
	go build -o bin/poker cmd/poker/*.go
	go build -o bin/poker cmd/daemon/*.go
	go build -o bin/exec internal/exec/exec.go

install:
	cp bin/poker /usr/bin/poker
	cp bin/poker /usr/local/bin/poker
	mkdir -p /var/lib/poker/images
	mkdir -p /var/lib/poker/containers
	mkdir -p /var/lib/poker/bin
	cp -r base /var/lib/poker/images/
	cp bin/exec /var/lib/poker/bin/

uninstall:
	rm -rf /usr/bin/poker /usr/local/bin/poker
	rm -rf /var/lib/poker

clean:
	rm -rf bin/poker