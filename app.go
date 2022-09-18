package main

import (
	"context"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"os"
	"runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) MarkdownerFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}

	out := bluemonday.UGCPolicy().SanitizeBytes(blackfriday.Run(file, blackfriday.WithExtensions(blackfriday.CommonExtensions)))

	return string(out)
}

func (a *App) ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return string(file)
}

func (a *App) ReadDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		return []string{err.Error()}
	}
	var out []string
	for _, file := range files {
		out = append(out, file.Name())
	}
	return out
}

func (a *App) GetPlatform() string {
	return runtime.GOOS
}
