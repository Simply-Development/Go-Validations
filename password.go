/**
 * Copyright 2020 Simply Development
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package simplyvalidations

import (
	"fmt"
	"unicode"
)

// ValidatesPassword ... Ensure that password pass the security requirements
func ValidatesPassword(password string) error {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return fmt.Errorf("Password doesn't meet security requirements")
		}
	}

	/*
	* At least one upper character
	* At least one lower character
	* At least one numeric character
	* At least one special character
	* Eight characters as minimum length
	* */
	if !upp || !low || !num || !sym || tot < 8 {
		return fmt.Errorf("Password doesn't meet security requirements")
	}

	return nil
}
