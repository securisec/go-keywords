package keywords

import (
	"fmt"
	"io/ioutil"
	"testing"

	keywords "github.com/securisec/go-keywords/languages"
)

func TestStringNoOptions(t *testing.T) {
	k, err := Extract(" How to solve a GeeTest 1 123 “slider CAPTCHA” with JS ")
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(k) != 5 {
		t.Errorf("Expected 5, got %d %s", len(k), k)
	}
}

func TestStringAllOptions(t *testing.T) {
	k, err := Extract(" How to solve solve a GeeTest 1 123 “slider CAPTCHA” with JS ", ExtractOptions{
		Lowercase:        true,
		AddStopwords:     []string{"captcha"},
		IgnorePattern:    "",
		Language:         keywords.English,
		RemoveDigits:     true,
		RemoveDuplicates: true,
		StripTags:        false,
	})
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(k)
	if len(k) != 4 {
		t.Errorf("Expected 5, got %d, %s", len(k), k)
	}
}

func TestHtmlStripped(t *testing.T) {
	var err error
	data, err := ioutil.ReadFile("./test/test.html")
	k, err := Extract(string(data), ExtractOptions{StripTags: true, IgnorePattern: "<.+>"})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 7 {
		t.Errorf("Expected 7, got %d %s", len(k), k)
	}
}

func TestLowercase(t *testing.T) {
	k, err := Extract("This is a test string for HTML Go Javascript random keywords")
	if err != nil {
		t.Error(err)
	}
	if len(k) != 6 {
		t.Errorf("Got wrong value. %d %s", len(k), k)
	}
}

func TestMatchOnly(t *testing.T) {
	k, err := Extract("This is a test string for HTML Go Javascript random keywords", ExtractOptions{
		MatchPattern: "\\b[a-z]{6}\\b",
		Lowercase:    true,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 2 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestSpanish(t *testing.T) {
	k, err := Extract("Lorem Ipsum es simplemente el texto de relleno de las imprentas y archivos de texto", ExtractOptions{
		Language: keywords.Spanish,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 8 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestFrench(t *testing.T) {
	k, err := Extract("Le Lorem Ipsum est simplement du faux texte employé dans la composition et la mise en page avant impression", ExtractOptions{
		Language:     keywords.French,
		MatchPattern: "\\b[a-z]{4}\\b",
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 3 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestDanish(t *testing.T) {
	k, err := Extract("Lorem Ipsum er ganske enkelt fyldtekst fra print- og typografiindustrien", ExtractOptions{
		Language: keywords.Danish,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 7 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestGerman(t *testing.T) {
	k, err := Extract("Lorem Ipsum ist ein einfacher Demo-Text für die Print- und Schriftindustrie", ExtractOptions{
		Language: keywords.German,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 6 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestItalian(t *testing.T) {
	k, err := Extract("Lorem Ipsum è un testo segnaposto utilizzato nel settore della tipografia e della stampa", ExtractOptions{
		Language: keywords.Italian,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 8 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestPortuguese(t *testing.T) {
	k, err := Extract("O Lorem Ipsum é um texto modelo da indústria tipográfica e de impressão", ExtractOptions{
		Language: keywords.Portuguese,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 7 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}

func TestRussian(t *testing.T) {
	k, err := Extract("Lorem Ipsum - это текст-\"рыба\", часто используемый в печати и вэб-дизайне", ExtractOptions{
		Language: keywords.Russian,
	})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 6 {
		t.Errorf("Got wrong value: %d, %s", len(k), k)
	}
}
