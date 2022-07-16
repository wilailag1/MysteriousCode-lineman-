// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
func toString6Bit(str string) string {
	chaAll := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	string6bit := ""
	for _, letter := range str {
		ab := fmt.Sprintf("%06b", strings.Index(chaAll, string(letter)))
		string6bit = string6bit + ab
	}
	return string6bit
}
func to8Bitlist(str6bit string) []string {
	var list8bit []string
	newstr8bit := ""
	for _, letter := range str6bit {
		newstr8bit = newstr8bit + string(letter)
		if len(newstr8bit) == 8 {
			list8bit = append(list8bit, newstr8bit)
			newstr8bit = ""
		}
	}
	return list8bit
}
func decodeAscii(list8bit []string) string {
	var textRes string
	for _, letter := range list8bit {
		asciiForm, _ := strconv.ParseInt(letter, 2, 64)
		character := rune(asciiForm)

		textRes = textRes + string(character)
	}
	return textRes
}

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	newSecret := strings.ReplaceAll(secret, "=", "A")
	string6bit := toString6Bit(newSecret)
	list8bit := to8Bitlist(string6bit)
	textDeCode := decodeAscii(list8bit)
	textRes := ReverseString(textDeCode)
	textRes = strings.ReplaceAll(textRes, ":", " ")
	fmt.Printf("Decoded textRes : %s\n", textRes)
	fmt.Printf("fullname : %s\n", "Veerapat Leatsilpa")
	fmt.Printf("phone : %s\n", "0919155764")
	fmt.Printf("email : %s\n", "weerapat.dev@gmail.com")
}
