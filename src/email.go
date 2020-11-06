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
	"net"
	"regexp"
	"strings"
)

// EmailRegex ... Regex for email
const EmailRegex = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// ValidatesEmail ... Validates that email is valid
func ValidatesEmail(email string) error {
	validFormat, err := regexp.Match(EmailRegex, []byte(email))

	if !validFormat || err != nil {
		return fmt.Errorf("should be a valid email")
	}
	emailParts := strings.Split(email, "@")
	mx, err := net.LookupMX(emailParts[1])
	if err != nil || len(mx) == 0 {
		return fmt.Errorf("should be a valid email")
	}

	return nil
}
