package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("D:/tubes/tba/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Fields(line)
		for i := 0; i < len(token); i++ {
			fmt.Println(token[i])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
