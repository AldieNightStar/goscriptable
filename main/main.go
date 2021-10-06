package main

import (
	"fmt"

	"github.com/AldieNightStar/goscriptable"
)

func main() {
	name := "go.s1um"
	fmt.Println("File is exist: ", goscriptable.IsFileExist(name))
}
