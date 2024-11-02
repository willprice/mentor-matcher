GO_SRC:=$(shell find . -iname '*.go')

.PHONY: all
all: mentor-matcher

mentor-matcher: $(GO_SRC)
	go build

.PHONY: clean
clean:
	@rm -f mentor-matcher
