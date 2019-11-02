package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"shu/pkg/mp"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "file path required\n")
		os.Exit(1)
	}
	path := args[1]
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile: %s: %s\n", path, err)
		os.Exit(1)
	}
	fmt.Printf("file ->\n%s\n", f)

	md := mp.Parse(string(f))
	html := mp.ToHTML(md)
	fmt.Printf("-> html\n%s", html)
}
