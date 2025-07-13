package status

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}
