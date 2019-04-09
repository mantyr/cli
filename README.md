# Golang Command Line Interface

[![Build Status](https://travis-ci.org/mantyr/cli.svg?branch=master)](https://travis-ci.org/mantyr/cli)
[![GoDoc](https://godoc.org/github.com/mantyr/cli?status.png)](http://godoc.org/github.com/mantyr/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/mantyr/cli?v=4)][goreport]
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This don't stable version

## Installation

	$ go get github.com/mantyr/cli

## Example

```GO
package main

import (
	"fmt"
	"os"

	"github.com/mantyr/cli"

	"commands"
)

func main() {
	cli := cli.New("program", os.Args[1:])
	cli.Add(
		command.NewA(),
		command.NewB(),
	)
	err := cli.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr

[build_status]: https://travis-ci.org/mantyr/cli
[godoc]:        http://godoc.org/github.com/mantyr/cli
[goreport]:     https://goreportcard.com/report/github.com/mantyr/cli
