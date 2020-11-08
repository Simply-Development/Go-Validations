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
	"testing"

	simplyvalidations "github.com/simply-development/go-validations"
)

// TestValidatesPassword ... Test if validation works for password entry
func TestValidatesPassword(t *testing.T) {
	t.Run("Password is valid", func(t *testing.T) {
		got := simplyvalidations.ValidatesPassword("contraseña")
		if got != nil {
			t.Errorf(`ValidatesPassword("contraseña") = %s, want nil`, got)
		}
	})
	t.Run("Password is invalid", func(t *testing.T) {
		got := simplyvalidations.ValidatesPassword("Contraseña-1")
		if got == nil {
			t.Errorf(`ValidatesPassword("Contraseña-1") = %s, want error`, got)
		}
	})
}
