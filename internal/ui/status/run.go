package status

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func RunDashboard() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running dashboard: %v\n", err)
	}
}
