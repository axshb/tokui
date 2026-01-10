package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// when a file is successfully read
type FileContentMsg string

// generic errors
type ErrMsg error

// read file asynchronously
func ReadFileCmd(path string) tea.Cmd {
	return func() tea.Msg {
		content, err := os.ReadFile(path)
		if err != nil {
			return ErrMsg(err)
		}
		return FileContentMsg(content)
	}
}
