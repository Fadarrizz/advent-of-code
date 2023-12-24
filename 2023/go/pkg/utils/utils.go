package utils

import (
	"fmt"
	"os"
)

// Allow unused variables to be included
func Use(x ...interface{}) {}

func Dump(x ...interface{}) {
    fmt.Printf("%+v\n", x)
}

func DD(x ...interface{}) {
    Dump(x)
    os.Exit(0)
}
