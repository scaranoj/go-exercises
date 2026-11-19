package main

import (
	"fmt"

	format "github.com/scaranoj/go-exercises/package_exercise/do-format"
	"github.com/scaranoj/go-exercises/package_exercise/math"
)

func main() {
	x := math.Double(5)
	y := format.Formatter(x)
	fmt.Println(y)
}
