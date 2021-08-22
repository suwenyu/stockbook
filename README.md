# Stockbook - TW Stock Crawler

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
