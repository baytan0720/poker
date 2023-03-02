all: clean build install

build:
	go build -o bin/poker *.go

install:
	cp bin/poker /usr/bin/poker
	cp bin/poker /usr/local/bin/poker
	mkdir -p /var/lib/poker/images
	mkdir -p /var/lib/poker/containers
	cp -r base /var/lib/poker/images/

uninstall:
	rm -rf /usr/bin/poker /usr/local/bin/poker
	rm -rf /var/lib/poker

clean:
	rm -rf bin/poker