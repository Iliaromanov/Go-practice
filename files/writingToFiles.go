package main

import (
	//"fmt"
	//"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err) // panic stops execution of current function
	}
}

func main() {
	// Get current working directory
	cwd, err := os.Getwd()
	checkError(err)

	writeFilePath := filepath.Join(cwd, "writeToThis.txt")
	// Write array of bytes (string) to file
	msg1 := []byte("This is first message!")
	msg2 := []byte("This is second message!!")
	err = ioutil.WriteFile(writeFilePath, msg1, 0644) // 0644 allows user to read and write, all others only read
	checkError(err)
	err = ioutil.WriteFile(writeFilePath, msg2, 0644) // msg2 overwrites msg1
	checkError(err)
}