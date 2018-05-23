BINARY = bot 

all: clean deps test build run

deps:
	dep ensure

build:
	go build -o ${BINARY} .

clean:
	rm -rf vendor
	rm -f ${BINARY}

run:
	./${BINARY}

test:
	go test

.PHONY: deps build clean run test