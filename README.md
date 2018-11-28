# Akamai CLI - IoT EC

### Compiling from Source

If you want to compile it from source, you will need Go 1.10 or later, and the [Dep](https://golang.github.io/dep/) package manager installed:

1. Fetch the package:
  `go get github.com/bbucko/cli-iec`
2. Change to the package directory:
  `cd $GOPATH/src/github.com/bbucko/cli-iec`
3. Install dependencies using `dep`:
  `dep ensure`
4. Check if binaries build successfully:
  - Linux/macOS/*nix: `go build ./...`
5. Install binaries (to $GOPATH/bin)
  - Linux/macOS/*nix: `go install ./..`


User flow:
* Generate key pair - `akamai jwt generate --name keys_1`
* Configurate new property - `akamai iec configure --auth jwt --hostname iec.kwapiszewski.com --mqtt --ws --https --jwtKey keys_1 --activate --namespace kwapiszewski_ns --jurisdiction EU`
* Subscribe - `akamai iec subscribe --namespace kwapiszewski_ns --jurisdiction EU`
* Publish - `akamai iec publish --namespace kwapiszewski_ns --jurisdiction EU --message "Hello world"`
* Generate token - `akamai jwt token --name keys_1 --clientId clientId --authGroup publisher`
