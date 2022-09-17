package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialModel(os.Args[1:]))
	if err := p.Start(); err != nil {
		os.Exit(1)
	}
}
