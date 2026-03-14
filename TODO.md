# TODO

## 🔥 Nästa steg?

- Add debug outputs in output, scanner etc.
- Tester 🔨
- Github Actions 🔨


.github/workflows/release.yml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - uses: goreleaser/goreleaser-action@v5
        with:
          version: latest
          args: release --clean



.goreleaser.yaml
project_name: fe

builds:
  - main: ./cmd/fe
    goos:
      - windows
      - darwin
      - linux
    goarch:
      - amd64
      - arm64

archives:
  - format: tar.gz
    files:
      - README.md

checksum:
  name_template: "{{ .ProjectName }}_{{ .Version }}_checksums.txt"

release:
  draft: true
