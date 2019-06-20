package main

import (
    "fmt"
    "io/ioutil"
)

func greaterThan10(number int) bool {
    if number > 10 {
        return true
    } else {
        return false
    }
}

func switchByIf(number int) string {
    if number == 1 {
        return "red"
    } else if number == 2 {
        return "blue"
    } else if number == 3 {
        return "white"
    } else {
        return "black"
    }
}

func main() {
    fmt.Println(greaterThan10(11))
    fmt.Println(greaterThan10(9))

    fmt.Println(switchByIf(1))
    fmt.Println(switchByIf(2))
    fmt.Println(switchByIf(3))

    const filename = "haha.txt"

    // way1
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }

    // way2
    if contents, err := ioutil.ReadFile(filename); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }

}
