package main

import "fmt"

func test(x int, y int) (t int) {
    t = x + y
    return
}

func main() {
    var a, b int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Println(test(a, b))
}
