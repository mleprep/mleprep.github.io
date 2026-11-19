//go:build ignore

// 535. Encode and Decode TinyURL (Medium)
// https://leetcode.com/problems/encode-and-decode-tinyurl/
// Run: go run 0535_encode_and_decode_tinyurl.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

type Codec struct {
	longToShort, shortToLong map[string]string
}

func Constructor() Codec {
	return Codec{longToShort: map[string]string{}, shortToLong: map[string]string{}}
}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62(n int) string {
	s := ""
	for {
		s = string(alphabet[n%62]) + s
		n /= 62
		if n == 0 {
			return s
		}
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	code, ok := this.longToShort[longUrl]
	if !ok { // a counter in base 62: short, unique, no collisions
		code = base62(len(this.shortToLong) + 1)
		this.longToShort[longUrl] = code
		this.shortToLong[code] = longUrl
	}
	return "http://tinyurl.com/" + code
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.shortToLong[shortUrl[strings.LastIndex(shortUrl, "/")+1:]]
}

func main() {
	c := Constructor()
	u := "https://leetcode.com/problems/design-tinyurl"
	assertEq(c.decode(c.encode(u)), u)
	assertEq(c.encode(u) == c.encode(u), true)
	for _, u := range []string{"https://example.com/a?b=c#d", "x"} {
		assertEq(c.decode(c.encode(u)), u)
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
