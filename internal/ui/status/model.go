package status

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width, height int

	// For log selection and preview
	selectedLogIdx  int
	recentLogFiles  []string
	logScrollOffset int
	focus           FocusedBox
}

type FocusedBox int

const (
	FocusRecentLogs FocusedBox = iota
	FocusLogView
)

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}
