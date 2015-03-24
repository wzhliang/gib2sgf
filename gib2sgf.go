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

func getMetaValue(line string) (v string) {
	v = strings.Split(line, "=")[1]
	v = strings.Trim(v, "\\]")

	return
}

func handleWhiteName(line string) string {
	return fmt.Sprintf("PW[%s]", getMetaValue(line))
}

func handleBlackName(line string) string {
	return fmt.Sprintf("PB[%s]", getMetaValue(line))
}

func handleResult(line string) string {
	m := map[string]string{
		"black": "B",
		"white": "W",
	}
	r := strings.Split(getMetaValue(line), " ")

	return fmt.Sprintf("RE[%s+%s]", m[r[0]], r[1])
	// XXX: need to support more format, like resign, etc
}

func HandleLine(line string) {
	l := strings.Trim(line, " ")

	handlers := map[string]Handler{
		"\\HS":             handleHS,
		"STO":              handleSTO,
		"\\GE":             handleGE,
		"\\[GAMEWHITENAME": handleWhiteName,
		"\\[GAMEBLACKNAME": handleBlackName,
		"\\[GAMERESULT":    handleResult,
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
