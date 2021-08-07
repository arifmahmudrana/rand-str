Random string generator
================

This CLI allows random string easyly from string set with given size.

## Releases
To run this CLI go to [releases](releases) and find the appropriate build from release page.

## Usage
This CLI supports two features
1. Setting up stores by adding necessary configurations time series etc.
2. Importing data of store

To run command first grab an executable from [release](#releases).
```sh
~ rand-str --help # For help or rand-str -h
# Outputs
# usage: main <string_len> <string_set>

~ rand-str 6 # For random string of 6 characters
~ rand-str 9 "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" # For random string of 9 characters from "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" set
```

## Test
To run tests locally make sure
- Need Go 1.16
- Setup project or clone outside `GOPATH` to check run `go env GOPATH`

To run test
```sh
go test -v -race ./...
```

To get test coverage
```sh
go test -v -race -coverprofile=/tmp/go-cover.`basename $(pwd)`.out ./... && go tool cover -func=/tmp/go-cover.`basename $(pwd)`.out && unlink /tmp/go-cover.`basename $(pwd)`.out
```
Or
```sh
go test -v -race -coverprofile=/tmp/go-cover.`basename $(pwd)`.out ./... && go tool cover -html=/tmp/go-cover.`basename $(pwd)`.out && unlink /tmp/go-cover.`basename $(pwd)`.out
```

## Build
First clone the repo then follow the [Tests](#test) step.
```sh
go build
```

## Contributing
First clone the repo then follow the [Tests](#test) step.

We use [svu](https://github.com/caarlos0/svu) for automatic tagging and build of our repo using Github actions see [svu commit message pattern](https://github.com/caarlos0/svu/blob/master/README.md#commit-messages-vs-what-they-do) for your commit message.
