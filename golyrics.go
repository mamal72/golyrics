package golyrics

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

const searchBaseURI = "http://lyrics.wikia.com/index.php?action=ajax&rs=getLinkSuggest&format=json&query="
const lyricsBaseURI = "http://lyrics.wikia.com/wiki/"

// Track is a music track containing Artist, Name and Lyrics.
type Track struct {
	Artist string
	Name   string
	Lyrics string
}

// FetchLyrics fetches the lyrics of a Track and sets it on that track.
func (track *Track) FetchLyrics() error {
	URI := fmt.Sprintf("%s%s:%s", lyricsBaseURI, track.Artist, track.Name)
	doc, err := goquery.NewDocument(URI)
	if err != nil {
		return err
	}

	lyricsHTML, err := doc.Find(".lyricbox").Html()
	if err != nil {
		return err
	}

	track.Lyrics = getFormattedLyrics(lyricsHTML)
	return nil
}

func breakToNewLine(HTML string) string {
	return strings.Replace(HTML, "<br/>", "\n", -1)
}

func stripeHTMLTags(HTML string) string {
	regex, _ := regexp.Compile("<[^>]+>")
	return regex.ReplaceAllString(HTML, "")
}

func fixApostrophes(text string) string {
	return strings.Replace(text, "&#39;", "'", -1)
}

func getSearchURI(query string) string {
	return fmt.Sprintf("%s%s", searchBaseURI, url.QueryEscape(query))
}

func getFormattedLyrics(text string) string {
	noBreaks := breakToNewLine(text)
	noHTMLTags := stripeHTMLTags(noBreaks)
	return fixApostrophes(noHTMLTags)
}

// SearchTrack searches for tracks
// using a string query that can be part of the track name or artist.
func SearchTrack(query string) ([]Track, error) {
	requestURI := getSearchURI(query)
	response, err := http.Get(requestURI)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	suggestions := []Track{}
	jsonparser.ArrayEach(data, func(value []byte, _ jsonparser.ValueType, offset int, _ error) {
		title := string(value[:])
		trackParts := strings.SplitN(title, ":", 2)
		if len(trackParts) < 2 {
			return
		}
		track := Track{
			Artist: trackParts[0],
			Name:   trackParts[1],
		}
		suggestions = append(suggestions, track)
	}, "suggestions")

	return suggestions, nil
}

// SearchTrackByArtistAndName searches for tracks
// using artist and name of the track.
func SearchTrackByArtistAndName(artist string, name string) ([]Track, error) {
	return SearchTrack(artist + ":" + name)
}
