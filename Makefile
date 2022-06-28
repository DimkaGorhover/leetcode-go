test:
	@go version && go test ./...

docker-test:
	@docker run --rm -i -t \
		--memory 128m \
		--memory-reservation 128m \
		--memory-swap 0 \
		-v $(shell pwd):/leetcode-go:ro \
		-w /leetcode-go \
		golang:1.18-alpine3.16 \
		sh -c 'go version && go test ./...'
