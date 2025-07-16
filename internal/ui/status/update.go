package status

import (
	"os"
	"path/filepath"
	"sort"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/2SSK/autoupd/internal/utils"
)

// Removed duplicate Model struct declaration. Model is defined in model.go

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.recentLogFiles) == 0 {
		files, _ := getRecentLogFiles()
		m.recentLogFiles = files
		m.selectedLogIdx = 0
	}
	if m.selectedLogIdx >= len(m.recentLogFiles) {
		m.selectedLogIdx = 0
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.focus == FocusRecentLogs {
				m.focus = FocusLogView
			} else {
				m.focus = FocusRecentLogs
			}
		case "up", "k":
			switch m.focus {
			case FocusRecentLogs:
				if m.selectedLogIdx > 0 {
					m.selectedLogIdx--
					m.logScrollOffset = 0 // reset scroll when changing log
				}
			case FocusLogView:
				if m.logScrollOffset > 0 {
					m.logScrollOffset--
				}
			}
		case "down", "j":
			switch m.focus {
			case FocusRecentLogs:
				if m.selectedLogIdx < len(m.recentLogFiles)-1 {
					m.selectedLogIdx++
					m.logScrollOffset = 0 // reset scroll when changing log
				}
			case FocusLogView:
				m.logScrollOffset++ // max scroll handled in view.go
			}
		}
	}

	return m, nil
}

// getRecentLogFiles returns the last 5 log files sorted by mod time desc
func getRecentLogFiles() ([]string, error) {
	files, err := filepath.Glob(utils.LogDir + "/*.log")
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := os.Stat(files[i])
		infoJ, _ := os.Stat(files[j])
		return infoI.ModTime().After(infoJ.ModTime())
	})
	if len(files) > 5 {
		files = files[:5]
	}
	return files, nil
}
