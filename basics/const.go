package main

import "fmt"

func const_() {
    const a, b = 1, 2
    const c, d float64 = 1.0, 2.0
    fmt.Println(a + b)    // 3
    fmt.Println(c + d)    // 3
    fmt.Println(a + 3.14) // 4.14
}

func enum() {
    const (
        red   = 0
        blue  = 1
        green = 2
        white = 3
    )

    fmt.Println(red, blue, green, white)
}

func enum_auto() {
    // auto add 1
    const (
        red   = iota
        blue
        green
        white
    )

    fmt.Println(red, blue, green, white)
}

func enmu_unit() {
    // we can calculate with iota
    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
    )

    // 1 1024 1048576 1073741824
    fmt.Println(b, kb, mb, gb)
}

func main() {
    const_()
    enum()
    enum_auto()
    enmu_unit()
}
