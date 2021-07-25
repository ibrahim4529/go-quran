package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Aya struct {
	Index int    `xml:"index,attr"`
	Text  string `xml:"text,attr"`
}

type Sura struct {
	Index int    `xml:"index,attr"`
	Name  string `xml:"name,attr"`
	Ayas  []Aya  `xml:"aya"`
}
type Quran struct {
	Suras []Sura `xml:"sura"`
}

type QuranType struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func QuranFromFile(path string) *Quran {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var quran Quran
	xml.Unmarshal(byteValue, &quran)
	return &quran
}
