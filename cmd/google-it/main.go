package main

import (
	"fmt"
	"os"

	googleit "github.com/val00274/go-google-it"
)

func main() {
	app := googleit.NewApp()
	fmt.Println(app.GoogleURL(os.Args[:1]...))
}
