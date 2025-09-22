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

	currentLine := ""
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

		parts := bytes.Split(buffer, []byte("\n"))
		if len(parts) == 1 {
			currentLine = fmt.Sprint(currentLine, string(parts[0]))
		}
		if len(parts) == 2 {
			currentLine = fmt.Sprint(currentLine, string(parts[0]))
			fmt.Printf("read: %s\n", currentLine)
			currentLine = fmt.Sprint(string(parts[1]))
		}
	}
}
