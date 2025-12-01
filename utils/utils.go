package utils

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadLines(filename string) []string {
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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
