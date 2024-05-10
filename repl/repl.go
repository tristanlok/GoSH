package repl

import (
	"bufio"
	"fmt"
	"gosh/scanner"
	"os"
)

func Start () {
    // Initiate variable and scanner
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var line string
    var strList *[]string

    // Beginning of Loop
    for true {
        fmt.Print("> ")
        scan.ScanLine(scanner, &line)
        strList = scan.GetArgs(&line)

        fmt.Printf("%q\n", *strList)
    }
}
