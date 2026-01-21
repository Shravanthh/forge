# Styling

Forge supports Tailwind CSS, inline styles, and custom CSS.

## Tailwind CSS

Enable Tailwind (recommended):

```go
func init() {
    ui.EnableTailwind()
}
```

Use Tailwind classes:

```go
ui.Div(
    ui.H1(ui.T("Hello")).WithClass("text-3xl font-bold text-blue-600"),
    ui.P(ui.T("Welcome")).WithClass("text-gray-600 mt-2"),
).WithClass("max-w-4xl mx-auto p-8")
```

## Style Builder

For inline styles, use the Style builder:

```go
ui.Div(
    ui.T("Styled content"),
).WithS(ui.S().
    Flex().
    FlexCol().
    Gap("16px").
    P("20px").
    Bg("#ffffff").
    Rounded("8px").
    Shadow("0 2px 4px rgba(0,0,0,0.1)"))
```

### Available Style Methods

**Layout:**
```go
S().Display("flex")
S().Flex()           // display: flex
S().Grid()           // display: grid
S().Block()          // display: block
S().None()           // display: none
S().FlexDir("row")
S().FlexRow()        // flex-direction: row
S().FlexCol()        // flex-direction: column
S().JustifyContent("center")
S().AlignItems("center")
S().Gap("16px")
```

**Spacing:**
```go
S().P("16px")        // padding
S().Px("16px")       // padding-left + padding-right
S().Py("16px")       // padding-top + padding-bottom
S().M("16px")        // margin
S().Mx("auto")       // margin-left + margin-right
S().My("16px")       // margin-top + margin-bottom
```

**Size:**
```go
S().W("100%")        // width
S().H("200px")       // height
S().MaxW("800px")    // max-width
S().MinH("100vh")    // min-height
```

**Colors:**
```go
S().Bg("#ffffff")    // background
S().Color("#333")    // color
```

**Typography:**
```go
S().Font("Arial")
S().FontSize("16px")
S().FontWeight("bold")
S().Bold()           // font-weight: bold
S().TextAlign("center")
S().Center()         // text-align: center
S().LineThrough()    // text-decoration: line-through
```

**Border:**
```go
S().Border("1px solid #ccc")
S().BorderRadius("8px")
S().Rounded("8px")   // border-radius
```

**Effects:**
```go
S().Shadow("0 2px 4px rgba(0,0,0,0.1)")
S().Opacity("0.5")
S().Transition("all 0.2s ease")
```

**Position:**
```go
S().Pos("absolute")
S().Absolute()
S().Relative()
S().Fixed()
```

**Cursor:**
```go
S().Cursor("pointer")
S().Pointer()        // cursor: pointer
```

## Custom CSS

Add global CSS:

```go
func init() {
    ui.AddCSS(`
        .my-button {
            background: linear-gradient(to right, #667eea, #764ba2);
            color: white;
            padding: 12px 24px;
            border-radius: 8px;
            border: none;
            cursor: pointer;
        }
        .my-button:hover {
            opacity: 0.9;
        }
    `)
}
```

## Built-in CSS

Forge includes preset styles:

```go
ui.AddCSS(ui.ResetStyles)      // CSS reset
ui.AddCSS(ui.BaseStyles)       // Basic utility classes
ui.AddCSS(ui.ComponentStyles)  // Component styles
ui.AddCSS(ui.AnimationStyles)  // Animation classes
```

## Animations

Add animations to elements:

```go
// Entry animations (play once)
ui.Card(...).Animate("fadeIn")
ui.Card(...).Animate("slideUp")
ui.Card(...).Animate("scaleIn")

// Looping animations
ui.Div().Animate("bounce")
ui.Div().Animate("pulse")
ui.Div().Animate("spin")

// Hover effects
ui.Card(...).Hover("lift")    // lifts up with shadow
ui.Card(...).Hover("scale")   // grows slightly
ui.Card(...).Hover("glow")    // blue glow
```

Available animations:
- `fadeIn`, `fadeOut`
- `slideUp`, `slideDown`
- `scaleIn`
- `bounce`, `pulse`, `shake`, `spin`

Hover effects:
- `lift`, `scale`, `glow`, `fade`
