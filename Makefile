


SRC=main.go \
    process.go


all: container_init sleeper

container_init: $(SRC)
	CGO_ENABLED=0 go build .

sleeper: t/sleeper.go
	go build t/sleeper.go


.PHONY: test
test: sleeper container_init
	./container_init -- ./sleeper 1 -- ./sleeper 5


.PHONY: clean
clean:
	-rm container_init sleeper
