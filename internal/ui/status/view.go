package status

func (m Model) View() string {
	if m.submitted {
		return SubmittedBoxStyle.Render("You entered: " + UserOutputStyle.Render(m.name))
	}

	return WelcomeStyle.Render("Welcome! Type your name:") + "\n" + m.textInput.View() + "\n\nPress q to quit"
}
