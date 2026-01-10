package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/axshb/tokui/internal/ui"
	"github.com/axshb/tokui/internal/util"
)

func main() {
	tokenizer, err := util.NewTokenizer()
	if err != nil {
		log.Fatal("Failed to initialize tokenizer:", err)
	}

	p := tea.NewProgram(ui.NewModel(tokenizer), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
