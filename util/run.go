package util

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func run(fn func(string) int, filename string) (fnName string, result int, elapsed string) {
	fnName = runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	if idx := strings.LastIndex(fnName, "."); idx >= 0 {
		fnName = fnName[idx+1:]
	}

	start := time.Now()
	result = fn(filename)
	duration := time.Since(start)

	if duration < time.Millisecond {
		elapsed = fmt.Sprintf("%dµs", duration.Microseconds())
	} else if duration < time.Second {
		elapsed = fmt.Sprintf("%dms", duration.Milliseconds())
	} else {
		elapsed = fmt.Sprintf("%fs", duration.Seconds())
	}

	return fnName, result, elapsed
}

func Run(fn func(string) int, filename string) {
	fnName, result, elapsed := run(fn, filename)
	fmt.Printf("%s (%s): %d in %s\n", fnName, filename, result, elapsed)
}

func RunAndCheck(fn func(string) int, filename string, expected int) {
	fnName, result, elapsed := run(fn, filename)
	if result != expected {
		log.Fatalf("Failed for %s (%s): expected %d, got %d", fnName, filename, expected, result)
	} else {
		fmt.Printf("%s (%s): %d in %s\n", fnName, filename, result, elapsed)
	}
}
