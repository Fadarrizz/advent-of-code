package conv

import (
	"strconv"
)

func StrToInt(s string) int {
    int, err := strconv.Atoi(s)
    if err != nil {
        panic(err.Error())
    }

    return int
}

func StrsToInts(s []string) []int {
    ints := make([]int, len(s))
    for i, str := range s {
        ints[i] = StrToInt(str)
    }

    return ints
}

func RuneToInt(r rune) int {
    return StrToInt(string(r))
}
