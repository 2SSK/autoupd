package status

import "github.com/charmbracelet/lipgloss"

var (
	WelcomeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("33")).
			Bold(true).
			Padding(0, 1)

	SubmittedBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Foreground(lipgloss.Color("39")).
				Padding(1, 2).
				Margin(1, 0)

	UserOutputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))
)
