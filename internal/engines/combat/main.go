package combat

import (
    // "fmt"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"internal/combat"
	"internal/creatures"
)

type Matrix[T any] struct {
	w, h int
	data []T
}

func MakeMatrix[T any](w, h int) Matrix[T]	{ return Matrix[T]{w, h, make([]T, w*h)} }
func (m Matrix[T]) At(x, y int) T			{ return m.data[y*m.w+x] }
func (m Matrix[T]) Set(x, y int, t T)		{ m.data[y*m.w+x] = t }
func (m Matrix[T]) Rows() int				{ return m.w }
func (m Matrix[T]) Columns() int			{ return m.h }

type CombatMap  struct {
	w, h int
	data []combat.Creature
}
func MakeCombatMap(w, h int) CombatMap 		{ return CombatMap{w, h, make([]combat.Creature, w*h)} }
func (cm CombatMap) At(x, y int) string		{ return cm.data[y*cm.w+x].GetIcon() }
func (cm CombatMap) Set(x, y int, c combat.Creature)		{ cm.data[y*cm.w+x] = c }
func (cm CombatMap) Rows() int				{ return cm.w }
func (cm CombatMap) Columns() int			{ return cm.h }

type Model struct {
	combatMap CombatMap
	conditions string
}

func InitialModel() Model {
	mapWidth := 10
	mapHeight := 10
	combatMap := MakeCombatMap(mapWidth, mapHeight)
	// for i:=0; i<mapHeight; i++ {
	// 	for j:=0; j<mapWidth; j++ {
	// 		combatMap.Set(i, j, " ")
	// 	}
	// }
	return Model{combatMap, "TODO Conditions"}
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
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
    s := "Battle Tamer v.0.mvp.1\n\n"
	m.combatMap.Set(3, 2, combat.Creature{"Dog", creatures.GenericWolf, creatures.Personality{0,0,0,0,0}})
	m.combatMap.Set(4, 2, combat.Creature{"Wolf", creatures.GenericWolf, creatures.Personality{0,0,0,0,0}})
	m.combatMap.Set(7, 6, combat.Creature{"Cannon", creatures.PlantCannon, creatures.Personality{0,0,0,0,0}})
	t := table.New().BorderRow(true).Rows(table.DataToMatrix(m.combatMap)...)

	return tea.NewView(lipgloss.JoinVertical(lipgloss.Center, s, t.Render()))
}
