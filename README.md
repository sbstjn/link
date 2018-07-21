# sbstjn/link

[![Current Release](https://badgen.now.sh/github/release/sbstjn/link)](https://github.com/sbstjn/link/releases)
[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](https://github.com/sbstjn/link/blob/master/LICENSE.md)
[![CircleCI Build Status](https://badgen.now.sh/circleci/github/sbstjn/link)](https://circleci.com/gh/sbstjn/link)

Go package to get links from a string.

## Install

```bash
$ > go get github.com/sbstjn/link
```

## Usage

```go
import "github.com/sbstjn/link"

func main() {
  item := link.Get(`Bonjour, https://sbstjn.com is nice.`)

  // Will show https://sbstjn.com
  fmt.Println(item.String())

  list := link.GetAll(`Bonjour, https://sbstjn.com is nice, but https://example.com too!`)

  // Will show [https://sbstjn.com https://example.com]
  fmt.Println(list)
}
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
