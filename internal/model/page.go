package model

import (
	"encoding/xml"
	"fmt"
	"os"
)

type QuranPage struct {
	Pages []struct {
		Index int `xml:"index,attr"`
		Sura  int `xml:"sura,attr"`
		Aya   int `xml:"aya,attr"`
	} `xml:"page"`
}

func QuranPageFromFile(path string) *QuranPage {
	xmlText, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error Bos")
	}
	var quranPage QuranPage
	xml.Unmarshal(xmlText, &quranPage)
	return &quranPage
}
