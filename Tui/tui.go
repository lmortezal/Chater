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
	serverMsg chan client.MsgStruct
}
type errMsg error

var messages = make(chan client.MsgStruct)
var serverMsg = make(chan client.MsgStruct)
type serverMessageMsg client.MsgStruct
var NameClient string = ""


func Tui_main(domain string , port int){
	go client.Startconnection(domain, port, messages , serverMsg)
	fmt.Println("Enter your name:")
	fmt.Scan(&NameClient)
	messages <- client.MsgStruct{Name: NameClient,Message: ""}

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
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#B5179E")),
		ReciveStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#4361EE")),
		err:         nil,
		serverMsg:   serverMsg,
	}

}


func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, listenServerMsg(m.serverMsg))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)



	

	switch msg := msg.(type) {
    case serverMessageMsg:
        // Format the server message and add it to the messages
        serverMessage := m.ReciveStyle.Render(msg.Name + ": ") + msg.Message
        m.messages = append(m.messages, serverMessage)
        m.viewport.SetContent(strings.Join(m.messages, "\n"))
        m.viewport.GotoBottom()
        return m, listenServerMsg(m.serverMsg)

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			//send the message to the client.go then that send it to the server
			messages <- client.MsgStruct{Name: NameClient,Message: m.textarea.Value()}
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil



	}
	return m, tea.Batch(tiCmd, vpCmd)
}

func listenServerMsg(serverMsg chan client.MsgStruct) tea.Cmd {
    return func() tea.Msg {
        return serverMessageMsg(<-serverMsg)
    }
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}