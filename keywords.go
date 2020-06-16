package keywords

import (
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	keywords "github.com/securisec/go-keywords/languages"
)

// ExtractOptions Options to pass to the Extract function
type ExtractOptions struct {
	RemoveDigits     bool     // Remove digits from found words. Defaults to true
	RemoveDuplicates bool     // Remove duplications from found words. Defaults to true
	Lowercase        bool     // Change all found words to lowercase. Defaults to true
	Language         []string // Language to use for keywords extraction. Defaults to "English"
	AddStopwords     []string // Append additional stopwords to ignore. Defaults to empty array
	IgnorePattern    string   // Ignore matches for the specified regex pattern. Defaults to ""
	MatchPattern     string   // Only patch the specified regex pattern. Defaults to ""
	StripTags        bool     // Strip HTML tags. Defaults to false
}

// Extract Extract keywords from a string.
func Extract(s string, options ...ExtractOptions) ([]string, error) {
	var (
		stopwords    []string
		results      []string
		matchWords   []string
		words        []string
		strippedHTML string
		o            ExtractOptions
	)

	// Set default values for settings
	if len(options) > 0 {
		o = options[0]
		if len(o.AddStopwords) == 0 {
			o.AddStopwords = []string{}
		}
		// set default language
		if len(o.Language) == 0 {
			stopwords = keywords.English
		} else {
			stopwords = o.Language
		}
	} else {
		o = ExtractOptions{
			RemoveDigits:     true,
			RemoveDuplicates: true,
			Lowercase:        true,
			Language:         keywords.English,
			AddStopwords:     []string{},
			IgnorePattern:    "",
			MatchPattern:     "",
			StripTags:        false,
		}
		stopwords = o.Language
	}

	// Add additional stopwords is specified
	if len(o.AddStopwords) > 0 {
		stopwords = append(stopwords, o.AddStopwords...)
	}

	strippedHTML = s

	// strip html tags
	if o.StripTags {
		strippedHTML = strip.StripTags(s)
	}
	// If ignore pattern is set, sub those patterns
	if o.IgnorePattern != "" {
		strippedHTML = regexp.MustCompile(o.IgnorePattern).ReplaceAllString(strippedHTML, "")
	}
	splitRe := regexp.MustCompile("\\s")
	words = splitRe.Split(strings.TrimSpace(strippedHTML), -1)
	if len(words) == 0 {
		return catchErr(nil)
	}

	specialChars := "\\.|,|;|!|\\?|\\(|\\)|:|\"|\\^'|\\$|“|”|‘|’|”|<|>"

	for _, w := range words {
		checkLink := regexp.MustCompile("^https?://.+")
		if checkLink.MatchString(w) {
			matchWords = append(matchWords, w)
		} else {
			w = regexp.MustCompile(specialChars).ReplaceAllString(w, "")
			if len(w) == 1 {
				w = regexp.MustCompile("\\-|_|@|&|#").ReplaceAllString(w, "")
			}
			if o.RemoveDigits {
				w = regexp.MustCompile("\\d").ReplaceAllString(w, "")
			}
			if len(w) > 0 {
				matchWords = append(matchWords, w)
			}
		}
	}

	for _, word := range matchWords {
		if matcher(strings.ToLower(word), stopwords) {
			if o.Lowercase {
				results = append(results, strings.ToLower(word))
			} else {
				results = append(results, word)
			}
		}
	}

	// Match only requested pattern
	if o.MatchPattern != "" {
		initialResult := results
		results = []string{}
		for _, r := range initialResult {
			if regexp.MustCompile(o.MatchPattern).MatchString(r) {
				results = append(results, r)
			}
		}
	}

	// remove duplicates and return
	if o.RemoveDuplicates {
		return unique(results), nil
	}
	return results, nil
}

func catchErr(err error) ([]string, error) {
	return []string{}, err
}

func matcher(s string, stopwords []string) bool {
	for _, w := range stopwords {
		if w == s {
			return false
		}
	}
	return true
}

func unique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
