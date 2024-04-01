// TUIpod an ipod for the terminal
package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/minchao/go-apple-music"
)

type model struct {
	list      list.Model
	spinner   spinner.Model 
	textInput textinput.Model

	ipodUI *IPodUI
}

func initialModel() model {
	return model{
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		spinner: spinner.New(),
		textInput: textinput.New(), 
                currentTrack: go-apple-music.Track,

		ipodUI: NewIPodUI(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// Handle messages

	case tea.KeyMsg:
		switch msg.String() {
		// Handle key presses	
		}
	}

	m.ipodUI.Render() // Render UI

	return m, cmd  
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1) 
	}
}
