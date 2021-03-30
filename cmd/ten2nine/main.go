// Copyright 2021 Pietro Mascolo

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
