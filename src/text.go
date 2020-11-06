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
	"regexp"
)

// TextValidationOptions ... Define the requirements that text has to pass
type TextValidationOptions struct {
	Required  bool
	MinLength int
	MaxLength int
	Format    string
}

// OnlyLettersRegex ... Regex for text with only letters
const OnlyLettersRegex = "^[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$"

// ValidatesText ... Validates that text have certain format
func ValidatesText(text string, requirements TextValidationOptions) error {
	if requirements.Required && text == "" {
		return fmt.Errorf("is required")
	}
	if len(text) < requirements.MinLength && text != "" {
		return fmt.Errorf("should have at least %d characters", requirements.MinLength)
	}
	if requirements.MaxLength > 0 {
		if len(text) > requirements.MaxLength {
			return fmt.Errorf("should have %d characters as maximum", requirements.MaxLength)
		}
	}
	if requirements.Format != "" {
		validFormat, err := regexp.Match(requirements.Format, []byte(text))

		if !validFormat || err != nil {
			return fmt.Errorf("should match with %s", requirements.Format)
		}
	}

	return nil
}
