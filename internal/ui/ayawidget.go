package ui

import (
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

type AyaWidget struct {
	*gtk.Frame

	SurahName       string
	AyahIndex       int
	ArabicText      string
	TranslationText string
}

func AyaWidgetNew(surahName, arabic, translation string, ayaIndex int) *AyaWidget {
	frame, _ := gtk.FrameNew(surahName + toArabicNumberText(ayaIndex+1))
	frame.SetMarginBottom(8)
	frame.SetMarginStart(8)
	frame.SetMarginTop(8)
	frame.SetMarginEnd(8)
	label, _ := gtk.LabelNew(arabic + toArabicNumberText(ayaIndex+1))
	cssContext, _ := label.GetStyleContext()
	cssContext.AddClass("quran")
	label.SetHAlign(gtk.ALIGN_END)
	tarjemah, _ := gtk.LabelNew(translation)
	tarjemah.SetHAlign(gtk.ALIGN_START)
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	vbox.SetMarginBottom(8)
	vbox.SetMarginStart(8)
	vbox.SetMarginTop(8)
	vbox.SetMarginEnd(8)
	vbox.Add(label)
	vbox.Add(tarjemah)

	label.SetLineWrap(true)
	label.SetLineWrapMode(pango.WRAP_WORD)
	tarjemah.SetLineWrap(true)
	tarjemah.SetLineWrapMode(pango.WRAP_WORD)
	frame.Add(vbox)
	frame.ShowAll()
	return &AyaWidget{
		Frame:           frame,
		SurahName:       surahName,
		TranslationText: translation,
		AyahIndex:       ayaIndex,
	}
}

func toArabicNumberText(number int) string {
	alphabet := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	arabic := []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"}
	numberStr := strconv.Itoa(number)
	numberStr = "﴿" + numberStr + "﴾"
	for i := range alphabet {
		numberStr = strings.ReplaceAll(numberStr, alphabet[i], arabic[i])
	}
	return numberStr
}
