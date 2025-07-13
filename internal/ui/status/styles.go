package status

import "github.com/charmbracelet/lipgloss"

var (
	mauve   = lipgloss.Color("#cba6f7")
	base    = lipgloss.Color("#1e1e2e")
	text    = lipgloss.Color("#cdd6f4")
	overlay = lipgloss.Color("#313244")
	green   = lipgloss.Color("#a6e3a1")
	red     = lipgloss.Color("#f38ba8")

	DashboardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(mauve))

	LeftBoxStyle = lipgloss.NewStyle().
			BorderForeground(red).
			Foreground(text)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(red).
			Foreground(text)

	FooterStyle = lipgloss.NewStyle().
			Foreground(green).
			Align(lipgloss.Center)
)
