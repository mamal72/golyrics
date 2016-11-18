package golyrics

import (
	"reflect"
	"testing"
)

func shouldNotPanic(test *testing.T) {
	if r := recover(); r != nil {
		test.Errorf("stripeHTMLTags() paniced")
	}
}

func Test_breakToNewLine(t *testing.T) {
	type args struct {
		HTML string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should work for breaks in the start of string",
			args: args{
				"<br/>Hello World",
			},
			want: "\nHello World",
		},
		{
			name: "test should work for breaks in the middle of string",
			args: args{
				"Hello<br/>World",
			},
			want: "Hello\nWorld",
		},
		{
			name: "test should work for breaks in the end of string",
			args: args{
				"Hello World<br/>",
			},
			want: "Hello World\n",
		},
	}
	for _, tt := range tests {
		if got := breakToNewLine(tt.args.HTML); got != tt.want {
			t.Errorf("%q. breakToNewLine() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_stripeHTMLTags(t *testing.T) {
	type args struct {
		HTML string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should work for different kinds of HTML tags",
			args: args{
				"<p>Hello World<br/>This is a TEST! :D</p>",
			},
			want: "Hello WorldThis is a TEST! :D",
		},
	}
	for _, tt := range tests {
		defer shouldNotPanic(t)
		if got := stripeHTMLTags(tt.args.HTML); got != tt.want {
			t.Errorf("%q. stripeHTMLTags() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getFormattedLyrics(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return a correct formatted string from HTML lyrics",
			args: args{
				"<div class='lyricbox'>The shortest song in the universe<br/>Really isn't much fun<br/>It only has one puny verse<br/>. . . and then it's done!<div class='lyricsbreak'></div>\n</div>",
			},
			want: "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
		},
	}
	for _, tt := range tests {
		if got := getFormattedLyrics(tt.args.text); got != tt.want {
			t.Errorf("%q. getFormattedLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getSearchURI(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return right URI",
			args: args{
				"blackfield:pain",
			},
			want: "http://lyrics.wikia.com/index.php?action=ajax&rs=getLinkSuggest&format=json&query=blackfield:pain",
		},
	}
	for _, tt := range tests {
		if got := getSearchURI(tt.args.query); got != tt.want {
			t.Errorf("%q. getSearchURI() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchLyrics(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test should return some results for some famous tracks",
			args: args{
				"metallica:unforgiven",
			},
			want: []string{
				"Metallica:Unforgiven",
				"Metallica:The Unforgiven II",
			},
		},
	}
	for _, tt := range tests {
		defer shouldNotPanic(t)
		if got := SearchLyrics(tt.args.query); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SearchLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchLyricsByArtistAndName(t *testing.T) {
	type args struct {
		artist string
		name   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test should return some results for some good real tracks",
			args: args{
				"blackfield",
				"end of the world",
			},
			want: []string{
				"Blackfield:End Of The World",
			},
		},
	}
	for _, tt := range tests {
		if got := SearchLyricsByArtistAndName(tt.args.artist, tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SearchLyricsByArtistAndName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetLyrics(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return right lyrics for a real track",
			args: args{
				"Sandra_Boynton:The_Shortest_Song_In_The_Universe",
			},
			want: "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
		},
	}
	for _, tt := range tests {
		defer shouldNotPanic(t)
		if got := GetLyrics(tt.args.title); got != tt.want {
			t.Errorf("%q. GetLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchAndGetLyrics(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return right lyrics for a real track",
			args: args{
				"Sandra_Boynton:The_Shortest_Song_In_The_Universe",
			},
			want: "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
		},
	}
	for _, tt := range tests {
		if got := SearchAndGetLyrics(tt.args.query); got != tt.want {
			t.Errorf("%q. SearchAndGetLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchAndGetLyricsByArtistAndName(t *testing.T) {
	type args struct {
		artist string
		name   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return right lyrics for a real track",
			args: args{
				"Sandra_Boynton",
				"The_Shortest_Song_In_The_Universe",
			},
			want: "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
		},
	}
	for _, tt := range tests {
		if got := SearchAndGetLyricsByArtistAndName(tt.args.artist, tt.args.name); got != tt.want {
			t.Errorf("%q. SearchAndGetLyricsByArtistAndName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
