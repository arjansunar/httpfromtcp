package main

import (
	"bytes"
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
	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	res := make(chan string)
	currentLine := ""
	go func(res chan string) {
		defer f.Close()
		defer close(res)
		for {
			buffer := make([]byte, 8)
			_, err := f.Read(buffer)
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal("Unable to read file")
			}

			parts := bytes.Split(buffer, []byte("\n"))
			for i, part := range parts {
				if i == 0 {
					currentLine += string(part)
				} else {
					res <- currentLine
					currentLine = ""
					currentLine += string(part)
				}
			}

		}
	}(res)
	return res
}
