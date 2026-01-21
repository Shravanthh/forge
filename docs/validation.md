# Validation

Validate forms and JSON data.

## Form Validation

### Basic Usage

```go
func handleSubmit(c *forge.Context) {
    email := c.String("email")
    password := c.String("password")
    
    v := ctx.NewValidator().
        Required("email", email, "Email is required").
        Email("email", email, "Invalid email format").
        Required("password", password, "Password is required").
        MinLen("password", password, 8, "Password must be at least 8 characters")
    
    if !v.Valid() {
        c.Set("errors", v.Errors())
        return
    }
    
    // Process valid form...
}
```

### Display Errors

```go
func FormPage(c *forge.Context) ui.UI {
    errors := c.Get("errors").(map[string]string)
    
    return ui.Form(
        ui.Stack("4px",
            ui.Label(ui.T("Email")),
            ui.Input().WithID("email").OnInput(c, ...),
            errorText(errors["email"]),
        ),
        // ...
    )
}

func errorText(err string) ui.UI {
    if err == "" {
        return ui.Span()
    }
    return ui.Span(ui.T(err)).WithClass("text-red-500 text-sm")
}
```

### Available Validators

```go
v := ctx.NewValidator()

// Required field
v.Required("field", value, "Field is required")

// String length
v.MinLen("field", value, 3, "Min 3 characters")
v.MaxLen("field", value, 100, "Max 100 characters")

// Email format
v.Email("field", value, "Invalid email")

// Regex pattern
v.Match("field", value, `^\d{5}$`, "Must be 5 digits")

// Numeric range
v.Min("field", intValue, 0, "Must be positive")
v.Max("field", intValue, 100, "Max 100")

// Custom validation
v.Custom("field", value == expected, "Values must match")
```

### Chaining

All validators return `*Validator` for chaining:

```go
v := ctx.NewValidator().
    Required("name", name, "Required").
    MinLen("name", name, 2, "Too short").
    MaxLen("name", name, 50, "Too long").
    Required("email", email, "Required").
    Email("email", email, "Invalid")
```

### Check Results

```go
if v.Valid() {
    // All validations passed
}

// Get all errors
errors := v.Errors() // map[string]string

// Get specific error
emailErr := v.Error("email")
```

## JSON Validation

Validate JSON data against a schema:

```go
data := []byte(`{"email": "test@example.com", "age": 25}`)

errors := ctx.ValidateJSON(data, map[string]string{
    "email": "required|string|email",
    "age":   "required|number",
    "name":  "string",
})

if len(errors) > 0 {
    // Handle validation errors
}
```

### Schema Rules

| Rule | Description |
|------|-------------|
| `required` | Field must exist |
| `string` | Must be a string |
| `number` | Must be a number |
| `bool` | Must be a boolean |
| `array` | Must be an array |
| `object` | Must be an object |
| `email` | Must be valid email |

Combine rules with `|`:

```go
schema := map[string]string{
    "email":    "required|string|email",
    "age":      "required|number",
    "tags":     "array",
    "metadata": "object",
    "active":   "bool",
}
```
