package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lmortezal/Chater/client"
)
type model struct{
	viewport viewport.Model
	messages []string
	textarea textarea.Model
	senderStyle lipgloss.Style
	ReciveStyle lipgloss.Style
	err error
}

type errMsg error

func Tui_main(){
	Messages := make(chan<- client.ServerMsg)
	
	client.Startconnection("localhost", 8080, Messages)

	p := tea.NewProgram(initialModel())
	if _ , err := p.Run(); err != nil{
		panic(err)
	}
}

func initialModel() model{
	ta := textarea.New()
	ta.Placeholder = "Type a message and hit enter to send"
	ta.Focus()

	ta.Prompt = "| "
	ta.CharLimit = 280
	ta.SetWidth(30)
	ta.SetHeight(5)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 10)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC300")),
		ReciveStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#001D3D")),
		err:         nil,
	}

}


func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)



	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}
	case client.ServerMsg:
        // Format the server message and add it to the messages
        serverMessage := m.ReciveStyle.Render(msg.Name + ": ") + msg.Message
        m.messages = append(m.messages, serverMessage)
        m.viewport.SetContent(strings.Join(m.messages, "\n"))
        m.viewport.GotoBottom()
    
	


	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}