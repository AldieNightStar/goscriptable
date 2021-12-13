package main

import (
	"fmt"

	"github.com/AldieNightStar/goscriptable"
)

func main() {
	url := "https://pastebin.com/zNiTn7hJ"
	data := goscriptable.FindInside(
		string(string(goscriptable.HttpGet(url))),
		"abc",
	)

	fmt.Println(data)
}
