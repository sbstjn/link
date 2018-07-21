# sbstjn/link

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

  list := link.Get(`Bonjour, https://sbstjn.com is nice. https://example.com too!`)

  // Will show [https://sbstjn.com https://example.com]
  fmt.Println(list)
}
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.

