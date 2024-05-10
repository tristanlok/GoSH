package scan

import (
	"bufio"
	"strings"
)

func ScanLine (scanner *bufio.Scanner, line *string) {
    scanner.Scan()
    *line = scanner.Text()
}

// split this into a list
func GetArgs (line *string) (strptr *[]string) {
    // Splits the string delimited by space (does not matter how many)
    var strList []string = strings.Fields(*line)
    strptr = &strList
    return
}
