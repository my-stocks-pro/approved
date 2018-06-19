package env

import (
	"os"
	"strings"
)

var ENV int

const (
	FIRSTRUN = 1
)

func Load() {
	ENV = 0

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "FIRSTRUN" && pair[1] == "1" {
			ENV |= FIRSTRUN
		}
	}
}