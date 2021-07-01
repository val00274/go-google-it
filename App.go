package googleit

import (
	"net/url"
	"strings"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) QueryEscape(keywords ...string) string {
	return url.QueryEscape(strings.Join(keywords, " "))
}

func (a *App) GoogleURL(keywords ...string) string {
	return "https://google.co.jp/search?q=" + a.QueryEscape(keywords...)
}
