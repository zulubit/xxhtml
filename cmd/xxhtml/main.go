package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/6oof/xxhtml/cmd/xxhtml/adapter"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input      textarea.Model
	generated  string
	fullFlag   bool
	errMessage string
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Paste your HTML here..."
	ta.Focus()
	ta.SetWidth(80)
	ta.SetHeight(20) // Set height to accommodate more lines of input
	return model{
		input: ta,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.input.Value() != "" {
				m.generated = processHTML(m.input.Value(), m.fullFlag)
			} else {
				m.errMessage = "No HTML provided!"
			}
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.generated != "" {
		return fmt.Sprintf("\033[31mGenerated Go code:\033[0m\n%s\nPress q to exit.", m.generated)
	} else if m.errMessage != "" {
		return fmt.Sprintf("\033[31mError:\033[0m %s\n%s", m.errMessage, m.input.View())
	} else {
		return fmt.Sprintf("Paste your HTML and press Enter to generate Go code:\n%s", m.input.View())
	}
}

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

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
