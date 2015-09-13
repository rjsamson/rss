## RSS
### A simple package for dealing with RSS in Go

Nothing too fancy here - just give rss.Parse an `io.Reader` containing RSS and it returns a struct (`rss.Feed`) representing the feed and all items.

### Usage

`go get github.com/rjsamson/rss`

```go
feed, err := rss.Parse(xmlData)
```

## Structure

### rss.Feed
```go
type Feed struct {
	Author      string
	Id          string
	Image       Image
	Language    string
	Generator   string
	Link        string
	Subtitle    string
	Summary     string
	Title       string
	Description string
	Copyright   string
	Items       []Item
}
```

### rss.Image
```go
type Image struct {
	Url         string
	Title       string
	Link        string
	Description string
	Width       string
	Height      string
}
```

### rss.Item
```go
type Item struct {
	Title       string
	Author      string
	Id          string
	Image       Image
	Duration    string
	Link        string
	Subtitle    string
	Summary     string
	Description string
	PubDate     string
	Enclosure   Enclosure
}
```

### rss.Enclosure
```go
type Enclosure struct {
	Url    string
	Length string
	Type   string
}
```

## Example

```go
package main

import (
	"fmt"
	"github.com/rjsamson/rss"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.feedforall.com/sample.xml")

	if err != nil {
		log.Fatal(err)
		return
	}

	xmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	feed, err := rss.Parse(xmlData)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(feed.Title)

	for _, item := range feed.Items {
		fmt.Println("\nTitle:", item.Title)
		fmt.Println("  URL:", item.Link)
	}
}
```
