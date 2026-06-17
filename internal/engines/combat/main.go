package combat

import (
    // "fmt"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"internal/combat"
	"internal/engines/combatmenu"
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
func MakeCombatMap(w, h int) CombatMap 					{ return CombatMap{w, h, make([]combat.Creature, w*h)} }
func (cm CombatMap) At(x, y int) string					{ return cm.data[y*cm.w+x].GetIcon() }
func (cm CombatMap) Get(x, y int) combat.Creature		{ return cm.data[y*cm.w+x] }
func (cm CombatMap) Set(x, y int, c combat.Creature)	{ cm.data[y*cm.w+x] = c }
func (cm CombatMap) UpdateEntity(e Entity)	{
	x := e.position[1]
	y := e.position[0]
	cm.data[y*cm.w+x] = e.base_creature
}
func (cm CombatMap) Rows() int							{ return cm.w }
func (cm CombatMap) Columns() int						{ return cm.h }

var mapStyle = lipgloss.NewStyle().
			Width(40).
			Height(30).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder())
var menuStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder())

type Entity struct {
	base_creature combat.Creature
	position [2]int
	team int
	leads_team int
}
			
type Model struct {
	menu combatmenu.Model
	combatMap CombatMap
	entities []Entity
	conditions string
}

func InitialModel() Model {
	menu := combatmenu.InitialModel()
	mapWidth := 10
	mapHeight := 10
	combatMap := MakeCombatMap(mapWidth, mapHeight)
	entities := []Entity{
		Entity{combat.Creature{"Dog", creatures.GenericWolf, creatures.Personality{0,0,0,0,0}}, [2]int{2, 7}, 1, 0,},
		Entity{combat.Creature{"Person", creatures.NoSpecies, creatures.Personality{0,0,0,0,0}}, [2]int{2, 4}, 1, 1,},
		Entity{combat.Creature{"Wolf", creatures.GenericWolf, creatures.Personality{0,0,0,0,0}}, [2]int{6, 7}, 2, 2,},
		Entity{combat.Creature{"Cannon", creatures.PlantCannon, creatures.Personality{0,0,0,0,0}}, [2]int{9, 4}, 2, 0,},
	}
	for i:=0; i<len(entities); i++ {
		e := entities[i]
		combatMap.UpdateEntity(e)
	}
	return Model{menu, combatMap, entities, "TODO Conditions"}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		newCombatMenu, newCmd := m.menu.Update(msg)
		combatMenu, ok := newCombatMenu.(combatmenu.Model)
        if !ok {
            panic("Could not perform assertion on combatmenu model")
        }
        m.menu = combatMenu
        cmd = newCmd
		if cmd == nil {
			return m, nil
		}
		switch cmd := cmd().(type) {
		case combatmenu.MoveCharacterMsg:
			var entityID = 0
			for i:=0; i<len(m.entities); i++ {
				// m.entities[i].base_creature.Rename("Zombie")
				if m.entities[i].leads_team == 1 {
					// m.entities[i].base_creature.Rename("Leader")
					entityID = i
				}
				// m.combatMap.UpdateEntity(m.entities[i])
			}
			for i:=0; i<len(cmd.Path); i++ {
				x := m.entities[entityID].position[0]
				y := m.entities[entityID].position[1]
				tmp := m.combatMap.Get(y, x)
				if cmd.Path[i] == 'N' {
					m.combatMap.Set(y, x, m.combatMap.Get(y-1, x))
					m.combatMap.Set(y-1, x, tmp)
					m.entities[entityID].position[1] -= 1
				} else if cmd.Path[i] == 'E' {
					m.combatMap.Set(y, x, m.combatMap.Get(y, x+1))
					m.combatMap.Set(y, x+1, tmp)
					m.entities[entityID].position[0] += 1
				}
			}
		}
		return m, cmd
	}
	return m, nil
}

func (m Model) View() tea.View {
    s := "Battle Tamer v.0.mvp.1\n\n"
	t := table.New().BorderRow(true).Rows(table.DataToMatrix(m.combatMap)...)
	main_section := lipgloss.JoinHorizontal(lipgloss.Center, mapStyle.Render(t.Render()), menuStyle.Render(m.menu.View().Content))
	full_screen := lipgloss.JoinVertical(lipgloss.Center, s, main_section)
	return tea.NewView(full_screen)
}
