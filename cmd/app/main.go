package main

import (
	"fmt"
	"internal/constants"
	"internal/engines/mainmenu"
	"internal/engines/options"
	"internal/engines/combat"
	"os"
	// "path/filepath"
	// "time"

	tea "charm.land/bubbletea/v2"
)

type sessionState int

const (
    mainMenuView sessionState = iota
    optionsView
	combatView
)

type model struct {
    state sessionState
    mainmenu tea.Model
    options tea.Model
	combat tea.Model
}

func initialModel() model {
	return model{
        state: mainMenuView,
        mainmenu: mainmenu.InitialModel(),
        options: options.InitialModel(),
		combat: combat.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
	case mainmenu.StartCombatMsg:
		m.state = combatView
    case mainmenu.OptionMsg:
        m.state = optionsView
    case options.BackMsg, constants.BackMsg:
        m.state = mainMenuView
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }
        
    switch m.state {
    case mainMenuView:
        newMainMenu, newCmd := m.mainmenu.Update(msg)
        mainMenuModel, ok := newMainMenu.(mainmenu.Model)
        if !ok {
            panic("Could not perform assertion on mainmenu model")
        }
        m.mainmenu = mainMenuModel
        cmd = newCmd
    case optionsView:
        newOptions, newCmd := m.options.Update(msg)
        optionsModel, ok := newOptions.(options.Model)
        if !ok {
            panic("Could not perform assertion on options model")
        }
        m.options = optionsModel
        cmd = newCmd
	case combatView:
        newCombat, newCmd := m.combat.Update(msg)
        combatModel, ok := newCombat.(combat.Model)
        if !ok {
            panic("Could not perform assertion on options model")
        }
        m.combat = combatModel
        cmd = newCmd
    }
    return m, cmd
}

func (m model) View() tea.View {
    switch m.state {
    case mainMenuView:
        return m.mainmenu.View()
    case optionsView:
        return m.options.View()
	case combatView:
		return m.combat.View()
    }
    return tea.NewView("Invalid Model has been selected")
}

func main() {
    m := initialModel()
    p := tea.NewProgram(m)
    if _, err := p.Run(); err != nil {
        fmt.Printf("You done fucked up A-A-Ron: %v", err)
        os.Exit(1)
    }
}
