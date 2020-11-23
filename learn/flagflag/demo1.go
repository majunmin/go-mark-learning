package main

import (
	"flag"
	"fmt"
)

var name1 string

func init() {
	flag.StringVar(&name1, "name", "everyone", "The greeting object.")
}

/**
go run demo1.go
*/
func learn() {
	fmt.Printf("Hello world, %s \n", name1)
}
