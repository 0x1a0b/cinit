

SRC=main.go \
    process.go

all: cinit sleeper

cinit: $(SRC)
	CGO_ENABLED=0 go build .

sleeper: t/sleeper.go
	go build t/sleeper.go


.PHONY: test
test: sleeper cinit
	./cinit -- ./sleeper 1 -- ./sleeper 5


.PHONY: clean
clean:
	-rm cinit sleeper
