package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/axshb/tokui/internal/util"
)

type FocusArea int

const (
	FocusSidebar FocusArea = iota
	FocusEditor
)

type Model struct {
	textarea   textarea.Model
	fileList   list.Model
	tokenizer  *util.Tokenizer
	focus      FocusArea
	tokenCount int
	err        error
	width, height int
}

func NewModel(tke *util.Tokenizer) Model {
	ti := textarea.New()
	ti.Placeholder = "Enter text here..."
	ti.Focus()

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetHeight(1)
	delegate.SetSpacing(0)

	items := util.GetFiles(".")
	l := list.New(items, delegate, SidebarWidth, 20)
	l.Title = "Files"
	l.SetShowHelp(false)

	return Model{
		textarea:  ti,
		fileList:  l,
		tokenizer: tke,
		focus:     FocusEditor,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.resizeComponents()

	case ErrMsg:
		m.err = msg
		return m, nil

	case FileContentMsg:
		m.textarea.SetValue(string(msg))
		m.textarea.Focus()
		m.focus = FocusEditor
		m.tokenCount = m.tokenizer.Count(m.textarea.Value())

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlX:
			m.textarea.Reset()
			m.tokenCount = 0
			m.focus = FocusEditor
			m.textarea.Focus()
			return m, nil
		case tea.KeyTab, tea.KeyShiftRight, tea.KeyShiftLeft:
			m.cycleFocus()
		case tea.KeyEnter:
			if m.focus == FocusSidebar {
				selected, ok := m.fileList.SelectedItem().(util.FileItem)
				if ok {
					return m, ReadFileCmd(selected.PathStr)
				}
			}
		}
	}

	if m.focus == FocusSidebar {
		m.fileList, cmd = m.fileList.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.textarea, cmd = m.textarea.Update(msg)
		cmds = append(cmds, cmd)
		// recalc tokens every text update
		m.tokenCount = m.tokenizer.Count(m.textarea.Value())
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	sidebar := SidebarStyle.Render(m.fileList.View())

	tokens := fmt.Sprintf("\nTokens: %d", m.tokenCount)
	keybinds := "Shift Arrows, Tab: Focus | Ctrl+X: Clear | Ctrl+C: Quit"
	if m.err != nil {
		keybinds = fmt.Sprintf("ERROR: %v", m.err)
	}

	editorContent := lipgloss.JoinVertical(
		lipgloss.Left,
		LogoStyle.Render(LogoText),
		m.textarea.View(),
		tokens,
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(keybinds),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, sidebar, EditorStyle.Render(editorContent))
}

func (m *Model) cycleFocus() {
	if m.focus == FocusSidebar {
		m.focus = FocusEditor
		m.textarea.Focus()
	} else {
		m.focus = FocusSidebar
		m.textarea.Blur()
	}
}

func (m *Model) resizeComponents() {
	totalSidebarWidth := SidebarWidth + 4
	m.fileList.SetSize(SidebarWidth, m.height-4)

	editorWidth := m.width - totalSidebarWidth - 4
	editorHeight := m.height - 12

	if editorWidth > 0 {
		m.textarea.SetWidth(editorWidth)
	}
	if editorHeight > 0 {
		m.textarea.SetHeight(editorHeight)
	}
}
