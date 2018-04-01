package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	writeFile()
	readFile()
}

func writeFile() {
	buf := []byte("hello")
	ioutil.WriteFile("./test.txt", buf, 666)
}

func readFile() {
	buf, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
