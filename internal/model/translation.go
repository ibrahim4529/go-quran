package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Translation struct {
	Suras []Sura `xml:"sura"`
}

type TranslationLanguage struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewTransLationLanguage(name, path string) *TranslationLanguage {
	return &TranslationLanguage{
		Name: name,
		Path: path,
	}
}

func TranslationFromFile(path string) *Translation {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var translation Translation
	xml.Unmarshal(byteValue, &translation)
	return &translation
}
