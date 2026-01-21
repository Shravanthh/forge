# Forge Architecture

## Overview

```
┌─────────────────────────────────────────────────────────────┐
│                         Server                               │
├─────────────────────────────────────────────────────────────┤
│  HTTP Request → Render Page → HTML Response                  │
│       ↓                                                      │
│  WebSocket ←→ Session Manager ←→ Context (State + Events)   │
│       ↓              ↓                                       │
│  Event Loop: Handle → Re-render → Diff → Send Patches       │
└─────────────────────────────────────────────────────────────┘
                           ↕ WebSocket
┌─────────────────────────────────────────────────────────────┐
│                         Client                               │
├─────────────────────────────────────────────────────────────┤
│  Event Delegation → Send to Server                          │
│  Receive Patches → Apply to DOM (morphdom-style)            │
└─────────────────────────────────────────────────────────────┘
```

## Core Components

### UI Package (`ui/`)

- `types.go` - Core types: `UI` interface, `Element`, `Text`, `Raw`
- `elements.go` - HTML element constructors and builder methods
- `events.go` - Event binding methods (`OnClick`, `OnInput`, etc.)

### Render Package (`render/`)

- `html.go` - Renders UI tree to HTML string with `data-forge-id` attributes

### Context Package (`ctx/`)

- `context.go` - State container with typed getters, event handler registry
- `session.go` - Session persistence interface and memory implementation

### Diff Package (`diff/`)

- `diff.go` - Compares UI trees, generates minimal patches

Patch types:
- `replace` - Replace entire element
- `attrs` - Update attributes only
- `text` - Update text content
- `insert` - Insert new child
- `remove` - Remove element

### Server Package (`server/`)

- `http.go` - HTTP server, routing, initial page render
- `websocket.go` - WebSocket handler, session management, event loop
- `client/forge.js` - Embedded client runtime

### Client Runtime (`server/client/forge.js`)

~2KB JavaScript that:
1. Connects WebSocket to server
2. Delegates DOM events to server
3. Applies incoming patches using morphdom-style diffing

## Data Flow

### Initial Load
1. Browser requests page
2. Server creates Context, renders UI tree to HTML
3. Server wraps HTML with `<script src="/forge.js">`
4. Browser renders HTML, loads forge.js
5. forge.js connects WebSocket

### User Interaction
1. User clicks button with `data-forge-click="handler_id"`
2. forge.js sends `{type: "event", id: "handler_id"}`
3. Server looks up handler, executes it
4. Handler mutates Context state
5. Server re-renders UI tree
6. Server diffs old vs new tree
7. Server sends patches via WebSocket
8. forge.js applies patches to DOM

## Component IDs

Every element gets a `data-forge-id` for targeting:
- Auto-generated from tree position: `0`, `0.0`, `0.1`, `0.0.2`
- Can be overridden with `.WithID("custom-id")`

## State Scoping

- **Connection-scoped** (default): State lives in memory, lost on disconnect
- **Session-scoped** (opt-in): Call `c.Persist("key")` to save across reconnects

## File Structure

```
forge/
├── ui/           # UI DSL
├── render/       # HTML renderer
├── ctx/          # Context & sessions
├── diff/         # Tree diffing
├── server/       # HTTP & WebSocket
│   └── client/   # JS runtime
├── cmd/forge/    # CLI tool
└── examples/     # Demo apps
```
