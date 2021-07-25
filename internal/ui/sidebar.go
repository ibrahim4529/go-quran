package ui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type SideBar struct {
	*gtk.Box
	AyaListore    *gtk.ListStore
	SuraListore   *gtk.ListStore
	TransListore  *gtk.ListStore
	CbSura        *gtk.ComboBoxText
	CbAya         *gtk.ComboBoxText
	CbTranslation *gtk.ComboBoxText
}

func SideBarNew() *SideBar {
	listStoreSura, _ := gtk.ListStoreNew(glib.TYPE_STRING)
	listStoreIndexAya, _ := gtk.ListStoreNew(glib.TYPE_STRING)
	listStoreTrans, _ := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	vbox.SetMarginBottom(8)
	vbox.SetMarginStart(8)
	vbox.SetMarginTop(8)
	vbox.SetMarginEnd(8)
	frameSeach, _ := gtk.FrameNew("Search Surah")
	frameSeach.SetMarginBottom(8)
	frameSeach.SetMarginStart(8)
	frameSeach.SetMarginTop(8)
	frameSeach.SetMarginEnd(8)

	gridSearch, _ := gtk.GridNew()
	gridSearch.SetRowSpacing(8)
	gridSearch.SetMarginBottom(8)
	gridSearch.SetMarginStart(8)
	gridSearch.SetMarginTop(8)
	gridSearch.SetMarginEnd(8)

	frameSeach.Add(gridSearch)
	labelSearchSura, _ := gtk.LabelNew("Sura Name: ")
	labelSearchSura.SetHAlign(gtk.ALIGN_START)
	labelSearchAya, _ := gtk.LabelNew("Aya: ")
	labelSearchAya.SetHAlign(gtk.ALIGN_START)

	gridSearch.Attach(labelSearchSura, 0, 0, 1, 1)
	gridSearch.Attach(labelSearchAya, 0, 1, 1, 1)

	cbSura, _ := gtk.ComboBoxTextNewWithEntry()
	cbSura.SetHExpand(true)
	cbSura.SetModel(listStoreSura)
	cbAya, _ := gtk.ComboBoxTextNewWithEntry()
	cbAya.SetHExpand(true)
	cbAya.SetModel(listStoreIndexAya)
	gridSearch.Attach(cbSura, 1, 0, 2, 1)
	gridSearch.Attach(cbAya, 1, 1, 2, 1)

	transFrame, _ := gtk.FrameNew("Configuration")
	transFrame.SetMarginBottom(8)
	transFrame.SetMarginStart(8)
	transFrame.SetMarginTop(8)
	transFrame.SetMarginEnd(8)

	gridTrans, _ := gtk.GridNew()
	gridTrans.SetRowSpacing(8)
	gridTrans.SetMarginBottom(8)
	gridTrans.SetMarginStart(8)
	gridTrans.SetMarginTop(8)
	gridTrans.SetMarginEnd(8)
	fontLabel, _ := gtk.LabelNew("Font: ")
	transLabel, _ := gtk.LabelNew("Translation: ")
	fontLabel.SetHAlign(gtk.ALIGN_START)
	transLabel.SetHAlign(gtk.ALIGN_START)
	cbTrans, _ := gtk.ComboBoxTextNewWithEntry()
	cbTrans.SetModel(listStoreTrans)
	cbFont, _ := gtk.FontButtonNewWithFont("LPMQ Isep Misbah 18")
	gridTrans.Attach(fontLabel, 0, 0, 1, 1)
	gridTrans.Attach(cbFont, 1, 0, 2, 1)
	gridTrans.Attach(transLabel, 0, 1, 1, 1)
	gridTrans.Attach(cbTrans, 1, 1, 2, 1)
	transFrame.Add(gridTrans)
	vbox.Add(frameSeach)
	vbox.Add(transFrame)

	return &SideBar{
		Box:           vbox,
		AyaListore:    listStoreIndexAya,
		SuraListore:   listStoreSura,
		TransListore:  listStoreTrans,
		CbSura:        cbSura,
		CbAya:         cbAya,
		CbTranslation: cbTrans,
	}
}

func (sb *SideBar) Init(win *MainWindow) {
	go func() {
		glib.IdleAdd(func() {
			for _, aya := range win.Quran.Suras {
				err := sb.SuraListore.Set(sb.SuraListore.Append(), []int{0}, []interface{}{aya.Name})
				if err != nil {
					return
				}
			}
			iter, _ := sb.SuraListore.GetIterFirst()
			sb.CbSura.SetActiveIter(iter)
		})
	}()

	go func() {
		glib.IdleAdd(func() {
			for _, trans := range win.WindowState.ListTrans {
				sb.TransListore.Set(sb.TransListore.Append(), []int{0, 1}, []interface{}{trans.Name, trans.Path})
			}
			iter, _ := sb.TransListore.GetIterFirst()
			sb.CbTranslation.SetActiveIter(iter)
		})
	}()
}
