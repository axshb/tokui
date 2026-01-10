package ui

import "github.com/charmbracelet/lipgloss"

const (
	SidebarWidth = 20
)

var (
	SidebarStyle = lipgloss.NewStyle().
			Width(SidebarWidth).
			Padding(1, 1).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("240"))

	EditorStyle = lipgloss.NewStyle().
			Padding(1, 2)

	LogoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("63")).
			Bold(true)

	LogoText = `
 _____     _             _
|_   _|__ | | ___ _   _ (_)
  | |/ _ \| |/ / | | | || |
  | | (_) |   <| |_| | || |
  |_|\___/|_|\_\\__,_|_||_|
`
)
