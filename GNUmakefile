GO_SRC:=$(shell find . -iname '*.go')

.PHONY: all
all: mentor-matcher

mentor-matcher: $(GO_SRC)
	go build

.PHONY: clean
clean:
	@rm -f mentor-matcher

.PHONY: docker-up
docker-up:
	docker-compose build
	docker-compose up -d

.PHONY: docker-down
docker-down:
	docker-compose down

