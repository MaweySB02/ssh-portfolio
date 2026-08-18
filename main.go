package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#5F3722"))

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8A583C"))

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#DAD2C6")).
			Background(lipgloss.Color("#5F3722"))

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F3722"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B59C85"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#8A583C")).
			Padding(1, 3).
			Width(70)
	asciiFaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8A583C")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#8A583C")).
			Padding(1, 2).
			Align(lipgloss.Center)
	asciiTextStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#8A583C")).
			Foreground(lipgloss.Color("#B59C85")).
			Padding(0, 2).
			Width(70).
			Align(lipgloss.Center)
)

var asciiArtText = `
███▄▄ ▄▄███ ▄███████▄ ███   ███ ████████ ███   ███
███▀███▀███ ███   ███ ███   ███ ███       ███▄███ 
███     ███ █████████ ███▄█▄███ ███▀▀▀     ▀███▀  
███     ███ ███   ███ ████▀████ ███▄▄▄▄▄    ███   
▀▀▀     ▀▀▀ ▀▀▀   ▀▀▀  ▀▀   ▀▀  ▀▀▀▀▀▀▀▀    ▀▀▀   
`

var asciiArtFace = `
***********************++--------====++++++*****************
*******************+:.                .-=++++***************
****************+-                       .=+++++************
***************+.                          :+++++++*********
*************+=                     .       .==++++++*******
************+=                      ..        -++++++++*****
***********++                    ...:::..      :+++++++++***
**********++                   .:=======--:.    =+++++++++**
*********++:                  :-=++++=====--.    +++++++++++
******++*+=                  :-=++++++++====-.   :++++++++++
******++++.                 :==+++++++++++===-.   ++++++++++
*****++++:                .----===++++++++====-   :+++++++++
****++++-               .-=====-----=+====-::::.   =++++++++
**++++++.             :---:.....:---====-------:   .++++++++
*++++++=           .-====-::..==:.==++=--.  =-.     .+++++++
+++++++:          :--===+++++====+++++=--======-     =++++++
++++++=           :--==+++++++++++++========++==-    =++++++
++++++=          ..:-===++++++++++==-======+++==-    =++++++
+++++=.           .:-===++++++++====+=======++==:    =++++++
+++++-            .:--===++++++++=------:-==+===    ==++++++
+++++:             :--========+++++++=+========    :==++++++
++++=.              :-============-------====-     :==++++++
++++-.              .:--========::--::::::-=-      :===+++++
+++=.              ....:---========-----===:       .====++++
++=-.              .::::::-==============-         :=====+++
++=-              ..::::::::--==========-          -=====+++
++-              ..::------:::::--===-:            :=====+++
+-.              .:----------:::                   :======++
+:.                 :--------:::.                  .-======+
-.                ......:-----::.                  .-=======
`

type screen int

const (
	menuScreen screen = iota
	aboutScreen
	projectsScreen
	skillsScreen
	experienceScreen
	contactScreen
)

type model struct {
	cursor int
	screen screen
}

var menuItems = []string{
	"About Me",
	"Projects",
	"Skills",
	"Experience",
	"Contact",
}

