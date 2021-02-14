BINARY_NAME=monkey

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: repl
repl: build
	./${BINARY_NAME} repl

.PHONY: clean
clean:
	rm ./${BINARY_NAME}
