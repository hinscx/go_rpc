package main

import (
	"fmt"

	"github.com/hinscx/goecho/echo"
)

func main() {
	res := echo.Echo("Hello World!")
	fmt.Printf("%s\n", res)
}
