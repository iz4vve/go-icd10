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
package icd10_test

import (
	"testing"

	"github.com/iz4vve/go-icd10"
)

func TestNineToTen(t *testing.T) {

}

func TestTenToNine(t *testing.T) {

}

func BenchmarkNineToTen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		icd10.NineToTen([]string{"010", "011"})
	}
}

func BenchmarkTenToNine(b *testing.B) { // run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		icd10.TenToNine([]string{"A10", "A99"})
	}
}
