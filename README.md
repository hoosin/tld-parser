# Go TLD Parser

A Go language library for parsing domain names and extracting their components (subdomain, domain, and suffix) based on the Public Suffix List.

This library is a Go implementation inspired by the functionality of the TypeScript `tld-parse` library.

## Features

- **Accurate Parsing**: Uses the Public Suffix List to accurately decompose domain names.
- **Private Domain Support**: Correctly identifies private domains (e.g., `github.io`).
- **IP Address Handling**: Gracefully handles IP addresses, identifying them as the domain.
- **Generic Function**: The `Parse` function uses Go generics to accept any type with an underlying `string` type.
- **Auto-generating List**: Includes a `Makefile` and a Go program to automatically download the latest Public Suffix List and generate the necessary Go code.

## Installation

To use this library in your project, you can use `go get`:

```sh
go get github.com/hoosin/tld-parser
```

## Usage

Here is a simple example of how to use the `parser.Parse` function.

```go
package main

import (
	"fmt"
	"github.com/hoosin/tld-parser/parser"
)

func main() {
	// Example 1: Simple Domain
	result1 := parser.Parse("www.google.com")
	fmt.Printf("Simple Domain: %+v\n", result1)

	// Example 2: Multi-part Suffix
	result2 := parser.Parse("forums.bbc.co.uk")
	fmt.Printf("Multi-part Suffix: %+v\n", result2)

	// Example 3: Private Suffix
	result3 := parser.Parse("my-project.github.io")
	fmt.Printf("Private Suffix: %+v\n", result3)

	// Example 4: IP Address
	result4 := parser.Parse("192.168.1.1")
	fmt.Printf("IP Address: %+v\n", result4)

	// Example 5: Invalid Input
	result5 := parser.Parse("")
	fmt.Printf("Invalid Input: %v\n", result5) // Outputs: <nil>
}
```

## Development

### Generating the Public Suffix List

This project can automatically generate the `public_suffix_list.go` file from the official Public Suffix List data source. To do this, you need to have `make` and `go` installed.

Run the following command in the project root:

```sh
make all
```

This command will:
1.  Clean any previous data or generated files.
2.  Download the latest `public_suffix_list.dat` from `publicsuffix.org` into the `/data` directory.
3.  Run the Go program in `/cmd/generate` to create the `parser/public_suffix_list.go` file.

### Running Tests

To run the unit tests for this library, navigate to the project root and run:

```sh
go test ./...
```

If you are on an Apple Silicon Mac and encounter a `dyld: missing LC_UUID` error, you can run the tests by disabling CGO:

```sh
CGO_ENABLED=0 go test ./...
```

To generate a test coverage report, you can use the following commands:

```sh
# Generate a coverage profile
CGO_ENABLED=0 go test -coverprofile=coverage.out ./...

# View the report in your browser
go tool cover -html=coverage.out
```
