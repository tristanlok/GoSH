package scan

import (
	"bufio"
	"fmt"
)

func ScanLine (scanner *bufio.Scanner, line *string) {
    scanner.Scan()
    *line = scanner.Text()
}

func GetArgs (line *string) {
    for i := 0; i < len(*line); i++ {
        fmt.Println((*line)[i])
    }
}
