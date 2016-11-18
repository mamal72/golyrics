package golyrics

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

const searchBaseURI = "http://lyrics.wikia.com/index.php?action=ajax&rs=getLinkSuggest&format=json&query="
const lyricsBaseURI = "http://lyrics.wikia.com/wiki/"

func breakToNewLine(HTML string) string {
	return strings.Replace(HTML, "<br/>", "\n", -1)
}

func stripeHTMLTags(HTML string) string {
	regex, err := regexp.Compile("<[^>]+>")
	if err != nil {
		log.Fatal(err)
	}
	return regex.ReplaceAllString(HTML, "")
}

func fixApostrophes(text string) string {
	return strings.Replace(text, "&#39;", "'", -1)
}

func getSearchURI(query string) string {
	return searchBaseURI + query
}

func getFormattedLyrics(text string) string {
	return fixApostrophes(stripeHTMLTags(breakToNewLine(text)))
}

// SearchLyrics searches for lyrics
// using a string query that can be part of the track name or artist
func SearchLyrics(query string) []string {
	requestURI := getSearchURI(query)
	response, err := http.Get(requestURI)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var suggestions []string
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if (err) != nil {
			log.Fatal(err)
		}
		if dataType != jsonparser.String {
			log.Fatal("error searching lyrics")
		}
		suggestions = append(suggestions, string(value[:]))
	}, "suggestions")

	return suggestions
}

// SearchLyricsByArtistAndName searches for lyrics
// using artist and name of the track
func SearchLyricsByArtistAndName(artist string, name string) []string {
	return SearchLyrics(artist + ":" + name)
}

// GetLyrics scrapes the lyrics of the given title
// The title should be a search result response
func GetLyrics(title string) string {
	URI := lyricsBaseURI + title
	doc, err := goquery.NewDocument(URI)
	if err != nil {
		log.Fatal(err)
	}

	lyricsHTML, err := doc.Find(".lyricbox").Html()
	if err != nil {
		log.Fatal(err)
	}

	return getFormattedLyrics(lyricsHTML)
}

// SearchAndGetLyrics searches for lyrics by a query
// and returns the lyrics of the first search result
func SearchAndGetLyrics(query string) string {
	suggestions := SearchLyrics(query)
	if len(suggestions) == 0 {
		return "No lyrics found for " + query
	}
	return GetLyrics(suggestions[0])
}

// SearchAndGetLyricsByArtistAndName searches for lyrics by artist and name
// and returns the lyrics of the first search result
func SearchAndGetLyricsByArtistAndName(artist string, name string) string {
	suggestions := SearchLyrics(artist + ":" + name)
	if len(suggestions) == 0 {
		return "No lyrics found for " + artist + " - " + name
	}
	return GetLyrics(suggestions[0])
}
