package status

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.renderDashboard(), m.renderFooter())
}

func (m Model) renderDashboard() string {
	totalWidth := m.width - 2
	totalHeight := m.height - 5

	LeftBoxWidth := totalWidth / 3
	RightBoxWidth := (totalWidth - LeftBoxWidth) - 2
	boxHeight := totalHeight - 2
	RightboxHeight := totalHeight - 2

	leftBox := m.LeftView(LeftBoxWidth, boxHeight)
	rightBox := m.LogView(RightBoxWidth, RightboxHeight, "Right Box Content")

	return DashboardStyle.
		Width(totalWidth).
		Height(totalHeight).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox))
}

func (m Model) LeftView(w, h int) string {
	msg := m.Title(w, "Left Box Title")

	width := w - 2
	height := h/3 - 1
	box1 := m.LogView(width, height, msg)
	box2 := m.LogView(width, height, msg)
	box3 := m.LogView(width, height, msg)

	return LeftBoxStyle.
		Width(w).
		Height(h).
		Render(lipgloss.JoinVertical(lipgloss.Left, box1, box2, box3))
}

func (m Model) LogView(w, h int, msg string) string {
	return BoxStyle.
		Width(w).
		Height(h).
		Render(m.Title(w, msg))
}

func (m Model) renderFooter() string {
	return FooterStyle.
		Width(m.width - 2).
		Render("Press q to quit. made with ❤️ by @2SSK")
}

func (m Model) Title(w int, msg string) string {
	return lipgloss.NewStyle().
		Width(w).
		Height(1).
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color("#000")).
		Render(msg)
}
