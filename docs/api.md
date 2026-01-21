# API Reference

Complete API reference for Forge.

## Package `forge`

### Types

```go
type Context = ctx.Context
type App = server.App
type DevServer = server.DevServer
type PageFunc = server.PageFunc
type LayoutFunc = server.LayoutFunc
```

### Functions

```go
func New() *App
func NewDev(watchDir string) *DevServer
```

---

## Package `forge/ui`

### Types

```go
type UI interface{ isUI() }

type Element struct {
    Tag      string
    ID       string
    Class    string
    Style    string
    Attrs    map[string]string
    Children []UI
    Events   map[string]string
}

type Text struct{ Value string }
type Raw struct{ HTML string }

type Style map[string]string

type TabItem struct {
    Key     string
    Label   string
    Content UI
}

type DropdownItem struct {
    Key     string
    Label   string
    OnClick func(*Context)
}
```

### Element Constructors

```go
func Div(children ...UI) Element
func Span(children ...UI) Element
func P(children ...UI) Element
func H1(children ...UI) Element
func H2(children ...UI) Element
func H3(children ...UI) Element
func H4(children ...UI) Element
func Button(children ...UI) Element
func A(children ...UI) Element
func Form(children ...UI) Element
func Label(children ...UI) Element
func Ul(children ...UI) Element
func Li(children ...UI) Element
func Tr(children ...UI) Element
func Th(children ...UI) Element
func Td(children ...UI) Element
func Main(children ...UI) Element
func Section(children ...UI) Element
func Article(children ...UI) Element
func Header(children ...UI) Element
func Footer(children ...UI) Element
func Nav(children ...UI) Element
func Input() Element
func Img() Element
func Br() Element
func Hr() Element
func El(tag string, children ...UI) Element
func T(s string) Text
```

### Element Methods

```go
func (e Element) WithID(id string) Element
func (e Element) WithClass(c string) Element
func (e Element) WithStyle(s string) Element
func (e Element) WithS(s Style) Element
func (e Element) WithAttr(k, v string) Element
func (e Element) WithChildren(children ...UI) Element
func (e Element) OnClick(c *Context, h EventHandler) Element
func (e Element) OnInput(c *Context, h EventHandler) Element
func (e Element) OnChange(c *Context, h EventHandler) Element
func (e Element) OnSubmit(c *Context, h EventHandler) Element
func (e Element) OnKeydown(c *Context, h EventHandler) Element
func (e Element) OnScroll(c *Context, h EventHandler) Element
func (e Element) Animate(name string) Element
func (e Element) Hover(effect string) Element
func (e Element) Draggable(id string) Element
func (e Element) DropZone(c *Context, onDrop func(*Context, string)) Element
```

### Components

```go
func Stack(gap string, children ...UI) Element
func Row(gap string, children ...UI) Element
func Card(children ...UI) Element
func Divider() Element
func Alert(msg, variant string) Element
func Badge(text string) Element
func Spinner() Element
func Progress(value int) Element
func Avatar(src, alt string) Element
func Table(headers []string, rows [][]UI) Element
func Modal(id string, c *Context, children ...UI) Element
func Tabs(id string, c *Context, tabs []TabItem) Element
func Dropdown(id string, c *Context, trigger UI, items []DropdownItem) Element
func VirtualList(id string, c *Context, height, itemHeight int, items []UI) Element
func SortableList(id string, c *Context, items []UI, onReorder func(*Context, int, int)) Element
func Embed(id string) Element
func IFrame(src string) Element
```

### Modal Helpers

```go
func OpenModal(id string) func(*Context)
func CloseModal(id string) func(*Context)
```

### Style Builder

```go
func S() Style
func (s Style) String() string
func (s Style) Display(v string) Style
func (s Style) Flex() Style
func (s Style) Grid() Style
func (s Style) Block() Style
func (s Style) None() Style
func (s Style) FlexDir(v string) Style
func (s Style) FlexRow() Style
func (s Style) FlexCol() Style
func (s Style) JustifyContent(v string) Style
func (s Style) AlignItems(v string) Style
func (s Style) Gap(v string) Style
func (s Style) P(v string) Style
func (s Style) Px(v string) Style
func (s Style) Py(v string) Style
func (s Style) M(v string) Style
func (s Style) Mx(v string) Style
func (s Style) My(v string) Style
func (s Style) W(v string) Style
func (s Style) H(v string) Style
func (s Style) MaxW(v string) Style
func (s Style) MinH(v string) Style
func (s Style) Bg(v string) Style
func (s Style) Color(v string) Style
func (s Style) Font(v string) Style
func (s Style) FontSize(v string) Style
func (s Style) FontWeight(v string) Style
func (s Style) Bold() Style
func (s Style) TextAlign(v string) Style
func (s Style) Center() Style
func (s Style) LineThrough() Style
func (s Style) Border(v string) Style
func (s Style) BorderRadius(v string) Style
func (s Style) Rounded(v string) Style
func (s Style) Shadow(v string) Style
func (s Style) Opacity(v string) Style
func (s Style) Transition(v string) Style
func (s Style) Pos(v string) Style
func (s Style) Absolute() Style
func (s Style) Relative() Style
func (s Style) Fixed() Style
func (s Style) Cursor(v string) Style
func (s Style) Pointer() Style
```

### CSS Functions

```go
func EnableTailwind()
func AddCSS(css string)
func GetCSS() string
func ResetCSS()
func AddHeadScript(script string)
func AddBodyScript(script string)
```

### CSS Constants

```go
const ResetStyles string
const BaseStyles string
const ComponentStyles string
const AnimationStyles string
const TailwindMinCSS string
```

---

## Package `forge/ctx`

### Types

```go
type EventHandler func(*Context)

type Context struct {
    Params map[string]string
}

type SessionStore interface {
    Save(id string, state map[string]any) error
    Load(id string) (map[string]any, error)
}

type MemoryStore struct{}
```

### Context Methods

```go
func New() *Context
func (c *Context) Set(key string, val any)
func (c *Context) Get(key string) any
func (c *Context) Int(key string) int
func (c *Context) String(key string) string
func (c *Context) Bool(key string) bool
func (c *Context) Persist(key string)
func (c *Context) PersistentState() map[string]any
func (c *Context) RestoreState(state map[string]any)
func (c *Context) On(id string, handler EventHandler)
func (c *Context) Handle(id string) bool
func (c *Context) HandleWithValue(id, value string) bool
func (c *Context) InputValue() string
```

### MemoryStore

```go
func NewMemoryStore() *MemoryStore
func (m *MemoryStore) Save(id string, state map[string]any) error
func (m *MemoryStore) Load(id string) (map[string]any, error)
```

---

## Package `forge/server`

### Types

```go
type App struct{}
type DevServer struct{ *App }
type PageFunc func(*ctx.Context) ui.UI
type LayoutFunc func(*ctx.Context, ui.UI) ui.UI
type UploadHandler func(filename string, data []byte) error
type StaticPage struct {
    Path   string
    Params map[string]string
}
```

### App Methods

```go
func New() *App
func (a *App) Route(path string, page PageFunc)
func (a *App) Layout(prefix string, layout LayoutFunc)
func (a *App) HandleUpload(path string, handler UploadHandler)
func (a *App) GenerateStatic(outDir string, pages []StaticPage) error
func (a *App) Run(addr string) error
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

### DevServer

```go
func NewDev(watchDir string) *DevServer
func (d *DevServer) Run(addr string) error
```

### Upload Helpers

```go
func SaveToDir(dir string) UploadHandler
```
