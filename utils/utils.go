package utils

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadContent(filename string) string {
	_, callerFilename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("unable to get caller's file path")
	}

	absolutePath := path.Join(path.Dir(callerFilename), filename)
	file, err := os.Open(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := content.String()
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}

	return result
}

func ReadLines(filename string) []string {
	content := ReadContent(filename)

	var lines []string
	for _, line := range strings.Split(content, "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		}
	}

	return lines
}
