package main

import (
	"github.com/giefferre/gettext"
	"log"
)

func main() {
	languages := gettext.NewCollection()
	err := languages.LoadDirectory(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(languages.Get("fr").Translate("Dial %s?", "asd"))
}
