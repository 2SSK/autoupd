package status

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/2SSK/autoupd/internal/utils"
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
	RightboxHeight := totalHeight

	leftBox := m.LeftView(LeftBoxWidth, boxHeight)
	rightBox := m.LogView(RightBoxWidth, RightboxHeight)

	return DashboardStyle.
		Width(totalWidth).
		Height(totalHeight).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox))
}

func (m Model) LeftView(w, h int) string {
	width := w - 2
	height := h / 3
	box1 := m.status(width, height)
	box2 := m.osInformation(width, height)
	box3 := m.recentLogs(width, height)

	return LeftBoxStyle.
		Width(w).
		Height(h).
		Render(lipgloss.JoinVertical(lipgloss.Left, box1, box2, box3))
}

func (m Model) LogView(w, h int) string {
	files, err := filepath.Glob((utils.LogDir + "/*.log"))
	if err != nil || len(files) == 0 {
		return BoxStyle.Width(w).Height(h).Render(m.Title(w, "Log Preview") + "\nNo log file selected")
	}
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := os.Stat(files[i])
		infoJ, _ := os.Stat(files[j])
		return infoI.ModTime().After(infoJ.ModTime())
	})
	idx := m.selectedLogIdx
	if idx >= len(files) {
		idx = 0
	}
	data, err := os.ReadFile(files[idx])
	content := ""
	if err != nil {
		content = "Error reading log file"
	} else {
		lines := strings.Split(string(data), "\n")
		visibleLines := h - 2 // leave space for title
		if visibleLines < 1 {
			visibleLines = 1
		}
		totalLines := len(lines)
		maxOffset := max(0, totalLines-visibleLines)
		start := m.logScrollOffset
		if start < 0 {
			start = 0
		}
		if start > maxOffset {
			start = maxOffset
		}
		end := start + visibleLines
		if end > totalLines {
			end = totalLines
		}
		content = strings.Join(lines[start:end], "\n")
	}
	title := filepath.Base(files[idx])
	style := BoxStyle.Width(w).Height(h)
	if m.focus == FocusLogView {
		style = style.BorderForeground(lipgloss.Color("#FFD700")) // highlight border
	}
	return style.Render(m.Title(w, title) + "\n" + content)
}

// Helper for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func (m Model) status(w, h int) string {
	status := "Unknown"
	if utils.WasUpdateSuccessful() {
		status = "Successful"
	} else {
		status = "Failed or Not Run Yet"
	}

	lastRun := getLastRun()
	nextRun := getNextRun()

	content := fmt.Sprintf("%s\n%s\n%s",
		"Today's Update: "+status,
		"Last Run: "+lastRun,
		"Next Run: "+nextRun,
	)

	return BoxStyle.
		Width(w).
		Height(h).
		Render(m.Title(w, "Status Information:") + "\n" + content)
}

func (m Model) osInformation(w, h int) string {
	osName := getOSName()
	pkgManager := utils.DetectPackageManager()
	kernel := getKernelVersion()
	timerStatus := "Inactive"
	if utils.IsTimerActive() {
		timerStatus = "Active"
	}

	content := fmt.Sprintf("OS: %s\nPackage Manager: %s\nKernel: %s\nTimer Status: %s", osName, pkgManager, kernel, timerStatus)

	return BoxStyle.
		Width(w).
		Height(h).
		Render(m.Title(w, "System Information:") + "\n" + content)
}

func (m Model) recentLogs(w, h int) string {
	files, err := filepath.Glob((utils.LogDir + "/*.log"))
	if err != nil {
		return "Error listing log files"
	}

	sort.Slice(files, func(i, j int) bool {
		infoI, _ := os.Stat(files[i])
		infoJ, _ := os.Stat(files[j])
		return infoI.ModTime().After(infoJ.ModTime())
	})

	count := min(len(files), 5)
	lines := make([]string, count)
	for i := 0; i < count; i++ {
		name := filepath.Base(files[i])
		if i == m.selectedLogIdx {
			lines[i] = lipgloss.NewStyle().Background(lipgloss.Color("#FFD700")).Foreground(lipgloss.Color("#000")).Render("> " + name)
		} else {
			lines[i] = "  " + name
		}
	}

	style := BoxStyle.Width(w).Height(h)
	if m.focus == FocusRecentLogs {
		style = style.BorderForeground(lipgloss.Color("#FFD700")) // highlight border
	}
	return style.Render(m.Title(w, "Recent Logs:") + "\n" + strings.Join(lines, "\n"))
}
