# Toast Notifications

Display toast messages for user feedback.

## Setup

Add the toast container to your layout:

```go
func Layout(c *forge.Context, content ui.UI) ui.UI {
    return ui.Div(
        ui.Toast(c),  // Toast container
        content,
    )
}
```

Add toast styles:

```go
ui.Style(ui.ToastStyles)
```

## Usage

```go
// Success toast
ui.ToastSuccess(c, "Changes saved!")

// Error toast
ui.ToastError(c, "Something went wrong")

// Info toast
ui.ToastInfo(c, "New message received")

// Warning toast
ui.ToastWarning(c, "Session expiring soon")
```

## Example

```go
func SaveHandler(c *forge.Context) {
    err := saveData(c)
    if err != nil {
        ui.ToastError(c, "Failed to save: " + err.Error())
        return
    }
    ui.ToastSuccess(c, "Saved successfully!")
}

func Page(c *forge.Context) ui.UI {
    return ui.Div(
        ui.Toast(c),
        ui.Button(ui.T("Save")).OnClick(c, SaveHandler),
    )
}
```

## Custom Variant

```go
ui.ShowToast(c, "Custom message", "custom-variant")
```

Then add CSS for `.toast-custom-variant`.

## Styling

Default styles included. Override with custom CSS:

```css
.toast-container { position: fixed; top: 20px; right: 20px; }
.toast { padding: 16px; border-radius: 8px; }
.toast-success { background: #10b981; }
.toast-error { background: #ef4444; }
.toast-info { background: #3b82f6; }
.toast-warning { background: #f59e0b; }
```
