build:
	@go build .

some: build
	# runs the interpreter against a sample script
	@./some ./scripts/addition.sm

clean:
	@rm some

test:
	@go test -v ./...
