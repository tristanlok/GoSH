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

    // Beginning of Loop
    for true {
        fmt.Print("> ")
        scan.ScanLine(scanner, &line)
        scan.GetArgs(&line)
    }
}
