build:
	go mod download
	go generate ./...
	cd cmd && CGO=0 GOARCH=arm go build -o timepiece.arm

release:
	goreleaser release

snapshot:
	goreleaser release --skip-publish --snapshot --rm-dist
