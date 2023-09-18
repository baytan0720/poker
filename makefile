POKER_VERSION=$(shell cat VERSION)
BUILD_TIME=$(shell date "+%FT%T%z")
GIT_REVISION=$(shell git rev-parse --short HEAD)
GO_VERSION=$(shell go version | awk '{print $$3}')
LD_FLAGS="-X 'main.pokerVersion=${POKER_VERSION}' -X 'main.buildTime=${BUILD_TIME}' -X 'main.gitRevision=${GIT_REVISION}' -X 'main.goVersion=${GO_VERSION}'"

POKER_CLI_SRC=cmd/cli/*.go
POKER_CLI_NAME=poker
POKER_DAEMON_SRC=cmd/daemon/*.go
POKER_DAEMON_NAME=poker-daemon
POKER_DAEMON_CONFIG=etc/config.yaml

all: mk_proto build install

mk_proto:
	sh sh/mk-proto.sh

build:
	go build -o sbin/$(POKER_CLI_NAME) $(POKER_CLI_SRC)
	go build -o sbin/$(POKER_DAEMON_NAME) -ldflags $(LD_FLAGS) $(POKER_DAEMON_SRC)

install:
	mkdir -p /etc/poker
	mkdir -p /var/poker/containers
	mkdir -p /var/poker/images
	cp sbin/$(POKER_CLI_NAME) /usr/local/bin/$(POKER_CLI_NAME)
	cp sbin/$(POKER_DAEMON_NAME) /usr/local/bin/$(POKER_DAEMON_NAME)
	cp $(POKER_DAEMON_CONFIG) /etc/poker/config.yaml

uninstall:
	rm -f /usr/local/bin/$(POKER_CLI_NAME)
	rm -f /usr/local/bin/$(POKER_DAEMON_NAME)
	rm -f /etc/poker/$(POKER_DAEMON_NAME).yaml