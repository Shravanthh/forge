# Forms

Building forms in Forge.

## Basic Form

```go
func ContactForm(c *forge.Context) ui.UI {
    return ui.Form(
        ui.Stack("16px",
            ui.Stack("4px",
                ui.Label(ui.T("Name")),
                ui.Input().
                    WithID("name").
                    WithAttr("type", "text").
                    WithAttr("value", c.String("name")).
                    WithClass("input").
                    OnInput(c, func(c *forge.Context) {
                        c.Set("name", c.InputValue())
                    }),
            ),
            ui.Stack("4px",
                ui.Label(ui.T("Email")),
                ui.Input().
                    WithID("email").
                    WithAttr("type", "email").
                    WithAttr("value", c.String("email")).
                    WithClass("input").
                    OnInput(c, func(c *forge.Context) {
                        c.Set("email", c.InputValue())
                    }),
            ),
            ui.Button(ui.T("Submit")).
                WithClass("btn btn-primary").
                WithAttr("type", "submit"),
        ),
    ).WithID("contact-form").OnSubmit(c, func(c *forge.Context) {
        name := c.String("name")
        email := c.String("email")
        // Process form...
        
        // Clear form
        c.Set("name", "")
        c.Set("email", "")
        c.Set("submitted", true)
    })
}
```

## Input Types

```go
// Text
ui.Input().WithAttr("type", "text")

// Email
ui.Input().WithAttr("type", "email")

// Password
ui.Input().WithAttr("type", "password")

// Number
ui.Input().WithAttr("type", "number")

// Date
ui.Input().WithAttr("type", "date")

// Checkbox
ui.Input().WithAttr("type", "checkbox")

// Radio
ui.Input().WithAttr("type", "radio").WithAttr("name", "group")
```

## Textarea

```go
ui.El("textarea",
    ui.T(c.String("message")),
).WithID("message").
  WithAttr("rows", "4").
  OnInput(c, func(c *forge.Context) {
      c.Set("message", c.InputValue())
  })
```

## Select

```go
ui.El("select",
    ui.El("option", ui.T("Select...")).WithAttr("value", ""),
    ui.El("option", ui.T("Option 1")).WithAttr("value", "1"),
    ui.El("option", ui.T("Option 2")).WithAttr("value", "2"),
    ui.El("option", ui.T("Option 3")).WithAttr("value", "3"),
).WithID("choice").OnChange(c, func(c *forge.Context) {
    c.Set("choice", c.InputValue())
})
```

## Checkbox

```go
func Checkbox(c *forge.Context, id, label string) ui.UI {
    checked := c.Bool(id)
    
    checkbox := ui.Input().
        WithID(id).
        WithAttr("type", "checkbox").
        OnClick(c, func(c *forge.Context) {
            c.Set(id, !c.Bool(id))
        })
    
    if checked {
        checkbox = checkbox.WithAttr("checked", "checked")
    }
    
    return ui.Label(
        checkbox,
        ui.Span(ui.T(" " + label)),
    )
}

// Usage
Checkbox(c, "agree", "I agree to the terms")
```

## Radio Group

```go
func RadioGroup(c *forge.Context, name string, options []string) ui.UI {
    selected := c.String(name)
    var items []ui.UI
    
    for _, opt := range options {
        opt := opt
        radio := ui.Input().
            WithID(name + "_" + opt).
            WithAttr("type", "radio").
            WithAttr("name", name).
            OnClick(c, func(c *forge.Context) {
                c.Set(name, opt)
            })
        
        if selected == opt {
            radio = radio.WithAttr("checked", "checked")
        }
        
        items = append(items, ui.Label(
            radio,
            ui.Span(ui.T(" " + opt)),
        ))
    }
    
    return ui.Stack("8px", items...)
}

// Usage
RadioGroup(c, "size", []string{"Small", "Medium", "Large"})
```

## Form Validation

```go
func validateEmail(email string) string {
    if email == "" {
        return "Email is required"
    }
    if !strings.Contains(email, "@") {
        return "Invalid email format"
    }
    return ""
}

func Form(c *forge.Context) ui.UI {
    emailError := c.String("email_error")
    
    return ui.Form(
        ui.Stack("4px",
            ui.Label(ui.T("Email")),
            ui.Input().
                WithID("email").
                WithAttr("type", "email").
                OnInput(c, func(c *forge.Context) {
                    email := c.InputValue()
                    c.Set("email", email)
                    c.Set("email_error", validateEmail(email))
                }),
            // Show error
            func() ui.UI {
                if emailError != "" {
                    return ui.Span(ui.T(emailError)).
                        WithClass("text-red-500 text-sm")
                }
                return ui.Span()
            }(),
        ),
        ui.Button(ui.T("Submit")).
            WithAttr("type", "submit"),
    ).OnSubmit(c, func(c *forge.Context) {
        email := c.String("email")
        if err := validateEmail(email); err != "" {
            c.Set("email_error", err)
            return
        }
        // Submit form...
    })
}
```

## Form Styling with Tailwind

```go
ui.Input().
    WithClass("w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500")

ui.Button(ui.T("Submit")).
    WithClass("px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition")
```
