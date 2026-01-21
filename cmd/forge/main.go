package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "new":
		if len(os.Args) < 3 {
			fatal("Usage: forge new <project-name>")
		}
		newProject(os.Args[2])
	case "dev":
		runDev()
	case "start":
		runStart()
	case "build":
		runBuild()
	case "test":
		runTest()
	case "version", "-v", "--version":
		fmt.Println("Forge v" + version)
	case "help", "-h", "--help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`Forge - Go-native UI Framework

Usage: forge <command> [options]

Commands:
  new <name>    Create a new Forge project
  dev           Start development server with hot reload
  start         Start production server
  build         Build for production
  test          Run tests
  version       Show version

Examples:
  forge new my-app
  forge dev
  forge start
  forge build`)
}

func newProject(name string) {
	fmt.Printf("Creating new Forge project: %s\n", name)

	// Create directory
	if err := os.MkdirAll(name, 0755); err != nil {
		fatal("Failed to create directory: " + err.Error())
	}

	// Detect forge location
	forgePath := detectForgePath()
	
	// Create go.mod
	gomod := fmt.Sprintf(`module %s

go 1.25

require forge v1.0.0
`, name)
	
	if forgePath != "" {
		gomod += fmt.Sprintf("\nreplace forge => %s\n", forgePath)
	}
	
	writeFile(filepath.Join(name, "go.mod"), gomod)

	// Create main.go
	writeFile(filepath.Join(name, "main.go"), `package main

import (
	"forge"
	"forge/ui"
)

func init() {
	ui.EnableTailwind()
	ui.AddCSS(ui.ResetStyles)
}

func main() {
	app := forge.New()
	app.Route("/", HomePage)
	app.Run(":3000")
}

func HomePage(c *forge.Context) ui.UI {
	count := c.Int("count")

	return ui.Div(
		ui.H1(ui.T("Welcome to Forge")).WithClass("text-3xl font-bold mb-4"),
		ui.P(ui.T("Build modern UIs with pure Go")).WithClass("text-gray-600 mb-8"),
		
		ui.Div(
			ui.Button(ui.T("-")).
				WithClass("px-4 py-2 bg-red-500 text-white rounded-l").
				WithID("dec").
				OnClick(c, func(c *forge.Context) {
					c.Set("count", c.Int("count")-1)
				}),
			ui.Span(ui.T(itoa(count))).WithClass("px-6 py-2 bg-gray-100"),
			ui.Button(ui.T("+")).
				WithClass("px-4 py-2 bg-green-500 text-white rounded-r").
				WithID("inc").
				OnClick(c, func(c *forge.Context) {
					c.Set("count", c.Int("count")+1)
				}),
		).WithClass("inline-flex"),
	).WithClass("min-h-screen flex flex-col items-center justify-center")
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	s := ""
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	if neg {
		s = "-" + s
	}
	return s
}
`)

	// Create pages directory
	os.MkdirAll(filepath.Join(name, "pages"), 0755)
	
	// Create components directory
	os.MkdirAll(filepath.Join(name, "components"), 0755)

	// Create .gitignore
	writeFile(filepath.Join(name, ".gitignore"), `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
/bin/

# Test binary
*.test

# Output
/dist/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
`)

	fmt.Println("✓ Created project structure")
	fmt.Println("\nNext steps:")
	fmt.Printf("  cd %s\n", name)
	fmt.Println("  go mod tidy")
	fmt.Println("  forge dev")
}

func runDev() {
	fmt.Println("Starting Forge development server...")
	
	// Check if main.go exists
	if _, err := os.Stat("main.go"); os.IsNotExist(err) {
		fatal("main.go not found. Are you in a Forge project directory?")
	}

	// Set dev mode environment variable
	os.Setenv("FORGE_ENV", "development")
	
	// Run with go run
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	
	if err := cmd.Run(); err != nil {
		fatal("Failed to start dev server: " + err.Error())
	}
}

func runStart() {
	fmt.Println("Starting Forge production server...")
	
	// Check for binary first
	binName := getBinaryName()
	if _, err := os.Stat(binName); err == nil {
		cmd := exec.Command("./" + binName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return
	}

	// Fall back to go run
	os.Setenv("FORGE_ENV", "production")
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func runBuild() {
	fmt.Println("Building Forge application...")
	
	binName := getBinaryName()
	
	cmd := exec.Command("go", "build", "-ldflags=-s -w", "-o", binName, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fatal("Build failed: " + err.Error())
	}
	
	fmt.Printf("✓ Built: %s\n", binName)
}

func runTest() {
	fmt.Println("Running tests...")
	
	cmd := exec.Command("go", "test", "./...", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func getBinaryName() string {
	dir, _ := os.Getwd()
	name := filepath.Base(dir)
	if strings.HasSuffix(os.Getenv("GOOS"), "windows") {
		return name + ".exe"
	}
	return name
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fatal("Failed to write " + path + ": " + err.Error())
	}
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}

func detectForgePath() string {
	// Get absolute path of current directory
	cwd, _ := os.Getwd()
	
	// Check common locations relative to where project will be created
	home, _ := os.UserHomeDir()
	paths := []string{
		filepath.Join(cwd, "../forge"),
		filepath.Join(cwd, "../../forge"),
		filepath.Join(home, "forge"),
		filepath.Join(home, "go/src/forge"),
	}
	
	for _, p := range paths {
		absPath, _ := filepath.Abs(p)
		if _, err := os.Stat(filepath.Join(absPath, "forge.go")); err == nil {
			// Return relative path from project location
			return absPath
		}
	}
	return ""
}
