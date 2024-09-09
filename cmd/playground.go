package cmd

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func Playground() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
	// 	ExecutablePath: playwright.String("C:/Program Files/Google/Chrome/Application/chrome.exe"),
	// })

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless:       playwright.Bool(false),
		ExecutablePath: playwright.String("C:/Program Files/Google/Chrome/Application/chrome.exe"),
	})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}
