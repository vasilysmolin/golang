package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Locale map[string]string

var Langs []string

var CurLocale map[string]Locale

func GetLocales() {
	// Загрузка переводов из файла
	translations := make(map[string]Locale)
	langs := []string{}

	files, err := ioutil.ReadDir("./locales")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && filepath.Ext(name) == ".json" {
			lang := strings.TrimSuffix(name, ".json")
			jsonFile, err := os.Open(filepath.Join("./locales", name))
			if err != nil {
				log.Fatal(err)
			}
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			var translation Locale
			json.Unmarshal(byteValue, &translation)

			translations[lang] = translation
			langs = append(langs, lang)
		}
	}
	CurLocale = translations
	Langs = langs
}
