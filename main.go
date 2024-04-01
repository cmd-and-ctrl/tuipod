package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbltea"
)

func main() {
    // Initialize your model here...
    
    p := tea.NewProgram(model)
    if err := p.Start(); err != nil {
        fmt.Printf("Oh no!\n\n%v\n\n", err)
        os.Exit(1)
    }
}
