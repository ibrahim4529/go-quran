package main

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ibrahim4529/go-quran/internal/model"
	"github.com/ibrahim4529/go-quran/internal/state"
	"github.com/ibrahim4529/go-quran/internal/ui"
)

func main() {
	app, _ := gtk.ApplicationNew("com.github.ibrahim4529.go-quran", glib.APPLICATION_FLAGS_NONE)

	app.Connect("activate", func() {
		activate(app)
	})

	app.Run(os.Args)
}

func activate(app *gtk.Application) {
	windowState := state.MainWindowStateNew(600, 1000)
	windowState.ListQuranType = []model.QuranType{
		{
			Name: "Utmani",
			Path: "res/quran-uthmani.xml",
		},
		{
			Name: "Simple Minimal",
			Path: "res/quran-simple-min.xml",
		},
	}
	windowState.ListTrans = []model.TranslationLanguage{
		{
			Name: "Indonesia",
			Path: "res/id.indonesian.xml",
		},
		{
			Name: "English Ahmedali",
			Path: "res/en.ahmedali.xml",
		},
	}
	windowState.QuranState = state.QuranState{
		QuranType: model.QuranType{
			Name: "Utmani",
			Path: "res/quran-uthmani.xml",
		},
		TransaltionLang: model.TranslationLanguage{
			Name: "Indonesia",
			Path: "res/id.indonesian.xml",
		},
		LastAya:  1,
		LastSura: 1,
	}
	window := ui.MainWindowNew(app, windowState)
	window.CssProvider.LoadFromPath("res/style.css")
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, window.CssProvider, 1)
	window.StartReadMode()
	window.ShowAll()
}
