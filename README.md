# golyrics

This is a simple scrapper package for fetching lyrics data from [Wikia](http://lyrics.wikia.com).


## Installation:

```bash
go get github.com/mamal72/golyrics
```


## Usage

```go
package main

import "github.com/mamal72/golyrics"

func main() {
    // Get lyrics suggestions by searching
    suggestions := golyrics.SearchLyrics("Blackfield Some Day") // []string
    // OR
    suggestions := golyrics.SearchLyricsByArtistAndName("Blackfield", "Some Day") // []string

    // Now fetch the lyrics
    if len(suggestions) == 0 {
        // No lyrics found for this track :(
        // Try some other keywords or show some error to user
    }
    lyrics := golyrics.GetLyrics(suggestions[0]) // string


    // You can also search and fetch the lyrics with only one call
    // It'll use the first search result for fetching lyrics 
    lyrics := golyrics.SearchAndGetLyrics("Blackfield Some Day") // string
    // OR
    lyrics := golyrics.SearchAndGetLyricsByArtistAndName("Blackfield", "Some Day") // string    
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