package util

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Run(fn func(string) int, filename string, expected int) {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	if idx := strings.LastIndex(fnName, "."); idx >= 0 {
		fnName = fnName[idx+1:]
	}

	start := time.Now()
	result := fn(filename)
	elapsed := time.Since(start)

	if result != expected {
		log.Fatalf("Failed for %s (%s): expected %d, got %d", fnName, filename, expected, result)
	} else {
		var timeString string
		if elapsed < time.Millisecond {
			timeString = fmt.Sprintf("%dµs", elapsed.Microseconds())
		} else if elapsed < time.Second {
			timeString = fmt.Sprintf("%dms", elapsed.Milliseconds())
		} else {
			timeString = fmt.Sprintf("%fs", elapsed.Seconds())
		}
		fmt.Printf("%s (%s): %d in %s\n", fnName, filename, result, timeString)
	}
}
