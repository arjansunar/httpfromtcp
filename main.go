package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Unable to open file.")
		return
	}

	for {
		buffer := make([]byte, 8)
		_, err = file.Read(buffer)
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			log.Fatal("Unable to read file")
			return
		}
		fmt.Printf("read: %s\n", buffer)
	}
}
