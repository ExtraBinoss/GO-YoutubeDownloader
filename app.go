package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DownloadVideo(url string, format string) string {
	format = strings.ToLower(format)
	fmt.Println("Adding to queue : ", url, format)
	//put all format string in lowercase
	cmd := exec.Command("yt-dlp", "-f", format, url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error downloading video : ", err)
		return "0"
	}

	// Convert output to string and check for "100%"
	outputStr := string(output)
	fmt.Println(outputStr)
	if strings.Contains(outputStr, "100%") {
		fmt.Println("I entered here")
		return "1"
	}
	return "0"
}
