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

// TestValidatesTextRequired ... Test if validation works for Required requirement
func TestValidatesTextRequired(t *testing.T) {
	t.Run("Text is required", func(t *testing.T) {
		t.Run("Text have value", func(t *testing.T) {
			got := simplyvalidations.ValidatesText("Test", simplyvalidations.TextValidationOptions{Required: true})
			if got != nil {
				t.Errorf(`ValidatesText("", simplyvalidations.TextValidationOptions{Required: false}) = %s, want nil`, got)
			}
		})
		t.Run("Text is an empty string", func(t *testing.T) {
			got := simplyvalidations.ValidatesText("", simplyvalidations.TextValidationOptions{Required: true})
			if got == nil {
				t.Errorf(`ValidatesText("", simplyvalidations.TextValidationOptions{Required: false}) = %s, want error`, got)
			}
		})
	})
	t.Run("Text is not required", func(t *testing.T) {
		t.Run("Text have value", func(t *testing.T) {
			got := simplyvalidations.ValidatesText("Test", simplyvalidations.TextValidationOptions{Required: false})
			if got != nil {
				t.Errorf(`ValidatesText("", simplyvalidations.TextValidationOptions{Required: false} = %s, want nil`, got)
			}
		})
		t.Run("Text is an empty string", func(t *testing.T) {
			got := simplyvalidations.ValidatesText("", simplyvalidations.TextValidationOptions{Required: false})
			if got != nil {
				t.Errorf(`ValidatesText("", simplyvalidations.TextValidationOptions{Required: false} = %s, want nil`, got)
			}
		})
	})
}

// TestValidatesTextMinLength ... Test if validation works for MinLength requirement
func TestValidatesTextMinLength(t *testing.T) {
	t.Run("Text have more characters than minimum length", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("Test", simplyvalidations.TextValidationOptions{MinLength: 3})
		if got != nil {
			t.Errorf(`ValidatesText("Test", simplyvalidations.TextValidationOptions{MinLength: 3}) = %s, want nil`, got)
		}
	})
	t.Run("Text have less characters than minimum length", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("U", simplyvalidations.TextValidationOptions{MinLength: 3})
		if got == nil {
			t.Errorf(`ValidatesText("U, simplyvalidations.TextValidationOptions{MinLength: 3}) = %s, want error`, got)
		}
	})
}

// TestValidatesTextMaxLength ... Test if validation works for MaxLength requirement
func TestValidatesTextMaxLength(t *testing.T) {
	t.Run("Text have less characters than maximum length", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("Test", simplyvalidations.TextValidationOptions{MaxLength: 5})
		if got != nil {
			t.Errorf(`ValidatesText("Test", simplyvalidations.TextValidationOptions{MaxLength: 5}) = %s, want nil`, got)
		}
	})
	t.Run("Text have more characters than maximum length", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("Test with more characters", simplyvalidations.TextValidationOptions{MaxLength: 5})
		if got == nil {
			t.Errorf(`ValidatesText("Test with more characters", simplyvalidations.TextValidationOptions{MaxLength: 5}) = %s, want error`, got)
		}
	})
}

// TestValidatesTextFormat ... Test if validation works for Format requirement
func TestValidatesTextFormat(t *testing.T) {
	t.Run("Text match with required format", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("Test", simplyvalidations.TextValidationOptions{Format: simplyvalidations.OnlyLettersRegex})
		if got != nil {
			t.Errorf(`ValidatesText("Test", simplyvalidations.TextValidationOptions{Format: simplyvalidations.OnlyLettersRegex}) = %s, want nil`, got)
		}
	})
	t.Run("Text doesn't match with required format", func(t *testing.T) {
		got := simplyvalidations.ValidatesText("Test1", simplyvalidations.TextValidationOptions{Format: simplyvalidations.OnlyLettersRegex})
		if got == nil {
			t.Errorf(`ValidatesText("Test1", simplyvalidations.TextValidationOptions{Format: simplyvalidations.OnlyLettersRegex}) = %s, want error`, got)
		}
	})
}
