package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s [mod]\n", os.Args[0])
		os.Exit(1)
	}
	resp, err := http.Get(fmt.Sprintf("https://proxy.golang.org/%s/@v/list", os.Args[1]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
