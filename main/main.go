package main

import (
	"time"

	"github.com/AldieNightStar/goscriptable"
)

func main() {
	for {
		println(goscriptable.RandomString())
		time.Sleep(time.Millisecond * 100)
	}
}
