package main

import "fmt"

func cast (a float64) int{
    return int(a)
}

func main() {
    var a float32 = 2.99
    b := cast(float64(a))
    fmt.Println(b)  // 2
}
