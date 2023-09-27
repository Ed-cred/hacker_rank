package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int) string {
	alphabetLower:= "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper:= "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var ret string
	for _, r := range s {
		switch {
		case strings.ContainsRune(alphabetLower, r):
			ret = ret + string(rotate(r, k, []rune(alphabetLower))) 
		case strings.ContainsRune(alphabetUpper, r):
			ret = ret + string(rotate(r, k, []rune(alphabetUpper))) 
		default:
			ret = ret + string(r)
		}
	}
	return ret
}
func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic ("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)


    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int(kTemp)

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
