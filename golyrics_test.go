package golyrics

import (
	"reflect"
	"testing"
)

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
		got := stripeHTMLTags(tt.args.HTML)
		if got != tt.want {
			t.Errorf("%q. stripeHTMLTags() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_fixApostrophes(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should work for apostrophes in the strings",
			args: args{
				"I&#39;m not strong enough to stay away\nCan&#39;t run from you...",
			},
			want: "I'm not strong enough to stay away\nCan't run from you...",
		},
	}
	for _, tt := range tests {
		if got := fixApostrophes(tt.args.text); got != tt.want {
			t.Errorf("%q. fixApostrophes() = %v, want %v", tt.name, got, tt.want)
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
		got := getFormattedLyrics(tt.args.text)
		if got != tt.want {
			t.Errorf("%q. getFormattedLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchTrack(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    []Track
		wantErr bool
	}{
		{
			name: "test should return some results for some famous tracks",
			args: args{
				"metallica:unforgiven",
			},
			want: []Track{{
				Artist: "Metallica",
				Name:   "Unforgiven",
			}, {
				Artist: "Metallica",
				Name:   "The Unforgiven II",
			}},
			wantErr: false,
		},
		{
			name: "test should return no results for wrong input query",
			args: args{
				"sadasdasfdsfkdsjfgkrjferjkgnf,gfdngirjdgfmv",
			},
			want:    []Track{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := SearchTrack(tt.args.query)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchLyrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SearchLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchTrackByArtistAndName(t *testing.T) {
	type args struct {
		artist string
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    []Track
		wantErr bool
	}{
		{
			name: "test should return some results for some good real tracks",
			args: args{
				"blackfield",
				"end of the world",
			},
			want: []Track{{
				Artist: "Blackfield",
				Name:   "End Of The World",
			}},
			wantErr: false,
		},
		{
			name: "test should return no results for unreal tracks",
			args: args{
				"sakfjweirufuxjn4hrfnmdxnjvdsfsdjhfhjsdfs",
				"askjdjkasdjksdjlfjsdkufslidfjlksdjklfjklsdfjklsdfs",
			},
			want:    []Track{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := SearchTrackByArtistAndName(tt.args.artist, tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchTrackByArtistAndName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SearchTrackByArtistAndName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTrack_FetchLyrics(t *testing.T) {
	type fields struct {
		Artist string
		Name   string
		Lyrics string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should fetch lyrics for the track and sets it on it",
			fields: fields{
				Artist: "Sandra_Boynton",
				Name:   "The_Shortest_Song_In_The_Universe",
				Lyrics: "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			track := &Track{
				Artist: tt.fields.Artist,
				Name:   tt.fields.Name,
				Lyrics: tt.fields.Lyrics,
			}
			if err := track.FetchLyrics(); (err != nil) != tt.wantErr {
				t.Errorf("Track.FetchLyrics() error = %v, wantErr %v", err, tt.wantErr)
			}
			if track.Lyrics != tt.fields.Lyrics {
				t.Errorf("%q. Track.FetchLyrics() = %v, want %v", tt.name, track.Lyrics, tt.fields.Lyrics)
			}
		})
	}
}
