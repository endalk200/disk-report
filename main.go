package main

import (
	"flag"
	"fmt"
)

func main() {
	var path *string = flag.String("p", ".", "Path to scan")

	var verbose *bool = flag.Bool("v", false, "Verbose output")

	flag.Parse()

	root := *path

	if !*verbose {
		fmt.Println(getSize(root))
	} else {
		getSizeVerbose(root)
	}
}
