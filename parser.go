package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) (map[string][]string, error) {
	// Buka file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Gagal membuka file: %v", err)
	}
	defer file.Close()

	// Membaca file baris per baris
	scanner := bufio.NewScanner(file)
	productions := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Pisahkan simbol non-terminal dan produksi dengan delimiter "->"
		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Format produksi tidak valid: %s", line)
		}

		// Trim spasi dari simbol non-terminal dan produksi
		symbol := strings.TrimSpace(parts[0])
		productionsStr := strings.TrimSpace(parts[1])

		// Pisahkan produksi dengan delimiter "|"
		productionsArr := strings.Split(productionsStr, "|")
		for i := 0; i < len(productionsArr); i++ {
			productionsArr[i] = strings.TrimSpace(productionsArr[i])
		}

		// Tambahkan aturan produksi ke dalam map
		productions[symbol] = productionsArr
	}

	// Cek error pembacaan file
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Gagal membaca file: %v", err)
	}

	return productions, nil
}
func isAccepted(input string) bool {
	filename := "d:/tubes/tba/production.txt"
	productions, _:= readFile(filename)

	var parse func(string, string) string
	parse = func(str, symbol string) string {
		if _, ok := productions[symbol]; !ok {
			if len(str) > 0 && str[0] == symbol[0] {
				return str[1:]
			}
			return ""
		}

		for _, production := range productions[symbol] {
			remainingStr := str
			valid := true
			for _, part := range production {
				remainingStr = parse(remainingStr, string(part))
				if remainingStr == "" {
					valid = false
					break
				}
			}
			if valid {
				return remainingStr
			}
		}

		return ""
	}

	result := parse(input, "S")
	return result == ""
}

func main() {
	inputString := "xyyxx"
	if isAccepted(inputString) {
		fmt.Printf("String %s diterima\n", inputString)
	} else {
		fmt.Printf("String %s tidak diterima\n", inputString)
	}
}

