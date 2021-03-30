package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iz4vve/go-icd10"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ten2nine <code1> <code2> <code3> ...")
		os.Exit(1)
	}

	codes := os.Args[1:]
	ret, err := icd10.TenToNine(codes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(ret, " "))
}
