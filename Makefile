release_version:= v0.1.1

export GO111MODULE=on

.PHONY: test
test:
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet
	go build -o bin/fleet github.com/mhausenblas/kcf/cmd/fleet

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

.PHONY: release
release:
	curl -sL https://git.io/goreleaser | bash -s -- --rm-dist --config .goreleaser.yml

.PHONY: publish
publish:
	git tag ${release_version}
	git push --tags