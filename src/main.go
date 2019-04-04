package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Started project to crawl web concurrently")
	fc := readTextFile("test.txt")
	fmt.Println(fc)
}

func readTextFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(b)
	return fileContent
}
