package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, file)
		file.Close()
	}
}
