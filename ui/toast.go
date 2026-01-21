package ui

import "github.com/Shravanthh/forge/ctx"

// Toast displays a toast notification.
func Toast(c *ctx.Context) Element {
	msg := c.String("_toast_msg")
	variant := c.String("_toast_variant")
	if msg == "" {
		return Div().WithID("toast-container").WithStyle("display:none")
	}

	return Div(
		Div(
			Span(T(msg)),
			Button(T("×")).WithID("toast-close").WithClass("toast-close").OnClick(c, func(c *ctx.Context) {
				c.Set("_toast_msg", "")
			}),
		).WithClass("toast toast-" + variant).Animate("slideUp"),
	).WithID("toast-container").WithClass("toast-container")
}

// ShowToast displays a toast message.
func ShowToast(c *ctx.Context, msg, variant string) {
	c.Set("_toast_msg", msg)
	c.Set("_toast_variant", variant)
}

// ToastSuccess shows a success toast.
func ToastSuccess(c *ctx.Context, msg string) { ShowToast(c, msg, "success") }

// ToastError shows an error toast.
func ToastError(c *ctx.Context, msg string) { ShowToast(c, msg, "error") }

// ToastInfo shows an info toast.
func ToastInfo(c *ctx.Context, msg string) { ShowToast(c, msg, "info") }

// ToastWarning shows a warning toast.
func ToastWarning(c *ctx.Context, msg string) { ShowToast(c, msg, "warning") }

// ToastStyles contains CSS for toast notifications.
const ToastStyles = `
.toast-container{position:fixed;top:20px;right:20px;z-index:1000}
.toast{display:flex;align-items:center;justify-content:space-between;min-width:300px;padding:16px;border-radius:8px;box-shadow:0 4px 12px rgba(0,0,0,0.15);gap:12px}
.toast-success{background:#10b981;color:#fff}
.toast-error{background:#ef4444;color:#fff}
.toast-info{background:#3b82f6;color:#fff}
.toast-warning{background:#f59e0b;color:#fff}
.toast-close{background:none;border:none;color:inherit;font-size:20px;cursor:pointer;padding:0;opacity:0.8}
.toast-close:hover{opacity:1}
`
