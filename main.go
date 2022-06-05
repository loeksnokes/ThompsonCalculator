package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/loeksnokes/treepair"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	fmt.Println("Input tree alphabet (for calculations, it will receive UTF8 ordering): ")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// delete any \n or \r
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	tp, err := treepair.NewTreePairAlpha(text)

	if nil != err {
		outStr := "NewTreePairAlpha(): Failed to create domaintree from `" + text + "'."
		fmt.Println(outStr)
		return
	}
	tp.ExpandDomainAt(text)
	permutation := map[int]int{
		0:  1,
		1:  4,
		2:  2,
		3:  0,
		4:  3,
		5:  15,
		6:  14,
		7:  13,
		8:  12,
		9:  11,
		10: 10,
		11: 9,
		12: 8,
		13: 7,
		14: 6,
		15: 5,
	}
	tp.ApplyPermRange(permutation)

	fmt.Println(tp.FullString() + " is the first treePair.")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// delete any \n or \r
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}
}

// reversePreservingCombiningCharacters interprets its argument as UTF-8
// and ignores bytes that do not form valid UTF-8.  return value is UTF-8.
// Code from: https://rosettacode.org/wiki/Reverse_a_string#Go

func reversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}
