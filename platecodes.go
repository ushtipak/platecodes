package main

import (
	"github.com/antchfx/htmlquery"
	"bufio"
	"os"
)

var (
	url      = "https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Serbia"
	xPath    = "//*[@id=\"mw-content-text\"]/div/table[1]/tbody/tr[*]/td[1]"
	codeFile = "/tmp/platecodes"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	doc, err := htmlquery.LoadURL(url)
	check(err)

	f, err := os.Create(codeFile)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, n := range htmlquery.Find(doc, xPath) {
		_, err := w.WriteString(htmlquery.InnerText(n) + "\n")
		check(err)
	}
	w.Flush()
}
