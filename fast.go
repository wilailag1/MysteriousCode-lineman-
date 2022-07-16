// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func ReverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func main() {
	rawDecodedText, err := base64.StdEncoding.DecodeString("aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K")
	if err != nil {
		panic(err)
	}
	res := ReverseString(string(rawDecodedText))
	res = strings.ReplaceAll(res, ":", " ")
	fmt.Printf("Decoded text: %s\n", res)
	fmt.Printf("fullname : %s\n", "Veerapat Leatsilpa")
	fmt.Printf("phone : %s\n", "0919155764")
	fmt.Printf("email : %s\n", "weerapat.dev@gmail.com")
}
