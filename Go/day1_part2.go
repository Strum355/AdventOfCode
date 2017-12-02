package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "1234"
    total := 0
    for i, curr := range s {
        if curr == rune(s[(i+len(s)/2)%len(s)]) {
            num, _ := strconv.Atoi(string(curr))
            total += num
        }
    }
    fmt.Println(total)
}