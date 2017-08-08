package main

import (
	"fmt"
	"os"
	"bufio"
)

func CountLines(fileName string){

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)
	_, isPrefix, err := reader.ReadLine()

	lines := 0
	for err == nil {
		lines++

		if isPrefix {
			for isPrefix {
				_, isPrefix, err = reader.ReadLine()
				if err != nil {
					break
				}
			}
		}
		_, isPrefix, err = reader.ReadLine()
	}

	fmt.Println(lines)
}

func main() {

	fileName := os.Args[1:]

	CountLines(fileName[0])
}
