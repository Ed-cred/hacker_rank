package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
    // Write your code here
	var ret string
	for _, r := range s {
		if unicode.IsLetter(r) {
			switch {
			case unicode.IsUpper(r):
				if r + k > 90 {
					r = 65 + (r + k - 90)
					ret += string(r)
				}
			case unicode.IsLower(r):
				if r + k > 172 {
					r = 97 + (r + k - 172)
					ret += string(r)
				}
			}
		} 
		ret += string(r)
	}
	return ret
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    // nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    // checkError(err)
    // n := int32(nTemp)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
