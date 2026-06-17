package combatmenu

import (
    "fmt"
	tea "charm.land/bubbletea/v2"
	// "charm.land/lipgloss/v2"
	// "charm.land/lipgloss/v2/table"
)

type menuOptions int

const (
	moveCharacter menuOptions = iota
	issueCommand
)

func (o menuOptions) String() string {
    switch o {
    case moveCharacter:
        return "Move"
    case issueCommand:
        return "Command"
    }
    return "ERROR"
}

type MoveCharacterMsg struct {
	Path []rune
}
type IssueCommandMsg struct {}

type Model struct {
    activeMenuOptions []menuOptions
    cursor int
}
	
func InitialModel() Model {
	return Model{[]menuOptions{moveCharacter, issueCommand,}, 0}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "h", "left":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.activeMenuOptions)-1 {
                m.cursor++
            }
        case "enter", " ", "l", "right":
            switch m.activeMenuOptions[m.cursor] {
            case moveCharacter:
				return m, func() tea.Msg { return MoveCharacterMsg{Path: []rune{'N',},} }
            case issueCommand:
                // return m, func() tea.Msg { return IssueCommandMsg{} }
				return m, func() tea.Msg { return MoveCharacterMsg{Path: []rune{'E',},} }
            }
        }
    }
    return m, nil
}

func (m Model) View() tea.View {
	s := ""
    for i, option := range m.activeMenuOptions {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %-10s\n", cursor, option)
    }
    return tea.NewView(s)
}
