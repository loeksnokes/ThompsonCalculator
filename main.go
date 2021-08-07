package main

import (
	"fmt"

	"github.com/loeksnokes/prefcode"
	"github.com/loeksnokes/treepair"
)

func main() {
	tp, _ := treepair.NewTreePairAlpha("01")
	baseCodeToExpand, _ := prefcode.NewPrefCodeAlphaString("日本語")
	fmt.Println("Got:  " + tp.FullString())
	fmt.Println("baseCodeToExpand: " + baseCodeToExpand.String())
}
