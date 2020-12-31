package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"launchpad.net/xmlpath"
)

func printSakka(personhtmlpath string) {
	fp, err := os.Open(personhtmlpath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	path := xmlpath.MustCompile(`/html/body/table/tr[1]/td[2]`)
	node, err := xmlpath.ParseHTML(fp)
	if err != nil {
		// euc-jpで死ぬ
		return
	}
	if text, ok := path.String(node); ok {
		fmt.Println(text)
	}
	return
}

func personfilelist() {
	dir := "aozorabunko-master/index_pages"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("error: aozorabunkoをダウンロードしてください %s", err.Error())
	}
	pattern := regexp.MustCompile(`person[0-9]+\.html`)
	for _, file := range files {
		if pattern.MatchString(file.Name()) {
			printSakka(filepath.Join(dir, file.Name()))
		}
	}
}

func main() {
	personfilelist()
	printSakka("aozorabunko-master/index_pages/person1117.html")
}
