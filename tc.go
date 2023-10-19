package main

import (
	"github.com/KM911/oslib"
	"github.com/KM911/oslib/adt"
	"os"
	"strings"
)

func main() {
	defer adt.TimerStart().End()
	oslib.RunStd(strings.Join(os.Args[1:], " "))
}
