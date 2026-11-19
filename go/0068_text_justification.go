//go:build ignore

// 68. Text Justification (Hard)
// https://leetcode.com/problems/text-justification/
// Run: go run 0068_text_justification.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	line, letters := []string{}, 0
	for _, w := range words {
		if letters+len(w)+len(line) > maxWidth { // w does not fit: justify the current line
			gaps := max(len(line)-1, 1)
			for i := 0; i < maxWidth-letters; i++ {
				line[i%gaps] += " " // round-robin puts extra spaces on the left
			}
			res = append(res, strings.Join(line, ""))
			line, letters = []string{}, 0
		}
		line = append(line, w)
		letters += len(w)
	}
	last := strings.Join(line, " ") // the last line is left-justified
	return append(res, last+strings.Repeat(" ", maxWidth-len(last)))
}

func main() {
	assertEq(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16), []string{
		"This    is    an", "example  of text", "justification.  "})
	assertEq(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16), []string{
		"What   must   be", "acknowledgment  ", "shall be        "})
	assertEq(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20), []string{
		"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
		"everything  else  we", "do                  "})
	assertEq(fullJustify([]string{"a"}, 1), []string{"a"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
