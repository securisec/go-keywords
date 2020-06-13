package keywords

import (
	"io/ioutil"
	"testing"
)

func TestStringNoOptions(t *testing.T) {
	k, err := Extract(" How to solve a GeeTest 1 123 “slider CAPTCHA” with JS ")
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(k) != 5 {
		t.Errorf("Expected 5, got %d", len(k))
	}
}

func TestStringAllOptions(t *testing.T) {
	k, err := Extract(" How to solve solve a GeeTest 1 123 “slider CAPTCHA” with JS ", ExtractOptions{
		Lowercase:        true,
		AddStopwords:     []string{"captcha"},
		IgnorePattern:    "",
		Language:         LangEnglish,
		RemoveDigits:     true,
		RemoveDuplicates: true,
		StripTags:        false,
	})
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(k) != 4 {
		t.Errorf("Expected 5, got %d", len(k))
	}
}

func TestHtmlStripped(t *testing.T) {
	var err error
	data, err := ioutil.ReadFile("./test/test.html")
	k, err := Extract(string(data), ExtractOptions{StripTags: true})
	if err != nil {
		t.Error(err)
	}
	if len(k) != 6 {
		t.Errorf("Expected 6, got %d", len(k))
	}
}

func TestLowercase(t *testing.T) {
	k, err := Extract("This is a test string for HTML Go Javascript random keywords")
	if err != nil {
		t.Error(err)
	}
	if len(k) != 6 {
		t.Errorf("Got wrong value")
	}
}
