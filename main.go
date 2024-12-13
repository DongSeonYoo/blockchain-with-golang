package main

import (
	"github.com/DongSeonYoo/go-coin/explorer"
	"github.com/DongSeonYoo/go-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
