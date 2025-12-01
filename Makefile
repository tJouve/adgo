.PHONY: test fmt run

run:
	@go run ./cmd/aoc --day $(DAY) --part $(PART) --input $(INPUT)

test:
	@go test ./...

fmt:
	@gofmt -s -w .

