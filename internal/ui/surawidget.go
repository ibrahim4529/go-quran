package ui

import (
	"strconv"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ibrahim4529/go-quran/internal/model"
)

type SurahWidget struct {
	*gtk.ListBox
}

func SurahWidgetNew() *SurahWidget {
	listView, _ := gtk.ListBoxNew()
	listView.SetSelectionMode(gtk.SELECTION_SINGLE)
	return &SurahWidget{
		ListBox: listView,
	}
}

func (sw *SurahWidget) RenderSurah(index int, quran *model.Quran, trans *model.Translation, sb *SideBar) {
	go func() {
		glib.IdleAdd(func() {
			sb.AyaListore.Clear()
			listAyaContent := sw.GetChildren()
			listAyaContent.Foreach(func(item interface{}) {
				sw.Remove(item.(gtk.IWidget))
			})
			for i, content := range quran.Suras[index].Ayas {
				sb.AyaListore.Set(sb.AyaListore.Append(), []int{0}, []interface{}{strconv.Itoa(i + 1)})
				ayaWidet := AyaWidgetNew(quran.Suras[index].Name, content.Text, trans.Suras[index].Ayas[i].Text, i)
				sw.Add(ayaWidet)
			}
			sw.SelectAll()
			iter, _ := sb.AyaListore.GetIterFirst()
			sb.CbAya.SetActiveIter(iter)
		})
	}()
}

func (sw *SurahWidget) GoToAya(index int) {
	row := sw.GetRowAtIndex(index)
	sw.SelectRow(row)
	if _, y, _ := row.TranslateCoordinates(sw, 0, 0); y >= 0 {
		if adj := sw.GetAdjustment(); adj != nil {
			_, rowHeight := row.GetPreferredHeight()
			adj.SetValue(float64(y) - (adj.GetPageSize()-float64(rowHeight))/2)
		}
	}
}

func (sw *SurahWidget) ReRender(win *MainWindow) {
	go func() {
		glib.IdleAdd(func() {
			listAyaContent := sw.GetChildren()
			listAyaContent.Foreach(func(item interface{}) {
				sw.Remove(item.(gtk.IWidget))
			})
			lastSura := win.WindowState.QuranState.LastSura - 1
			lastAya := win.WindowState.QuranState.LastAya - 1
			for i, content := range win.Quran.Suras[lastSura].Ayas {
				// sb.AyaListore.Set(sb.AyaListore.Append(), []int{0}, []interface{}{strconv.Itoa(i + 1)})
				ayaWidet := AyaWidgetNew(win.Quran.Suras[lastSura].Name, content.Text, win.Translation.Suras[lastSura].Ayas[i].Text, i)
				sw.Add(ayaWidet)
			}
			sw.GoToAya(lastAya)
		})
	}()
}
