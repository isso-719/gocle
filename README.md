# gocle

Gocle is a CLI tool to remove the Go Packages installed on GOBIN.

## Installation

```sh
$ go install github.com/isso-719/gocle@latest
```

## Usage

```sh
 $ gocle -h                 // help for gocle
 $ gocle air                // remove air package from GOBIN
 $ gocle air -y             // remove air package from GOBIN, and do not show confirmation message
 $ gocle air golangci-lint  // you can remove multiple packages at once
```

## License

```
MIT License
```

## Author
- [isso-719](https://github.com/isso-719)