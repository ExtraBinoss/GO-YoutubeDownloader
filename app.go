package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gitlab.com/toby3d/oembed"
)

func (a *App) GetVideoAuthor(url string) string {
	// Get the video author
	data, err := oembed.Extract(url, nil)
	if err != nil {
		fmt.Println("Error getting author : ", err)
		return ""
	}
	author := data.AuthorName
	return author
}

func (a *App) GetVideoTitle(url string) string {
	// Get the video title
	data, err := oembed.Extract(url, nil)
	if err != nil {
		fmt.Println("Error getting title : ", err)
		return ""
	}
	title := data.Title
	return title
}

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

	outputStr := string(output)
	fmt.Println(outputStr)
	os.WriteFile("output.txt", []byte(outputStr), 0644)
	if strings.Contains(outputStr, "100%") {
		return "1"
	}
	return "0"
}

//write to a log file the output of the download, string as parameter, returns void
func (a *App) WriteALogFile(output string) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println(output)
	file.Close()
}
