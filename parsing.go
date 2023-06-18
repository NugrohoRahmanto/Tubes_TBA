package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct untuk menyimpan produksi dalam grammar
type Production struct {
	NonTerminal string
	Rules       [][]string
}

func main() {
	// Baca aturan produksi dari file teks
	productions, err := readProductionsFromFile("D:/tubes/tba/production.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Masukkan string yang akan diperiksa
	inputString := ""
	for inputString != "404" {
		fmt.Println("masukkan string yang akan di cek (404)")
		fmt.Scan(&inputString)
		lastTwo := inputString[len(inputString)-2:]
		isAccepted := isStringAccepted(inputString, "S", productions)
		if isAccepted && lastTwo != "xx" && lastTwo != "aa" {
			fmt.Println("String diterima oleh grammar.")
		} else {
			fmt.Println("String tidak diterima oleh grammar.")
		}

	}
}

// Fungsi untuk membaca aturan produksi dari file teks
func readProductionsFromFile(filename string) ([]Production, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	productions := make([]Production, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Format aturan produksi tidak valid: %s", line)
		}

		nonTerminal := strings.TrimSpace(parts[0])
		rulesStr := strings.TrimSpace(parts[1])
		rules := strings.Split(rulesStr, "|")

		prodRules := make([][]string, len(rules))
		for i, rule := range rules {
			prodRules[i] = strings.Fields(strings.TrimSpace(rule))
		}

		productions = append(productions, Production{
			NonTerminal: nonTerminal,
			Rules:       prodRules,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return productions, nil
}

func isStringAccepted(input string, symbol string, productions []Production) bool {
	// Jika input kosong, cek apakah simbol saat ini adalah simbol akhir
	if input == "" {
		for _, production := range productions {
			if production.NonTerminal == symbol {
				for _, rule := range production.Rules {
					if len(rule) == 0 {
						return true
					}
				}
			}
		}
		return false
	}

	// Cari aturan produksi yang cocok dengan simbol saat ini
	for _, production := range productions {
		if production.NonTerminal == symbol {
			for _, rule := range production.Rules {
				// Pemanggilan rekursif untuk setiap simbol dalam aturan produksi
				if isRuleAccepted(input, rule, productions) {
					return true
				}
			}
		}
	}

	return false
}

// Fungsi untuk memeriksa apakah suatu aturan produksi diterima
func isRuleAccepted(input string, rule []string, productions []Production) bool {
	// Memeriksa setiap simbol dalam aturan produksi
	for _, symbol := range rule {
		// Jika simbol adalah simbol non-terminal, panggil rekursif
		if isNonTerminal(symbol, productions) {
			// Jika tidak ada input yang tersisa, aturan produksi tidak dapat dilanjutkan
			if input == "" {
				return false
			}
			// Pemanggilan rekursif untuk simbol non-terminal
			if isStringAccepted(input, symbol, productions) {
				// Pemotongan input sesuai dengan panjang simbol
				input = input[len(symbol):]
			} else {
				return false
			}
		} else {
			// Jika simbol adalah simbol terminal, cocokkan dengan input
			if len(input) >= len(symbol) && input[:len(symbol)] == symbol {
				input = input[len(symbol):]
			} else {
				return false
			}
		}
	}

	return true
}

// Fungsi untuk memeriksa apakah suatu simbol adalah simbol non-terminal
func isNonTerminal(symbol string, productions []Production) bool {
	for _, production := range productions {
		if production.NonTerminal == symbol {
			return true
		}
	}
	return false
}