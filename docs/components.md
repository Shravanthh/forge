# Components

Forge includes pre-built components for common UI patterns.

## Layout Components

### Stack (Vertical)

```go
ui.Stack("16px",
    ui.T("Item 1"),
    ui.T("Item 2"),
    ui.T("Item 3"),
)
```

### Row (Horizontal)

```go
ui.Row("8px",
    ui.Button(ui.T("Save")),
    ui.Button(ui.T("Cancel")),
)
```

### Card

```go
ui.Card(
    ui.H3(ui.T("Card Title")),
    ui.P(ui.T("Card content goes here.")),
)
```

### Divider

```go
ui.Divider()
```

## Feedback Components

### Alert

```go
ui.Alert("Operation successful!", "success")
ui.Alert("Please check your input.", "warning")
ui.Alert("An error occurred.", "error")
ui.Alert("Here's some information.", "info")
```

### Badge

```go
ui.Badge("New")
ui.Badge("Featured")
ui.Badge("Sale")
```

### Spinner

```go
ui.Spinner()
```

### Progress

```go
ui.Progress(25)   // 25%
ui.Progress(50)   // 50%
ui.Progress(100)  // 100%
```

## Data Components

### Table

```go
ui.Table(
    []string{"Name", "Email", "Role"},
    [][]ui.UI{
        {ui.T("John"), ui.T("john@example.com"), ui.Badge("Admin")},
        {ui.T("Jane"), ui.T("jane@example.com"), ui.Badge("User")},
    },
)
```

## Interactive Components

### Modal

```go
// Trigger button
ui.Button(ui.T("Open Modal")).
    WithID("open-btn").
    OnClick(c, ui.OpenModal("my-modal"))

// Modal
ui.Modal("my-modal", c,
    ui.H2(ui.T("Modal Title")),
    ui.P(ui.T("Modal content here.")),
    ui.Divider(),
    ui.Row("8px",
        ui.Button(ui.T("Cancel")).OnClick(c, ui.CloseModal("my-modal")),
        ui.Button(ui.T("Confirm")).OnClick(c, func(c *forge.Context) {
            // Handle confirm
            c.Set("modal_my-modal", false)
        }),
    ),
)
```

### Tabs

```go
ui.Tabs("my-tabs", c, []ui.TabItem{
    {Key: "profile", Label: "Profile", Content: ProfileContent(c)},
    {Key: "settings", Label: "Settings", Content: SettingsContent(c)},
    {Key: "notifications", Label: "Notifications", Content: NotificationsContent(c)},
})
```

### Dropdown

```go
ui.Dropdown("actions", c,
    ui.Button(ui.T("Actions ▼")),
    []ui.DropdownItem{
        {Key: "edit", Label: "Edit", OnClick: func(c *forge.Context) {
            // Handle edit
        }},
        {Key: "duplicate", Label: "Duplicate", OnClick: func(c *forge.Context) {
            // Handle duplicate
        }},
        {Key: "delete", Label: "Delete", OnClick: func(c *forge.Context) {
            // Handle delete
        }},
    },
)
```

## Media Components

### Avatar

```go
ui.Avatar("/images/user.jpg", "User Name")
```

### IFrame

```go
ui.IFrame("https://www.youtube.com/embed/VIDEO_ID").
    WithStyle("width:560px;height:315px")
```

## Creating Custom Components

```go
// Define component
func UserCard(name, email, avatar string) ui.UI {
    return ui.Card(
        ui.Row("12px",
            ui.Avatar(avatar, name),
            ui.Stack("4px",
                ui.Span(ui.T(name)).WithClass("font-bold"),
                ui.Span(ui.T(email)).WithClass("text-gray-500 text-sm"),
            ),
        ),
    )
}

// Use component
UserCard("John Doe", "john@example.com", "/avatars/john.jpg")
```

## Component Styling

All components accept standard styling:

```go
ui.Card(
    ui.T("Content"),
).WithClass("my-custom-class").
  WithStyle("max-width: 400px").
  Animate("fadeIn").
  Hover("lift")
```
