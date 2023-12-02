package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func GetFileContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	reader := bufio.NewReader(f)
	contents, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}

func GetTestInput(day int) string {
	return GetFileContent(fmt.Sprintf("inputs/day%d/test.txt", day))
}

func GetInput(day int) string {
	return GetFileContent(fmt.Sprintf("inputs/day%d/input.txt", day))
}
