# os-release

[![Build Status](https://travis-ci.org/sonnysideup/os-release.svg?branch=master)](https://travis-ci.org/sonnysideup/os-release)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonnysideup/os-release)](https://goreportcard.com/report/github.com/sonnysideup/os-release)
[![GoDoc](https://godoc.org/github.com/sonnysideup/os-release?status.svg)](https://godoc.org/github.com/sonnysideup/os-release)

Parse and load OS identification data.

## Usage

Example usage is show below.

```go
package main

import (
	"fmt"
	"io/ioutil"

	osr "github.com/sonnysideup/os-release"
)

func main() {
	contents, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}
	info := osr.Parse(string(contents))

	// Inspect the distro lineage
	fmt.Printf("Is %q like fedora? %t", info.Name, info.IsLikeFedora())
	fmt.Printf("Is %q like debian? %t", info.Name, info.IsLikeDebian())

	// List all of the fields on the Data struct, such as ID/Name/Version and others.
	fmt.Printf("%#v", info)
}
```
