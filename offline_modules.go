package main

import (
	"fmt"

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
	_ "github.com/fatih/color"
	_ "github.com/gizak/termui/v3"
	_ "github.com/gizak/termui/v3/widgets"
	_ "github.com/go-rod/rod"
	_ "github.com/gocolly/colly"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/olekukonko/tablewriter"
	_ "github.com/schollz/progressbar/v3"
	_ "github.com/sijms/go-ora/v2"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	_ "github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/suite"
	_ "github.com/xuri/excelize/v2"
	_ "gopkg.in/gomail.v2"
	_ "modernc.org/sqlite"
)

func main() {
	fmt.Print("안녕")
}
