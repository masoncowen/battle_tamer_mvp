package mainmenu

import (
    "fmt"
    // "os"
    // "path/filepath"
	tea "charm.land/bubbletea/v2"
)

type menuOptions int

const (
    newBattle menuOptions = iota
    loadBattle
    optionsMenu
    quit
)

func (o menuOptions) String() string {
    switch o {
    case newBattle:
        return "New Battle"
    case loadBattle:
        return "Load Battle"
    case optionsMenu:
        return "Options"
    case quit:
        return "Quit"
    }
    return "ERROR"
}

type StartCombatMsg struct {}
type OptionMsg struct {}

type Model struct {
    activeMenuOptions []menuOptions
    cursor int
}

func InitialModel() Model {
    return Model{[]menuOptions{newBattle, optionsMenu, quit,}, 0}
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
            case newBattle:
                return m, func() tea.Msg { return StartCombatMsg{} }
            case optionsMenu:
                return m, func() tea.Msg { return OptionMsg{} }
            case quit:
                return m, tea.Quit
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

        s += fmt.Sprintf("%s %s\n", cursor, option)
    }
    return tea.NewView(s)
}
