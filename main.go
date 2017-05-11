package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	directory := flag.String("directory", "", "")
	flag.Parse()
	err := filepath.Walk(*directory, visit)
	if err != nil {
		panic(err)
	}
}

func visit(path string, fileInfo os.FileInfo, err error) error {
	log.Println(path)
	return nil
}
