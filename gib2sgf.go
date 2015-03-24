package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Handler func(line string) string

func handleHS(line string) string {
	return "(;GM[1]FF[4]CA[UTF-8]AP[gib2sgf]"
}

func handleSTO(line string) string {
	// STO 0 3 2 15 15
	atot := "abcdefghijklmnopqrst"
	colors := []string{"", "B", "W"}
	tok := strings.Split(line, " ")
	c, _ := strconv.Atoi(tok[3])
	x, _ := strconv.Atoi(tok[4])
	y, _ := strconv.Atoi(tok[5])

	return fmt.Sprintf(";%s[%s%s]", colors[c], atot[x:x+1], atot[y:y+1])
}

func handleGE(line string) string {
	return ")"
}

func HandleLine(line string) {
	l := strings.Trim(line, " ")

	handlers := map[string]Handler{
		"\\HS": handleHS,
		"STO":  handleSTO,
		"\\GE": handleGE,
	}

	for prefix, h := range handlers {
		if strings.HasPrefix(l, prefix) {
			fmt.Println(h(l))
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need two arguments.")
	}

	f, _ := os.Open(os.Args[1])
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		HandleLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
