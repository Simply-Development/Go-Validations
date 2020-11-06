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

package simplyvalidationstests

import (
	simplyvalidations "simply-go-validations/src"
	"testing"
)

// TestValidatesFormat ... Test if validation works for Format requirement
func TestValidatesEmail(t *testing.T) {
	t.Run("Email is valid", func(t *testing.T) {
		got := simplyvalidations.ValidatesEmail("hello@simply-development.com")
		if got != nil {
			t.Errorf(`ValidatesEmail("hello@simply-development.com") = %s, want nil`, got)
		}
	})
	t.Run("Email is invalid", func(t *testing.T) {
		got := simplyvalidations.ValidatesEmail("hello@noexiste.com")
		if got == nil {
			t.Errorf(`ValidatesEmail("hello@noexiste.com") = %s, want error`, got)
		}
	})
}
