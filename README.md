# Awesome Miner Go SDK

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-round)](https://godoc.org/github.com/johanmeiring/awesomeminer-go-sdk) [![Go Report Card](https://goreportcard.com/badge/github.com/johanmeiring/awesomeminer-go-sdk)](https://goreportcard.com/report/github.com/johanmeiring/awesomeminer-go-sdk) [![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-round)](https://github.com/johanmeiring/awesomeminer-go-sdk/blob/master/LICENSE)

This Go package provides a wrapper for the [Awesome Miner API](http://www.awesomeminer.com/help/api.aspx).  

Note: **This is currently a work in progress and does not yet contain implementations for all available endpoints.**  Contributions are welcome ;-)

## Installing
To install this SDK to your `GOPATH`:

    go get -u github.com/johanmeiring/awesomeminer-go-sdk

... or you can use a dependency tool such as Govendor or dep.

## Basic usage
```go
// If your Awesome Miner instance is behind a proxy with HTTP Basic Auth:
c := awesomeminer.NewClient("http://proxy", "username", "password")
// otherwise:
c := awesomeminer.NewClient("http://miner:17790", "", "")

// Fetch all notifications:
n, err := c.GetNotifications()

// Fetch all miners:
m, err := c.GetMiners()

// Fetch a specific miner:
miner, err := c.GetMiner(5)
```

## License
This SDK is distributed under the MIT License.  See the LICENSE file for more details.

## Donations
Donations are very welcome, and can be made to the following addresses:
* BTC: 1AWHJcUBha35FnuuWat9urRW2FNc4ftztv
* ETH: 0xAF1Aac4c40446F4C46e55614F14d9b32d712ECBc
