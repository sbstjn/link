package link

import (
	"net/url"
	"regexp"
)

// Based on https://stackoverflow.com/a/8234912/713518
var expression = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[\S.-]+|(?:www.|[-;:&=\+\$,\w]+@)[\S.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[ˆ\S]*))?)`)

// Link reflects an URL link
type Link string

// List is a list of Links
type List []*Link

func (l Link) String() string {
	return string(l)
}

// GetAll returns a list of all links from a string
func GetAll(in string) List {
	var list List

	for _, result := range expression.FindAllString(in, -1) {
		if _, err := url.Parse(result); err == nil {
			link := Link(result)
			list = append(list, &link)
		}
	}

	return list
}

// Get returns the first link from a string
func Get(in string) *Link {
	if data := GetAll(in); len(data) > 0 {
		return data[0]
	}

	return nil
}
