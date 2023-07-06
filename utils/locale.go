package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Locale map[string]string

var Langs []string

var CurLocale map[string]Locale

func GetLocales(root string) {
	// Загрузка переводов из файла
	translations := make(map[string]Locale)
	langs := []string{}

	files, err := ioutil.ReadDir(root + "/locales")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && filepath.Ext(name) == ".json" {
			lang := strings.TrimSuffix(name, ".json")
			jsonFile, err := os.Open(filepath.Join(root+"/locales", name))
			if err != nil {
				log.Fatal(err)
			}
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			var translation Locale
			err = json.Unmarshal(byteValue, &translation)
			if err != nil {
				logrus.Fatal(err)
			}

			translations[lang] = translation
			langs = append(langs, lang)
		}
	}
	CurLocale = translations
	Langs = langs
}
