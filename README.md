[![Build Status](https://travis-ci.org/mamal72/golyrics.svg?branch=master)](https://travis-ci.org/mamal72/golyrics)
[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/golyrics)](https://goreportcard.com/report/github.com/mamal72/golyrics)
[![Coverage Status](https://coveralls.io/repos/github/mamal72/golyrics/badge.svg?branch=master)](https://coveralls.io/github/mamal72/golyrics?branch=master)
[![GoDoc](https://godoc.org/github.com/mamal72/golyrics?status.svg)](https://godoc.org/github.com/mamal72/golyrics)
[![license](https://img.shields.io/github/license/mamal72/golyrics.svg)](https://github.com/mamal72/golyrics/blob/master/LICENSE)

# golyrics

This is a simple scrapper package to fetch lyrics data from the [Wikia](http://lyrics.wikia.com) website.


## Installation

```bash
go get github.com/mamal72/golyrics
```


## Usage

```go
package main

import "github.com/mamal72/golyrics"

func main() {
    // Get track suggestions by searching
    suggestions, err := golyrics.SearchTrack("Blackfield Some Day") // []Track, error
    // OR
    suggestions, err := golyrics.SearchTrackByArtistAndName("Blackfield", "Some Day") // []Track, error

    // Let's check results
    if err != nil || len(suggestions) == 0 {
        // No lyrics found for this track :(
        // Try some other keywords or show some error to user
    }

    // Assign first result to the track
    track := suggestions[0] // Track

    // Now fetch the lyrics and set it back on the track    
    err := track.FetchLyrics() // error
    if err != nil {
        // Error fetching lyrics for the track
    }
    fmt.Printf("Lyrics of %s by %s: %s", track.Name, track.Artist, track.Lyrics)
}
```


## Tests

```bash
go test
```


## Ideas || Issues
Just fill an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. :heart:

Remember to write a few tests for your code before sending pull requests.


## License
> MIT