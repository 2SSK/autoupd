package status

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textInput textinput.Model
	submitted bool
	name      string
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter your name"
	ti.Focus()
	ti.Width = 20

	return Model{
		textInput: ti,
		submitted: false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
