package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./TestNaja.png")
	text, _ := client.Text()
	fmt.Println(text)
}
