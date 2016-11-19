[![Build Status](https://travis-ci.org/mamal72/golyrics.svg?branch=master)](https://travis-ci.org/mamal72/golyrics)
[![codecov](https://codecov.io/gh/mamal72/golyrics/branch/master/graph/badge.svg)](https://codecov.io/gh/mamal72/golyrics)
[![GoDoc](https://godoc.org/github.com/mamal72/golyrics?status.svg)](https://godoc.org/github.com/mamal72/golyrics)

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
    // Get lyrics suggestions by searching
    suggestions, err := golyrics.SearchLyrics("Blackfield Some Day") // []string, error
    // OR
    suggestions, err := golyrics.SearchLyricsByArtistAndName("Blackfield", "Some Day") // []string, error

    // Now fetch the lyrics
    if len(suggestions) == 0 {
        // No lyrics found for this track :(
        // Try some other keywords or show some error to user
    }
    lyrics, err := golyrics.GetLyrics(suggestions[0]) // string, error


    // You can also search and fetch the lyrics with only one call
    // It'll use the first search result for fetching lyrics 
    lyrics, err := golyrics.SearchAndGetLyrics("Blackfield Some Day") // string, error
    // OR
    lyrics, err := golyrics.SearchAndGetLyricsByArtistAndName("Blackfield", "Some Day") // string, error
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