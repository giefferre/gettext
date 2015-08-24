package gettext

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_LANGUAGE = "en"
)

type Collection struct {
	catalogs        map[string]*Catalog
	defaultLanguage string
}

func NewCollection() *Collection {
	return &Collection{
		catalogs:        map[string]*Catalog{},
		defaultLanguage: "en",
	}
}

func (c *Collection) LoadDirectory(path string) error {
	directoryPath := fmt.Sprintf("%s%s*.mo", path, string(os.PathSeparator))
	files, err := filepath.Glob(directoryPath)
	if err != nil {
		return err
	}

	c.catalogs = make(map[string]*Catalog, 0)

	for _, fileName := range files {
		language := strings.ToLower(fmt.Sprintf(strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))))

		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}

		catalog := NewCatalog()
		if err := catalog.ReadMo(bytes.NewReader(fileBytes)); err != nil {
			return err
		}

		c.catalogs[language] = catalog
	}

	return nil
}

func (c *Collection) SetDefaultLanguage(langCode string) {
	c.defaultLanguage = langCode
}

func (c *Collection) Get(langCode string) *Catalog {
	langCode = strings.ToLower(langCode)
	catalog, ok := c.catalogs[langCode]
	if ok {
		return catalog
	}

	catalog, ok = c.catalogs[strings.Split(langCode, "-")[0]]
	if ok {
		return catalog
	}

	catalog, ok = c.catalogs[c.defaultLanguage]
	if ok {
		return catalog
	}

	return NewCatalog()
}
