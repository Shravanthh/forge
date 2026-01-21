package ctx

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Validator validates form fields.
type Validator struct {
	errors map[string]string
}

// NewValidator creates a new validator.
func NewValidator() *Validator {
	return &Validator{errors: make(map[string]string)}
}

// Required checks if field is not empty.
func (v *Validator) Required(field, value, msg string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors[field] = msg
	}
	return v
}

// MinLen checks minimum length.
func (v *Validator) MinLen(field, value string, min int, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if len(value) < min {
		v.errors[field] = msg
	}
	return v
}

// MaxLen checks maximum length.
func (v *Validator) MaxLen(field, value string, max int, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if len(value) > max {
		v.errors[field] = msg
	}
	return v
}

// Email validates email format.
func (v *Validator) Email(field, value, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		v.errors[field] = msg
	}
	return v
}

// Match validates against regex pattern.
func (v *Validator) Match(field, value, pattern, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if matched, _ := regexp.MatchString(pattern, value); !matched {
		v.errors[field] = msg
	}
	return v
}

// Min checks minimum numeric value.
func (v *Validator) Min(field string, value, min int, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if value < min {
		v.errors[field] = msg
	}
	return v
}

// Max checks maximum numeric value.
func (v *Validator) Max(field string, value, max int, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if value > max {
		v.errors[field] = msg
	}
	return v
}

// Custom adds a custom validation.
func (v *Validator) Custom(field string, valid bool, msg string) *Validator {
	if _, ok := v.errors[field]; ok {
		return v
	}
	if !valid {
		v.errors[field] = msg
	}
	return v
}

// Valid returns true if no errors.
func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

// Errors returns all validation errors.
func (v *Validator) Errors() map[string]string {
	return v.errors
}

// Error returns error for a specific field.
func (v *Validator) Error(field string) string {
	return v.errors[field]
}

// ValidateJSON validates JSON data against a schema.
func ValidateJSON(data []byte, schema map[string]string) map[string]string {
	errors := make(map[string]string)
	var obj map[string]any
	
	if err := json.Unmarshal(data, &obj); err != nil {
		errors["_json"] = "Invalid JSON"
		return errors
	}

	for field, rules := range schema {
		value, exists := obj[field]
		ruleList := strings.Split(rules, "|")

		for _, rule := range ruleList {
			rule = strings.TrimSpace(rule)
			
			if rule == "required" && !exists {
				errors[field] = field + " is required"
				break
			}
			
			if !exists {
				continue
			}

			switch {
			case rule == "string":
				if _, ok := value.(string); !ok {
					errors[field] = field + " must be a string"
				}
			case rule == "number":
				if _, ok := value.(float64); !ok {
					errors[field] = field + " must be a number"
				}
			case rule == "bool":
				if _, ok := value.(bool); !ok {
					errors[field] = field + " must be a boolean"
				}
			case rule == "array":
				if _, ok := value.([]any); !ok {
					errors[field] = field + " must be an array"
				}
			case rule == "object":
				if _, ok := value.(map[string]any); !ok {
					errors[field] = field + " must be an object"
				}
			case rule == "email":
				if s, ok := value.(string); ok {
					emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
					if !emailRegex.MatchString(s) {
						errors[field] = field + " must be a valid email"
					}
				}
			}
		}
	}

	return errors
}
