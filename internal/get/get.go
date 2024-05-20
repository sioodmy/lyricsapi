package get

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

var cache = make(map[string]string)

func GetHandle(w http.ResponseWriter, r *http.Request) {
	query := r.PathValue("query")
	song_path := SearchSong(query)
	lyrics := GetSong(song_path)

	w.Write([]byte(lyrics))
}

func SearchSong(query string) string {

	url := fmt.Sprintf("https://www.tekstowo.pl/wyszukaj.html?search-query=%s", query)

	c := colly.NewCollector()
	var song_path string

	found := false
	c.OnHTML("div.card.mb-4 a[href]", func(e *colly.HTMLElement) {
		if !found {
			song_path = e.Attr("href")
			found = true
		}
	})
	c.Visit(url)

	return song_path
}

func GetSong(path string) string {
	url := fmt.Sprintf("https://www.tekstowo.pl%s", path)
	c := colly.NewCollector()

	cached, exists := cache[path]

	if exists {
		return cached
	}

	var lyrics string
	c.OnHTML("div#songText div.inner-text", func(e *colly.HTMLElement) {
		lyrics = e.Text
	})

	cache[path] = lyrics

	c.Visit(url)
	return lyrics
}
