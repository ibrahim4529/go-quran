package state

import "github.com/ibrahim4529/go-quran/internal/model"

type MainWindowState struct {
	Height        int                         `json:"height"`
	Width         int                         `json:"width"`
	FullScreen    bool                        `json:"full_screen"`
	QuranState    QuranState                  `json:"quran_state"`
	ListTrans     []model.TranslationLanguage `json:"list_trans"`
	ListQuranType []model.QuranType           `json:"list_quran_type"`
}

type QuranState struct {
	QuranType       model.QuranType           `json:"quran_type"`
	TransaltionLang model.TranslationLanguage `json:"transaltion_lang"`
	LastAya         int                       `json:"last_aya"`
	LastSura        int                       `json:"last_sura"`
}

func MainWindowStateNew(height, width int) *MainWindowState {
	return &MainWindowState{
		Height:     height,
		Width:      width,
		FullScreen: false,
	}
}
