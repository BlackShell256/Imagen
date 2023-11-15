package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) < 2 {
		panic("faltan argumentos")
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	bytes := fmt.Sprintf("[]byte{%s}", ConvertToString(file))

	clipboard.WriteAll(bytes)

}

func ConvertToString(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}
