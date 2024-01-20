SHELL := /usr/bin/env /bin/bash -e -u -o pipefail -o errexit -o nounset

.SILENT:
test:
	GOCACHE=/Volumes/RAMDisk/go_cache:/go_cache \
	go version && go test ./...

.SILENT:
docker-test:
	docker run --rm -i -t \
		--memory 128m \
		--memory-reservation 128m \
		--memory-swap 0 \
		-u $(shell id -u):$(shell id -g) \
		-v $(shell pwd):/leetcode-go:ro \
		-v /Volumes/RAMDisk/go_cache:/go_cache \
		-e HOME=/dev/shm \
		-e GOCACHE=/go_cache \
		-w /leetcode-go \
		golang:1.19-alpine \
		sh -c 'go version && go test -parallel -race ./...'
