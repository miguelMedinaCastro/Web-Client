package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(label string) string {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

func ReadHeaders() map[string] string {
	headers := make(map[string]string)

	for {
		resp := ReadInput("adicionar header? (s/n): ")
		if strings.ToLower(resp) != "s" {
			break
		}
		
		key := ReadInput("header (a.g: Content-Type): ")
		val := ReadInput("valor: ")
		headers[key] = val
	}

	return headers
}

func ReadBody() string {
	fmt.Println("qual requisição (JSON/HTML). Finalize com linha vazia: ")

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
