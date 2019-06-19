# Go 语言笔记

## 你好，World！
```go
package main

import "fmt"

func main() {
    fmt.Println("你好，world！")
}
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

### 流程控制


