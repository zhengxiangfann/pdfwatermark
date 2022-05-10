package main

import "C"
import "fmt"

//export go_fib
func go_fib(n C.int) C.double {
    var i C.int = 0
    var a, b C.double = 0.0, 1.0
    for ; i < n; i++ {
        a, b = a + b, a
    }
    fmt.Println("斐波那契计算完毕，我是 Go 语言")
    return a
}

func main() {}
