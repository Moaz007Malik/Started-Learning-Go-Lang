package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I Learn Go Lang")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", byteWritten)

	bs := []byte("Go programming language is cool")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
