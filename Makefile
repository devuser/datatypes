GO = go
GO_GET = $(GO) get -v
GO_TEST = $(GO) test -buildvcs=false -v
GO_BUILD = $(GO) build -buildvcs=false


all: test

dep:
	-$(GO_GET) github.com/smartystreets/goconvey/convey

tidy:
	$(GO) mod tidy

fmt:
	@echo 'fmt noop'

test: dep
	$(GO_TEST)

.PHONY: all dep tidy test
