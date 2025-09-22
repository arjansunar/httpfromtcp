package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Unable to listen.", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Unable to accept connections.", err)
			return
		}
		fmt.Println("a connection has been accepted")
		for line := range getLinesChannel(conn) {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("connection has been closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	res := make(chan string, 1)
	currentLine := ""
	go func() {
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
	}()
	return res
}
