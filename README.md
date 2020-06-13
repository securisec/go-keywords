![](https://github.com/securisec/go-keywords/workflows/tests/badge.svg)

# go-keywords

`go-keywords` is an attempt to port over the npm package `keyword-extractor`. It is not as in depth as its npm package, but gets the work done. `go-keywords` works with both regular strings, and html data.

**Why?** While in search for a simple keyword extractor, i came across a few packages that uses the RAKE concept, but for my use case to simply extract individual keywords, it wasnt sufficient.

## Install
```go
go get github.com/securisec/go-keywords
```

## Use

### keywords.ExtractOptions
The `Extract` functions default behavior can be altered passing the `ExtractOptions` struct. The options and their default values are:

```go
type ExtractOptions struct {
	RemoveDigits     bool     // Remove digits from found words. Defaults to true
	RemoveDuplicates bool     // Remove duplications from found words. Defaults to true
	Lowercase        bool     // Change all found words to lowercase. Defaults to true
	Language         string   // Language to use for keywords extraction. Defaults to "en"
	AddStopwords     []string // Append additional stopwords to ignore. Defaults to empty array
	IgnorePattern    string   // Ignore matches for the specified regex pattern. Defaults to ""
	StripTags        bool     // Strip HTML tags. Defaults to false
}
```

### Simple use case

```go
package main

import (
	"fmt"

	"github.com/securisec/go-keywords"
)

func main() {
	data := " How to solve a GeeTest 1 123 “slider CAPTCHA” with JS "
	k, _ := keywords.Extract(data)
	fmt.Println(k)
}

// [solve geetest slider captcha js]
```

### Get keywords from html
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/securisec/go-keywords"
)

func main() {
	req, _ := http.Get("http://example.com")
	data, _ := ioutil.ReadAll(req.Body)
	k, _ := keywords.Extract(string(data), keywords.ExtractOptions{
		StripTags:        true,
		RemoveDuplicates: true,
		IgnorePattern:    "<.+>",
		Lowercase:        true,
	})
	fmt.Println(k)
}


// [illustrative examples documents information domain literature prior coordination permission]
```