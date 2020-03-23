build:
	goreleaser release

snapshot:
	goreleaser release --skip-publish --snapshot --rm-dist