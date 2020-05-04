/*
Description:
Encrypt this!

You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter needs to be converted to its ASCII code.
The second letter needs to be switched with the last letter
Keepin' it simple: There are no special characters in input.
Examples:
EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"
*/

package kata

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return text
	}
	w := strings.Split(text, " ")
	fin := ""
	for _, v := range w {
		fin += enc(v) + " "
	}
	return strings.Trim(fin, " ")
}

func enc(s string) string {
	enc := ""
	switch {
	case len(s) == 1:
		asciiValue := int(s[0])
		enc += fmt.Sprintf("%d", asciiValue)
	case len(s) == 2:
		asciiValue := int(s[0])
		enc += fmt.Sprintf("%d%s", asciiValue, string(s[len(s)-1]))
	case len(s) == 3:
		asciiValue := int(s[0])
		enc += fmt.Sprintf("%d%s%s", asciiValue, string(s[len(s)-1]), string(s[1]))
	case len(s) > 3:
		asciiValue := int(s[0])
		enc += fmt.Sprintf("%d%s%s%s", asciiValue, string(s[len(s)-1]), string(s[2:len(s)-1]), string(s[1]))
	}
	return enc
}
