This probject is a simple slugifier for golang projects it has 2 main methods `Slugify` and `IDify` that both force ASCII text.

`Slugify` lowercases the output string and `IDify` does not. There is also a public `SanatizeText` method that takes a string and runs the transliterations on it which can still return unicode.

There is a public `OK` variable that is a sting containing additional characters that are allowed to exist in your slugified string. The `TO_DASH` variable is a list of additional characters to turn into a dash that may otherwise be allowed in a slugified string.

```
import "github.com/digitalxero/slugify"

var slugged = slugify.Slugify("日本語の手紙をテスト", 0)
var slugged6 = slugify.Slugify("日本語の手紙をテスト", 6)
```

You can also install the cli tool `go get -u github.com/digitalxero/slugify/cmd/slugify`
```
$ slugify --help
CLI Tool to slugify a string

Usage:
  slugify [flags]

Flags:
  -h, --help             help for slugify
      --lower
  -l, --max-len int
      --ok string        Non alphanumeric values that are OK to have in your output (default "-_")
      --to-dash string   Convert these to a dash instead of stripping them from the output
                         (default "/\\—–.~!@#$%^&*(){}[]+=?><;:`")

$ slugify --lower "日本語の手紙をテスト"
ri-ben-yu-noshou-zhi-wotesuto
$ slugify --lower "日本語の手紙をテスト" --max-len 6
ri-ben
```

