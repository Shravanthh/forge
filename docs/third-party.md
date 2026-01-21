# Third-Party Integration

Embed third-party widgets and scripts in Forge.

## Adding Scripts

### Head Scripts

Load early (before body):

```go
func init() {
    // Analytics
    ui.AddHeadScript(`<script async src="https://www.googletagmanager.com/gtag/js?id=GA_ID"></script>`)
    
    // Fonts
    ui.AddHeadScript(`<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">`)
}
```

### Body Scripts

Load after DOM (end of body):

```go
func init() {
    // Chat widget
    ui.AddBodyScript(`<script src="https://widget.intercom.io/widget/APP_ID"></script>`)
    
    // Custom initialization
    ui.AddBodyScript(`
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            // Initialize widgets
        });
    </script>
    `)
}
```

## Embedding Widgets

### Container for Widget

```go
func Page(c *forge.Context) ui.UI {
    return ui.Div(
        ui.H1(ui.T("Dashboard")),
        
        // Widget will render here
        ui.Embed("analytics-widget"),
    )
}

func init() {
    ui.AddBodyScript(`
    <script>
        // Widget targets #analytics-widget
        AnalyticsWidget.init({
            container: '#analytics-widget',
            apiKey: 'xxx'
        });
    </script>
    `)
}
```

### IFrame Embeds

```go
// YouTube
ui.IFrame("https://www.youtube.com/embed/VIDEO_ID").
    WithStyle("width:560px;height:315px")

// Google Maps
ui.IFrame("https://www.google.com/maps/embed?pb=...").
    WithStyle("width:100%;height:400px")

// Twitter
ui.IFrame("https://platform.twitter.com/embed/...").
    WithStyle("width:100%;height:300px")
```

## Common Integrations

### Google Analytics

```go
func init() {
    ui.AddHeadScript(`
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-XXXXXXX"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag(){dataLayer.push(arguments);}
        gtag('js', new Date());
        gtag('config', 'G-XXXXXXX');
    </script>
    `)
}
```

### Stripe

```go
func init() {
    ui.AddHeadScript(`<script src="https://js.stripe.com/v3/"></script>`)
}

func CheckoutPage(c *forge.Context) ui.UI {
    return ui.Div(
        ui.Embed("card-element"),
        ui.Button(ui.T("Pay")).WithID("pay-btn"),
    )
}

func init() {
    ui.AddBodyScript(`
    <script>
        const stripe = Stripe('pk_xxx');
        const elements = stripe.elements();
        const card = elements.create('card');
        card.mount('#card-element');
    </script>
    `)
}
```

### Intercom

```go
func init() {
    ui.AddBodyScript(`
    <script>
        window.intercomSettings = {
            api_base: "https://api-iam.intercom.io",
            app_id: "APP_ID"
        };
    </script>
    <script src="https://widget.intercom.io/widget/APP_ID"></script>
    `)
}
```

### Crisp Chat

```go
func init() {
    ui.AddBodyScript(`
    <script>
        window.$crisp=[];window.CRISP_WEBSITE_ID="WEBSITE_ID";
        (function(){d=document;s=d.createElement("script");
        s.src="https://client.crisp.chat/l.js";s.async=1;
        d.getElementsByTagName("head")[0].appendChild(s);})();
    </script>
    `)
}
```

## Security Considerations

1. Only embed trusted scripts
2. Use Subresource Integrity (SRI) when available
3. Review third-party code before adding
4. Consider Content Security Policy (CSP)

```go
// With SRI
ui.AddHeadScript(`
<script src="https://example.com/lib.js" 
        integrity="sha384-xxx" 
        crossorigin="anonymous"></script>
`)
```
