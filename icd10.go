// Package icd10
//
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
package icd10

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// NineToTen converts ICD9 codes to ICD10 when a conversion
// entry is present in the lookup table
func NineToTen(codes []string) ([]string, error) {
	var ret = []string{}

	m, _, err := setup()
	if err != nil {
		return ret, err
	}

	for _, code := range codes {
		icd10, ok := m[code]

		if !ok {
			ret = append(ret, fmt.Sprintf("UNKNOWN_ICD9:%s", code))
		} else {
			ret = append(ret, icd10)
		}
	}

	return ret, nil
}

// TenToNine converts ICD10 codes to ICD9 when a conversion
// entry is present in the lookup table
func TenToNine(codes []string) ([]string, error) {
	var ret = []string{}

	_, m, err := setup()
	if err != nil {
		return ret, err
	}

	for _, code := range codes {
		icd9, ok := m[code]

		if !ok {
			ret = append(ret, fmt.Sprintf("UNKNOWN_ICD10:%s", code))
		} else {
			ret = append(ret, icd9)
		}
	}

	return ret, nil
}

// setup reaads the conversion charts and returns them in a way that is easy to serve.
func setup() (map[string]string, map[string]string, error) {
	nine2ten := map[string]string{}
	ten2nine := map[string]string{}

	nines, err := ioutil.ReadFile("./resources/9-10.csv")
	if err != nil {
		return nine2ten, ten2nine, err
	}

	tens, err := ioutil.ReadFile("./resources/10-9.csv")
	if err != nil {
		return nine2ten, ten2nine, err
	}

	for _, line := range strings.Split(string(nines), "\n") {
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			continue
		}
		nine2ten[fields[0]] = fields[1]
	}

	for _, line := range strings.Split(string(tens), "\n") {
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			continue
		}
		ten2nine[fields[0]] = fields[1]
	}
	return nine2ten, ten2nine, err
}
