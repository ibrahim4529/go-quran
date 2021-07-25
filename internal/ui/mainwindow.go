package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/ibrahim4529/go-quran/internal/model"
	"github.com/ibrahim4529/go-quran/internal/state"
)

type MainWindow struct {
	*gtk.ApplicationWindow
	WindowState state.MainWindowState
	Quran       *model.Quran
	Translation *model.Translation
	CssProvider *gtk.CssProvider
}

func MainWindowNew(app *gtk.Application, state *state.MainWindowState) *MainWindow {
	window, _ := gtk.ApplicationWindowNew(app)
	cssProvider, _ := gtk.CssProviderNew()
	quran := model.QuranFromFile(state.QuranState.QuranType.Path)
	translation := model.TranslationFromFile(state.QuranState.TransaltionLang.Path)
	window.SetTitle("Go-Quran")
	window.SetDefaultSize(state.Width, state.Height)
	window.SetPosition(gtk.WIN_POS_CENTER)

	return &MainWindow{
		ApplicationWindow: window,
		Quran:             quran,
		Translation:       translation,
		WindowState:       *state,
		CssProvider:       cssProvider,
	}
}

func (win *MainWindow) StartReadMode() {

	mainPaned, _ := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	scrolledWindow, _ := gtk.ScrolledWindowNew(nil, nil)
	scrolledWindow.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)

	sideBar := SideBarNew()
	mainContent := SurahWidgetNew()
	scrolledWindow.Add(mainContent)

	mainPaned.Pack1(sideBar, false, false)
	mainPaned.SetPosition(200)
	mainPaned.Pack2(scrolledWindow, false, false)

	sideBar.Init(win)

	sideBar.CbSura.Connect("changed", func(box *gtk.ComboBoxText) {
		treeIter, _ := box.GetActiveIter()
		index := box.GetActive()
		if treeIter != nil {
			win.WindowState.QuranState.LastSura = index + 1
			mainContent.RenderSurah(index, win.Quran, win.Translation, sideBar)
		}
	})

	sideBar.CbAya.Connect("changed", func(box *gtk.ComboBoxText) {
		iter, _ := box.GetActiveIter()
		index := box.GetActive()
		if iter != nil {
			win.WindowState.QuranState.LastAya = index + 1
			mainContent.GoToAya(index)
		}
	})

	mainContent.Connect("row_selected", func() {
		row := mainContent.GetSelectedRow()
		if row != nil {
			index := row.GetIndex()
			sideBar.CbAya.SetActive(index)
		}
	})

	sideBar.CbTranslation.Connect("changed", func(box *gtk.ComboBoxText) {
		iter, _ := box.GetActiveIter()

		if iter != nil {
			modelBox, _ := box.GetModel()
			val, _ := modelBox.ToTreeModel().GetValue(iter, 1)
			path, _ := val.GetString()
			win.Translation = model.TranslationFromFile(path)
		}
		mainContent.ReRender(win)
	})

	win.Add(mainPaned)
	win.ShowAll()
}
