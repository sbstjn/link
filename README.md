# sbstjn/link

Get links from a string.

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