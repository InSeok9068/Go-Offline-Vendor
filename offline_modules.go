package main

import (
	"fmt"
	"offline_modules/cmd"

	_ "github.com/AlecAivazis/survey/v2"
	_ "github.com/atotto/clipboard"
	_ "github.com/briandowns/spinner"
	_ "github.com/charmbracelet/bubbles/cursor"
	_ "github.com/charmbracelet/bubbles/filepicker"
	_ "github.com/charmbracelet/bubbles/help"
	_ "github.com/charmbracelet/bubbles/key"
	_ "github.com/charmbracelet/bubbles/list"
	_ "github.com/charmbracelet/bubbles/paginator"
	_ "github.com/charmbracelet/bubbles/progress"
	_ "github.com/charmbracelet/bubbles/runeutil"
	_ "github.com/charmbracelet/bubbles/spinner"
	_ "github.com/charmbracelet/bubbles/stopwatch"
	_ "github.com/charmbracelet/bubbles/table"
	_ "github.com/charmbracelet/bubbles/textarea"
	_ "github.com/charmbracelet/bubbles/textinput"
	_ "github.com/charmbracelet/bubbles/timer"
	_ "github.com/charmbracelet/bubbles/viewport"
	_ "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/glamour"
	_ "github.com/charmbracelet/lipgloss"
	_ "github.com/chromedp/chromedp"
	_ "github.com/fatih/color"
	_ "github.com/gizak/termui/v3"
	_ "github.com/gocolly/colly"
	_ "github.com/olekukonko/tablewriter"
	_ "github.com/playwright-community/playwright-go"
	_ "github.com/schollz/progressbar/v3"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

func main() {
	fmt.Println("안녕")
	cmd.Playground()
}
