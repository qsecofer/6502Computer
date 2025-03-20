package main

import (
	"computer/src/computer"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {	
	comp *computer.Computer

	width int
}

type UpdateModelMsg struct {
	PC  *uint16
	SP  *uint8
	LDA *uint8
	LDX *uint8
	LDY *uint8
	FLG *uint8	
}

func initialModel(c *computer.Computer) model {
	return model{			
		comp:c,

		width: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// func setModelCmd(newValues UpdateModelMsg) tea.Cmd {
// 	return func() tea.Msg {
// 		return newValues
// 	}
// }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit		
		case "s":
			m.comp.Cpu.ExecuteInstruction(m.comp.Bus)			
			return m,nil
		}
	case UpdateModelMsg:
		// if msg.PC != nil {			
		// 	m.pc = *msg.PC
		// }
		// if msg.SP != nil {
		// 	m.sp = *msg.SP
		// }
		// if msg.LDA != nil {
		// 	m.lda = *msg.LDA
		// }
		// if msg.LDX != nil {
		// 	m.ldx = *msg.LDX
		// }
		// if msg.LDY != nil {
		// 	m.ldy = *msg.LDY
		// }
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}


const (
	headerColor = "12"
	valueColor = "7"
	flagColor ="9"
	ProgramCounterWidth = 14
	registerWidth = 9
	flagsWidth = 14
)

func ProgramCounterStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Bold(true).
	Foreground(lipgloss.Color(headerColor)).
	Width(ProgramCounterWidth).
	Align(lipgloss.Center)
}

func HeaderValueStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Foreground(lipgloss.Color(valueColor)).
	Width(ProgramCounterWidth).
	Align(lipgloss.Center)
}

func RegistersStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Foreground(lipgloss.Color(headerColor)).
	Width(registerWidth).
	Align(lipgloss.Center)
}

func FlagsStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Foreground(lipgloss.Color(headerColor)).
	Width(flagsWidth).
	Align(lipgloss.Center)
}

func RegisterValueStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Foreground(lipgloss.Color(valueColor)).
	Width(registerWidth).
	Align(lipgloss.Center)
}

func FlagsValueStyle() lipgloss.Style {
	return lipgloss.NewStyle().		
	Foreground(lipgloss.Color(flagColor)).
	Width(flagsWidth).
	Align(lipgloss.Center)
}

func LedString(value uint8, tp ...string) string {
	var sb strings.Builder
	for i := 0; i < 8; i++ {		
		if len(tp) > 0 && tp[0]=="flags" && i==2 {
			sb.WriteString("-")	
			continue
		}		
		shifted := int16(value << i)
		if (shifted & 128) == 128 {
			sb.WriteString("●")
		} else {
			sb.WriteString("○")
		}
	}
	return sb.String()
}

func CPUHeader() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		ProgramCounterStyle().Render("PC"),
		RegistersStyle().Render("SP"),
		RegistersStyle().Render("LDA"),
		RegistersStyle().Render("LDX"),
		RegistersStyle().Render("LDY"),
		FlagsStyle().Render("NV-BDIZC"),
	)
}

func CPUValues(m model) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		HeaderValueStyle().Render(fmt.Sprintf("0x%04X", m.comp.Cpu.PC)),		
		RegisterValueStyle().Render(fmt.Sprintf("0x%02X", m.comp.Cpu.SP)),
		RegisterValueStyle().Render(fmt.Sprintf("0x%02X\n%s", m.comp.Cpu.A, LedString(m.comp.Cpu.A))),
		RegisterValueStyle().Render(fmt.Sprintf("0x%02X\n%s", m.comp.Cpu.X, LedString(m.comp.Cpu.X))),
		RegisterValueStyle().Render(fmt.Sprintf("0x%02X\n%s", m.comp.Cpu.Y, LedString(m.comp.Cpu.Y))),
		FlagsValueStyle().Render(LedString(m.comp.Cpu.Flags2Byte(), "flags")),
	)
}

func CPUView(m model) string {
	headers := CPUHeader()
	values := CPUValues(m)

	data := lipgloss.JoinVertical(lipgloss.Left, headers, values)

	dataStyle := lipgloss.NewStyle().
		Width(m.width - 6).
		Align(lipgloss.Left)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(headerColor)).
		Width(m.width - 4).
		Align(lipgloss.Left).
		Padding(1).
		Margin(1)

	return borderStyle.Render(dataStyle.Render(data))
}


func (m model) View() string {
	



	

	return  CPUView(m)+ "\n\nPress q to quit."
}