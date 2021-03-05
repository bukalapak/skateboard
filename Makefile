DRAFTER_VERSION    ?= 5.0.0
DRAFTER_PLATFORM   ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
DRAFTER_REPOSITORY ?= https://github.com/bukalapak/drafter-go

drafter:
	@mkdir -p bin
	wget -P bin -nc -O bin/skateboard-rpc \
		${DRAFTER_REPOSITORY}/releases/download/v${DRAFTER_VERSION}/drafter-rpc-${DRAFTER_VERSION}-${DRAFTER_PLATFORM}
	@chmod +x bin/skateboard-rpc
