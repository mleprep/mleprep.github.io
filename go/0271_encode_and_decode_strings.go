//go:build ignore

// 271. Encode and Decode Strings (Medium)
// https://leetcode.com/problems/encode-and-decode-strings/
// Run: go run 0271_encode_and_decode_strings.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Encode writes each string as "<length>#<string>".
func (codec *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)) + "#" + s)
	}
	return sb.String()
}

func (codec *Codec) Decode(s string) []string {
	res := []string{}
	for i := 0; i < len(s); {
		j := i + strings.IndexByte(s[i:], '#')
		n, _ := strconv.Atoi(s[i:j])
		res = append(res, s[j+1:j+1+n])
		i = j + 1 + n
	}
	return res
}

func main() {
	c := Constructor()
	for _, data := range [][]string{{"neet", "code", "love", "you"}, {"we", "say", ":", "yes", "#4#", ""}, {}} {
		assertEq(c.Decode(c.Encode(data)), data)
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
