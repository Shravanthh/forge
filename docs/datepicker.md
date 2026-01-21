# Date Picker

Interactive date picker component.

## Usage

```go
func FormPage(c *forge.Context) ui.UI {
    return ui.Form(
        ui.Label(ui.T("Birth Date")),
        ui.DatePicker("birthdate", c),
        ui.Button(ui.T("Submit")).OnClick(c, handleSubmit),
    )
}

func handleSubmit(c *forge.Context) {
    date := c.String("birthdate") // "2024-01-15"
    // Process date...
}
```

## Features

- Click input to open calendar dropdown
- Navigate months with arrows
- Click day to select
- Returns date as `YYYY-MM-DD` string
- Closes on selection

## API

```go
ui.DatePicker(id string, c *ctx.Context) Element
```

- `id` - State key for the selected date
- `c` - Context

Get selected value:

```go
date := c.String("birthdate") // "2024-01-15" or ""
```

## Example with Validation

```go
func BookingForm(c *forge.Context) ui.UI {
    return ui.Div(
        ui.Stack("12px",
            ui.Label(ui.T("Check-in Date")),
            ui.DatePicker("checkin", c),
            
            ui.Label(ui.T("Check-out Date")),
            ui.DatePicker("checkout", c),
        ),
        ui.Button(ui.T("Book")).OnClick(c, func(c *forge.Context) {
            checkin := c.String("checkin")
            checkout := c.String("checkout")
            
            v := ctx.NewValidator().
                Required("checkin", checkin, "Check-in date required").
                Required("checkout", checkout, "Check-out date required")
            
            if !v.Valid() {
                c.Set("errors", v.Errors())
                return
            }
            
            // Process booking...
        }),
    )
}
```

## Styling

Add date picker styles:

```go
ui.Style(ui.DatePickerStyles)
```

Or customize:

```css
.datepicker { position: relative; }
.datepicker-input { padding: 8px 12px; border: 1px solid #d1d5db; }
.datepicker-dropdown { position: absolute; background: white; }
.dp-day { width: 32px; height: 32px; }
.dp-day.selected { background: #3b82f6; color: white; }
```
