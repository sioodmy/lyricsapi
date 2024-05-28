package get

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

var cache = make(map[string]string)

func GetHandle(w http.ResponseWriter, r *http.Request) {
	query := r.PathValue("query")
	lyrics, err := QueryLyrics(query)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write([]byte(lyrics))
}

func QueryLyrics(query string) (string, error) {
	song_path, err := SearchSong(query)
	lyrics := GetSong(song_path)

	return lyrics, err
}

func SearchSong(query string) (string, error) {
	url := fmt.Sprintf("https://www.tekstowo.pl/wyszukaj.html?search-query=%s", query)

	c := colly.NewCollector()
	var song_path string

	found := false

	echan := make(chan error, 1)
	c.OnHTML("div.card.mb-4 a[href]", func(e *colly.HTMLElement) {
		if !found {
			song_path = e.Attr("href")
			found = true
		}
	})
	c.OnScraped(func(r *colly.Response) {

		if !found {
			echan <- fmt.Errorf("Couldn't find any matching songs")

		}
	})

	go func() {
		c.Visit(url)
		close(echan)
	}()

	err := <-echan

	return song_path, err
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
		cache[path] = lyrics
	})

	c.Visit(url)
	return lyrics
}
