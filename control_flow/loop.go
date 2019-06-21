package main

import (
    "fmt"
    "strconv"
)

func forStatement() {
    sum := 0
    for i := 0; i < 100; i++ {
        sum += 1
    }
    fmt.Println(sum)

    sum = 10
    s := ""
    for ; sum > 5; sum-- {
        if s == "" {
            s = strconv.Itoa(sum)
        } else {
            s = s + " " + strconv.Itoa(sum)
        }
    }
    fmt.Println(s)

    sum = 10
    for sum > 0 {
        sum -= 1
    }
    fmt.Println(sum)

    sum = 10
    for {
        if sum < 0 {
            break
        }
        sum -= 1
    }
    fmt.Println(sum)
}


func main() {
    forStatement()
}
