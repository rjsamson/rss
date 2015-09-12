package rss

import (
	"bytes"
	"encoding/xml"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"io"
	"io/ioutil"
)

type Feed struct {
	XMLName     xml.Name `xml:"rss"`
	Author      string   `xml:"channel>author"`
	Id          string   `xml:"channel>id"`
	Image       Image    `xml:"channel>image"`
	Language    string   `xml:"channel>language"`
	Generator   string   `xml:"channel>generator"`
	Link        string   `xml:"channel>link"`
	Subtitle    string   `xml:"channel>subtitle"`
	Summary     string   `xml:"channel>summary"`
	Title       string   `xml:"channel>title"`
	Description string   `xml:"channel>description"`
	Copyright   string   `xml:"channel>copyright"`
	Items       []Item   `xml:"channel>item"`
}

type Image struct {
	Url         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Width       string `xml:"width"`
	Height      string `xml:"height"`
}

type Item struct {
	Title       string    `xml:"title"`
	Author      string    `xml:"author"`
	Id          string    `xml:"id"`
	Image       Image     `xml:"image"`
	Duration    string    `xml:"duration"`
	Link        string    `xml:"link"`
	Subtitle    string    `xml:"subtitle"`
	Summary     string    `xml:"summary"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

func Parse(reader io.Reader) (Feed, error) {
	var result Feed
	xmlData, err := ioutil.ReadAll(reader)

	if err != nil {
		return result, err
	}

	newReader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(newReader)
	decoder.CharsetReader = charset.NewReader
	err = decoder.Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
