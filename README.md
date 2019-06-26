# Go 语言笔记

## 你好，World！
```go
package main

import "fmt"

func main() {
    fmt.Println("你好，world！")
}
```

## 函数
函数定义语法
```go
func 函数名(变量1, 变量2 变量类型1, 变量3 变量类型2, 变量4 变量类型4) 函数返回值类型1 {
    // 一些操作
}

func 函数名(变量1, 变量2 变量类型1, 变量3 变量类型2) (函数返回值类型1, 函数返回值类型2) {
    // 一些操作
}

func 函数名(变量1, 变量2 变量类型1) (别名1 函数返回值类型1, 别名2 函数返回值类型2) {
    // 一些操作
}
```

示例代码
```go
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


func add(a int, b int) int {
    return a + b
}


// function can be a parameter
// e.g. result = apply(add, 1, 2)
// result equals 3
func apply(op func(int, int) int, a, b int) int {
    return op(a, b)
}
```

**匿名函数**
```go
result := apply(func(a int, b int) int {
    return b - a
}, 3, 2)
fmt.Println(result)
```

**可变参数列表**，传入...代表可以有任意个数的参数
```go
func sum(numbers ...int) int {
    var result int = 0
    for i := range numbers {
        result += numbers[i]
    }
    return result
}

fmt.Println(sum(1, 2, 3))   // 6
```

## 基础语法

### 变量与声明变量
#### 声明
使用`var`定义变量
* 定义了变量后就有初值
* 定义后必须使用
```go
var 变量名 变量类型
```

示例：
```go
func zeroValue() {
    // a = 0, s = ""
    var a, b int
    var s string
    fmt.Println("1. ", a, b, s)
    fmt.Printf("2. a = %d, b = %d, s = %s\n", a, b, s)
    // %q print empty string as ""
    fmt.Printf("3. a = %d, b = %d, s = %q\n", a, b, s)
}

func initialValue() {
    var a, b int = 1, 2
    var s string = "hello"

    fmt.Println(a, b, s)
}

func auto() {
    var a, b, bool_, s = 1, 2, true, "hello"
    var (
        c int = 3
        s2 string = "world"
    )
    fmt.Println(a, b, bool_,s)
    fmt.Println(c, s2)
}

func shortcut() {
    a, b, bool_, s := 1, 2, true, "hello"
    fmt.Println(a, b, bool_,s)
}
```

#### 内建变量类型
* 整型
    * (u)int：根据操作系统自动定义
    * (u)int8, (u)int32, (u)int64：8位、32位、64位
    * uintptr：指针
* 浮点型
    * float32， float64
    * complex64， complex128
* bool
* 字符类型
    * string：字符串类型
    * byte
    * rune: 字符型（与C的char很像，但是4字节）

#### 强制类型转换
Go语言只有强制类型转换，没有隐式类型转换。float32和float64也是不同的类型。
```go
func cast (a float64) int{
    return int(a)
}

func main() {
    var a float32 = 2.99
    b := cast(float64(a))
    fmt.Println(b)  // 2
}
```

#### 常量
常量不规定类型时，其类型在运算时才确定，所以可以作为各种类型来使用（类似C语言的#define）
```go
func const_() {
    const a, b = 1, 2
    const c, d float64 = 1.0, 2.0
    fmt.Println(a + b)      // 3
    fmt.Println(c + d)      // 3
    fmt.Println(a + 3.14)   // 4.14
}
```

我们可以使用const来定义一个枚举类型
```go
func enum() {
    const (
        red   = 0
        blue  = 1
        green = 2
        white = 3
    )

    fmt.Println(red, blue, green, white)
}
```

使用`iota`关键字可以自增值
```go
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
```

`iota`关键字可以参与运算
```go
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
```


### 流程控制
#### if
```go
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
```

#### for
```go
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
```

## License
[Apache License 2.0](./LICENSE)


