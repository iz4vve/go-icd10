package main

import (
	"fmt"

	"github.com/iz4vve/go-icd10"
)

func main() {
	fmt.Println("10 2 9!")
	fmt.Println(icd10.TenToNine([]string{"A000"}))
}
