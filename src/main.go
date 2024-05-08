package main

import "fmt"

func main() {
    var str string = "testing 123"
    var a, b int = 4, 5
    fmt.Println("Hello!!!", str, "Numbers are:", a, "and", b)

    var c string
    fmt.Scanln(&c)
    fmt.Println(c)
}
