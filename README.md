# Stockbook - TW Stock Crawler

[![codecov](https://codecov.io/gh/suwenyu/stockbook/branch/master/graph/badge.svg?token=2N9LS902PJ)](https://codecov.io/gh/suwenyu/stockbook)
[![Build Status](https://app.travis-ci.com/suwenyu/stockbook.svg?branch=master)](https://app.travis-ci.com/github/suwenyu/stockbook)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsuwenyu%2Fstockbook.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsuwenyu%2Fstockbook?ref=badge_shield)

Stockbook is a stock crawler written in Go (Golang). It is fast, flexible, yet simple module for fetching data from Taiwan Stock Exchange(TWSE). Here express my greatest gratitude to [gogrs](https://github.com/toomore/gogrs).

## Contents

- [Stockbook - TW Stock Crawler](#stockbook---tw-stock-crawler)
	- [Contents](#contents)
	- [Installation](#installation)
	- [Quick start](#quick-start)

## Installation

To install Stockbook package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.16+ is required**), then you can use the below Go command to install Stockbook.

```sh
$ go get -u github.com/suwenyu/stockbook
```

2. Import it in your code:

```go
import "github.com/suwenyu/stockbook"
```

## Quick start

```go
func main() {
	data := stockbook.NewTWSE("2330", time.Now())
	data.RetrievePrevMonth(2)

	fmt.Println(data.FormatData())

}
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsuwenyu%2Fstockbook.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsuwenyu%2Fstockbook?ref=badge_large)