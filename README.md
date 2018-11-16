# Akamai CLI - IoT EC

### Compiling from Source

If you want to compile it from source, you will need Go 1.10 or later, and the [Dep](https://golang.github.io/dep/) package manager installed:

1. Fetch the package:
  `go get github.com/akamai/cli-iec`
2. Change to the package directory:
  `cd $GOPATH/src/github.com/akamai/cli-iec`
3. Install dependencies using `dep`:
  `dep ensure`
4. Check if binaries build successfully:
  - Linux/macOS/*nix: `go build ./...`
5. Install binaries (to $GOPATH/bin)
  - Linux/macOS/*nix: `go install ./..`
