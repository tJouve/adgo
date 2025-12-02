.PHONY: test fmt run

run:
	@go run ./cmd/aoc

test:
	@go test ./...

fmt:
	@gofmt -s -w .

