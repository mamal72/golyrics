package golyrics

import (
	"errors"
	"io/ioutil"
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

func stripeHTMLTags(HTML string) (string, error) {
	regex, err := regexp.Compile("<[^>]+>")
	if err != nil {
		return "", err
	}
	return regex.ReplaceAllString(HTML, ""), nil
}

func fixApostrophes(text string) string {
	return strings.Replace(text, "&#39;", "'", -1)
}

func getSearchURI(query string) string {
	return searchBaseURI + query
}

func getFormattedLyrics(text string) (string, error) {
	noBreaks := breakToNewLine(text)
	noHTMLTags, err := stripeHTMLTags(noBreaks)
	if err != nil {
		return "", err
	}
	return fixApostrophes(noHTMLTags), nil
}

// SearchLyrics searches for lyrics
// using a string query that can be part of the track name or artist
func SearchLyrics(query string) ([]string, error) {
	requestURI := getSearchURI(query)
	response, err := http.Get(requestURI)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	suggestions := []string{}
	jsonparser.ArrayEach(data, func(value []byte, _ jsonparser.ValueType, offset int, _ error) {
		suggestions = append(suggestions, string(value[:]))
	}, "suggestions")

	return suggestions, nil
}

// SearchLyricsByArtistAndName searches for lyrics
// using artist and name of the track
func SearchLyricsByArtistAndName(artist string, name string) ([]string, error) {
	return SearchLyrics(artist + ":" + name)
}

// GetLyrics scrapes the lyrics of the given title
// The title should be a search result response
func GetLyrics(title string) (string, error) {
	URI := lyricsBaseURI + title
	doc, err := goquery.NewDocument(URI)
	if err != nil {
		return "", err
	}

	lyricsHTML, err := doc.Find(".lyricbox").Html()
	if err != nil {
		return "", err
	}

	return getFormattedLyrics(lyricsHTML)
}

// SearchAndGetLyrics searches for lyrics by a query
// and returns the lyrics of the first search result
func SearchAndGetLyrics(query string) (string, error) {
	suggestions, err := SearchLyrics(query)
	if err != nil {
		return "", err
	}
	if len(suggestions) == 0 {
		return "", errors.New("No lyrics found for " + query)
	}
	return GetLyrics(suggestions[0])
}

// SearchAndGetLyricsByArtistAndName searches for lyrics by artist and name
// and returns the lyrics of the first search result
func SearchAndGetLyricsByArtistAndName(artist string, name string) (string, error) {
	suggestions, err := SearchLyrics(artist + ":" + name)
	if err != nil {
		return "", err
	}
	if len(suggestions) == 0 {
		return "", errors.New("No lyrics found for " + artist + " - " + name)
	}
	return GetLyrics(suggestions[0])
}