func initialModel() model {
	return model{
		cursor: 0,
		screen: menuScreen,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		key := msg.String()

		// Quit from anywhere
		if key == "q" || key == "ctrl+c" {
			return m, tea.Quit
		}

		// If we're on a portfolio page
		if m.screen != menuScreen {

			switch key {

			case "esc", "backspace":
				m.screen = menuScreen
			}

			return m, nil
		}

		// Menu navigation
		switch key {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(menuItems)-1 {
				m.cursor++
			}

		case "enter":

			switch m.cursor {

			case 0:
				m.screen = aboutScreen

			case 1:
				m.screen = projectsScreen

			case 2:
				m.screen = skillsScreen

			case 3:
				m.screen = experienceScreen

			case 4:
				m.screen = contactScreen
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.screen {

	case aboutScreen:
		return pageLayout(aboutView())

	case projectsScreen:
		return pageLayout(projectsView())

	case skillsScreen:
		return pageLayout(skillsView())

	case experienceScreen:
		return pageLayout(experienceView())

	case contactScreen:
		return pageLayout(contactView())

	default:
		return pageLayout(menuView(m))
	}
}

func pageLayout(content string) string {

	// LEFT: permanent face
	left := asciiFaceStyle.Render(asciiArtFace)

	// TOP RIGHT: permanent MAWEY logo
	maweyBox := asciiTextStyle.Render(asciiArtText)

	// BOTTOM RIGHT: current page
	right := boxStyle.Render(content)

	// Stack MAWEY ascii and content vertically
	rightColumn := lipgloss.JoinVertical(
		lipgloss.Left,
		maweyBox,
		right,
	)

	// Put face beside the right column
	gap := lipgloss.NewStyle().Width(2).Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		gap,
		rightColumn,
	)
}

func menuView(m model) string {

	name := titleStyle.Render("MAWEY BACELONIA")

	subtitle := subtitleStyle.Render(
		"Software Developer",
	)

	menu := ""

	for i, item := range menuItems {

		if i == m.cursor {

			menu += selectedStyle.Render(
				"✦  " + item,
			)

		} else {

			menu += itemStyle.Render(
				"   " + item,
			)
		}

		menu += "\n"
	}

	help := helpStyle.Render(
		"↑/↓ or j/k navigate • enter select • q quit",
	)

	return name +
		"\n" +
		subtitle +
		"\n\n" +
		menu +
		"\n" +
		help
}

func aboutView() string {

	title := titleStyle.Render("ABOUT ME")

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAD2C6")).
		Bold(true)

	content := fmt.Sprintf(`
I'm a %s who loves to create things
that feel %s and a little more human.


I enjoy building websites, experimenting with
interfaces, and learning how software works
under the hood.


%s
• Software development
• Algorithms & data structures
• UI/UX
• Creative web experiences
• Different coffee flavors
`,
		highlightStyle.Render("software developer"),
		highlightStyle.Render("creative"),
		highlightStyle.Render("Currently exploring:"),
	)

	help := helpStyle.Render(
		"\nesc / backspace → back",
	)

	return title + "\n" + content + help
}

func projectsView() string {

	title := titleStyle.Render("PROJECTS")

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAD2C6")).
		Bold(true)

	content := fmt.Sprintf(`
%s
Personalized books with a scrapbook-inspired
customization experience.

Built with:
Django • Python • MySQL • HTML • CSS • JavaScript

Role:
Full-stack Developer

%s
A cemetery management system with AR-assisted
navigation.

Built with:
Laravel • PHP • MySQL • HTML • CSS • JavaScript 

Role:
Software Developer & Project Manager
`,
		highlightStyle.Render("BOOKLIT"),
		highlightStyle.Render("CEMETERY MANAGEMENT SYSTEM"),
	)

	help := helpStyle.Render(
		"\nesc / backspace → back",
	)

	return title + "\n" + content + help
}

func skillsView() string {

	title := titleStyle.Render("SKILLS")

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAD2C6")).
		Bold(true)

	content := fmt.Sprintf(`
%s
PHP • JavaScript • Python • Java • SQL

%s
HTML • CSS • Laravel • Django

%s
MySQL

%s
Git • VS Code • Figma • Notion • Trello

%s
Go • Data Structures • Algorithms
`,
		highlightStyle.Render("LANGUAGES"),
		highlightStyle.Render("WEB"),
		highlightStyle.Render("DATABASE"),
		highlightStyle.Render("TOOLS"),
		highlightStyle.Render("CURRENTLY LEARNING"),
	)

	help := helpStyle.Render(
		"\nesc / backspace → back",
	)

	return title + "\n" + content + help
}

func experienceView() string {

	title := titleStyle.Render("EXPERIENCE")

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAD2C6")).
		Bold(true)

	content := fmt.Sprintf(`
%s
Capstone Project
2024 – 2025

Developed a cemetery management system
with AR-assisted navigation.

Responsibilities included:

• Software development
• Project management
• System design
• Documentation
• Team coordination


%s
The Umonics Method

505-hour internship involving software
development and technical work.
`,
		highlightStyle.Render("SOFTWARE DEVELOPER / PROJECT MANAGER"),
		highlightStyle.Render("SOFTWARE DEVELOPMENT INTERN"),
	)

	help := helpStyle.Render(
		"\nesc / backspace → back",
	)

	return title + "\n" + content + help
}

func contactView() string {

	title := titleStyle.Render("CONTACT")

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAD2C6")).
		Bold(true)

	content := fmt.Sprintf(`
%s
github.com/MaweySB02

%s
linkedin.com/in/mawey-bacelonia-120a05350

%s
maweysb@gmail.com

%s
Working on it!
`,
		highlightStyle.Render("GitHub"),
		highlightStyle.Render("LinkedIn"),
		highlightStyle.Render("Email"),
		highlightStyle.Render("Portfolio"),
	)

	help := helpStyle.Render(
		"\nesc / backspace → back",
	)

	return title + "\n" + content + help
}

func main() {

	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
