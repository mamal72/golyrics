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
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test should work for different kinds of HTML tags",
			args: args{
				"<p>Hello World<br/>This is a TEST! :D</p>",
			},
			want:    "Hello WorldThis is a TEST! :D",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := stripeHTMLTags(tt.args.HTML)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. stripeHTMLTags() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
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
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test should return a correct formatted string from HTML lyrics",
			args: args{
				"<div class='lyricbox'>The shortest song in the universe<br/>Really isn't much fun<br/>It only has one puny verse<br/>. . . and then it's done!<div class='lyricsbreak'></div>\n</div>",
			},
			want:    "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := getFormattedLyrics(tt.args.text)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getFormattedLyrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. getFormattedLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchLyrics(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
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
			wantErr: false,
		},
		{
			name: "test should return no results for wrong input query",
			args: args{
				"sadasdasfdsfkdsjfgkrjferjkgnf,gfdngirjdgfmv",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := SearchLyrics(tt.args.query)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchLyrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    []string
		wantErr bool
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
			wantErr: false,
		},
		{
			name: "test should return no results for unreal tracks",
			args: args{
				"sakfjweirufuxjn4hrfnmdxnjvdsfsdjhfhjsdfs",
				"askjdjkasdjksdjlfjsdkufslidfjlksdjklfjklsdfjklsdfs",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := SearchLyricsByArtistAndName(tt.args.artist, tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchLyricsByArtistAndName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SearchLyricsByArtistAndName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetLyrics(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test should return right lyrics for a real track",
			args: args{
				"Sandra_Boynton:The_Shortest_Song_In_The_Universe",
			},
			want:    "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := GetLyrics(tt.args.title)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetLyrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. GetLyrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearchAndGetLyrics(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test should search and return right lyrics for a real track",
			args: args{
				"Sandra_Boynton:The_Shortest_Song_In_The_Universe",
			},
			want:    "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
			wantErr: false,
		},
		{
			name: "test should right an error for an unreal track",
			args: args{
				"kjsjdkajskdjkasjkldkljadklasuid397r8fjdsfjsdkffsdfksdjfks",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := SearchAndGetLyrics(tt.args.query)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchAndGetLyrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
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
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test should return right lyrics for a real track by artist and name",
			args: args{
				"Sandra_Boynton",
				"The_Shortest_Song_In_The_Universe",
			},
			want:    "The shortest song in the universe\nReally isn't much fun\nIt only has one puny verse\n. . . and then it's done!\n",
			wantErr: false,
		},
		{
			name: "test should return an error for an unreal track by artist and name",
			args: args{
				"sf.kjjkrjk4jk53k4mm34n,5j3",
				"m43m5krdk.fgjkdjguijkgfkdgdf",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := SearchAndGetLyricsByArtistAndName(tt.args.artist, tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SearchAndGetLyricsByArtistAndName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SearchAndGetLyricsByArtistAndName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
