# File Uploads

Handle file uploads in Forge.

## Setup

Register an upload handler:

```go
app := forge.New()

// Save to directory
app.HandleUpload("/upload", server.SaveToDir("./uploads"))

// Or custom handler
app.HandleUpload("/upload", func(filename string, data []byte) error {
    // Process file...
    return nil
})
```

## Upload Form

```go
func UploadForm(c *forge.Context) ui.UI {
    return ui.Div(
        ui.Form(
            ui.Input().
                WithAttr("type", "file").
                WithAttr("name", "file").
                WithID("file-input"),
            ui.Button(ui.T("Upload")).
                WithAttr("type", "submit"),
        ).WithAttr("action", "/upload").
          WithAttr("method", "POST").
          WithAttr("enctype", "multipart/form-data"),
    )
}
```

## Custom Upload Handler

```go
app.HandleUpload("/upload/avatar", func(filename string, data []byte) error {
    // Validate file type
    if !strings.HasSuffix(filename, ".jpg") && !strings.HasSuffix(filename, ".png") {
        return fmt.Errorf("only JPG and PNG allowed")
    }
    
    // Validate size (max 5MB)
    if len(data) > 5*1024*1024 {
        return fmt.Errorf("file too large")
    }
    
    // Generate unique filename
    ext := filepath.Ext(filename)
    newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
    
    // Save file
    return os.WriteFile(filepath.Join("./uploads/avatars", newName), data, 0644)
})
```

## Multiple Upload Endpoints

```go
app.HandleUpload("/upload/images", server.SaveToDir("./uploads/images"))
app.HandleUpload("/upload/documents", server.SaveToDir("./uploads/documents"))
app.HandleUpload("/upload/avatars", avatarHandler)
```

## Serving Uploaded Files

Add a static file handler:

```go
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Serve uploaded files
    if strings.HasPrefix(r.URL.Path, "/uploads/") {
        http.ServeFile(w, r, "."+r.URL.Path)
        return
    }
    // ... rest of handler
}
```

## Upload Progress

For large files, consider using a JavaScript-based uploader with progress. Inject the script:

```go
ui.AddBodyScript(`
<script>
function uploadWithProgress(file, url, onProgress) {
    return new Promise((resolve, reject) => {
        const xhr = new XMLHttpRequest();
        const formData = new FormData();
        formData.append('file', file);
        
        xhr.upload.onprogress = (e) => {
            if (e.lengthComputable) {
                onProgress(Math.round((e.loaded / e.total) * 100));
            }
        };
        
        xhr.onload = () => resolve(xhr.response);
        xhr.onerror = () => reject(xhr.statusText);
        
        xhr.open('POST', url);
        xhr.send(formData);
    });
}
</script>
`)
```
