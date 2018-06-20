package env

import (
	"os"
	"strings"
)

var FLAG int

const (
	FIRSTRUN = 1
	COOKIE = 2
	HEADER = 4
)

func Load() {
	FLAG = 0

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "FIRSTRUN" && pair[1] == "1" {
			FLAG |= FIRSTRUN
		}
	}
}