package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	keywords := strings.Join(os.Args[1:], " ")
	q := url.QueryEscape(keywords)

	fmt.Println("https://google.co.jp/search?q=" + q)
}
