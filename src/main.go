package main

import (
	"computer/src/computer"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	comp := computer.New()
	p := tea.NewProgram(initialModel(comp), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
