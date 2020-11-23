package main

import (
	"flag"
	"fmt"
)

var name2 = flag.String("name", "everyone", "The greeting object.")

func init() {
}

/**
go run demo2.go -name "Robort"
*/
func main() {
	flag.Parse()
	fmt.Printf("Hello world, %s \n", *name2)
}
