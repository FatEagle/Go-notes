package main

import "fmt"

func my_print(s string) {
    fmt.Println(s)
}

func add_suffix(s string) string {
    s = s + " suffix"
    return s
}

func calculate(a, b int, operation string) int {
    switch operation {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        q, _ := div(a, b)
        return q
    default:
        panic("unsupported operation: " + operation)
    }
}

func div(a, b int) (q, r int) {
    return a / b, a % b
}

func div2(a, b int) (q, r int) {
    q = a / b
    r = a % b
    return
}

func div3(a, b int) (int, int) {
    q := a / b
    r := a % b
    return q, r
}

func main() {
    my_print("hello world")
    s := add_suffix("hello world")
    fmt.Println(s)
    fmt.Println(calculate(1, 2, "*"))
    fmt.Println(calculate(7, 2, "/"))

    q, r := div(11, 5)
    fmt.Println(q, r)

    q, r = div2(11, 5)
    fmt.Println(q, r)

    q, r = div3(11, 5)
    fmt.Println(q, r)
}
