package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	bpm        int
	step       int
	kick       []bool
	snare      []bool
	hat        []bool
	message    string
}

func initialModel() model {
	return model{
		bpm:   140,
		step: 0,
		kick:  make([]bool, 16),
		snare: make([]bool, 16),
		hat:   make([]bool, 16),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "left":
			if m.step > 0 {
				m.step--
			}

		case "right":
			if m.step < 15 {
				m.step++
			}

		case "k":
			m.kick[m.step] = !m.kick[m.step]

		case "s":
			m.snare[m.step] = !m.snare[m.step]

		case "h":
			m.hat[m.step] = !m.hat[m.step]

		case "up":
			m.bpm += 2

		case "down":
			if m.bpm > 40 {
				m.bpm -= 2
			}

		case "enter":
			render(m)
			m.message = "saved as output.wav"
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "BEATUI\n\n"
	s += fmt.Sprintf("BPM: %d   \"← →\" step | \"k/s/h\" toggle | \"⏎\" export | \"q\" quit\n\n", m.bpm)

	s += "Kick : " + line(m.kick, m.step) + "\n"
	s += "Snare: " + line(m.snare, m.step) + "\n"
	s += "Hat  : " + line(m.hat, m.step) + "\n\n"

	if m.message != "" {
		s += m.message + "\n"
	}

	return s
}

func line(track []bool, cursor int) string {
	out := ""
	for i, v := range track {
		c := "-"
		if v {
			c = "●"
		}
		if i == cursor {
			out += "[" + c + "]"
		} else {
			out += " " + c + " "
		}
	}
	return out
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
