package xxhtmlcli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/6oof/xxhtml/cmd/xxhtml/adapter"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// Define Lipgloss styles
var (
	styleGeneratedTitle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#a3be8c")) // Style for the title
	styleCode           = lipgloss.NewStyle().Foreground(lipgloss.Color("#ebdbb2"))            // Style for the code
	styleError          = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff6c6b"))            // Style for errors
)

func processHTML(htmlContent string, fullFlag bool) string {
	var nodes []string
	if fullFlag {
		doc, err := adapter.ParseFull(htmlContent)
		if err != nil {
			return fmt.Sprintf("Error parsing HTML document: %v", err)
		}
		nodes = []string{adapter.ConvertNode(doc)}
	} else {
		doc, err := adapter.ParseFragment(htmlContent)
		if err != nil {
			return fmt.Sprintf("Error parsing HTML fragment: %v", err)
		}
		for _, node := range doc {
			nodes = append(nodes, adapter.ConvertNode(node))
		}
	}

	var result strings.Builder
	for _, node := range nodes {
		result.WriteString(node)
	}
	s := result.String()

	// Remove trailing comma for proper syntax
	if len(s) > 0 && s[len(s)-1] == ',' {
		s = s[:len(s)-1]
	}
	return s
}

func Run() error {
	var htmlInput string
	var fullFlag bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Paste your HTML here:").
				Description("Char limit of 100,000").
				Validate(func(str string) error {
					if str == "" {
						return errors.New("Can't be left empty.")
					}
					return nil
				}).
				Value(&htmlInput).
				CharLimit(100000),
			huh.NewConfirm().
				Title("Full HTML document?").
				Value(&fullFlag),
		),
	).WithTheme(huh.ThemeCatppuccin())

	// Display the form and handle user input
	err := form.Run()
	if err != nil {
		return fmt.Errorf("Error displaying form: %w", err)
	}

	if htmlInput == "" {
		return fmt.Errorf(styleError.Render("No HTML provided!"))
	}

	generatedCode := processHTML(htmlInput, fullFlag)
	fmt.Println(styleGeneratedTitle.Render("Generated Go code:"))
	fmt.Println(styleCode.Render(generatedCode))

	return nil
}
