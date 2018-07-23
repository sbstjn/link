package link_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/sbstjn/link"
)

var _ = Describe("Link", func() {
	DescribeTable("Single link in Text",
		func(in string, link string) {
			found := Get(in)

			Expect(found).NotTo(BeNil())
			Expect(found.String()).To(Equal(link))
		},

		Entry("spec #1", `scheme://userinfo:pass@host/path?query#fragment`, "scheme://userinfo:pass@host/path?query#fragment"),
		Entry("spec #2", `scheme://userinfo@host/path?query#fragment`, "scheme://userinfo@host/path?query#fragment"),

		Entry("basic", `This is a http://example.com link`, "http://example.com"),
		Entry("basic w/ slash", `This is a http://example.com/ link`, "http://example.com/"),

		Entry("subdomain #1", `This is a http://a.example.com link`, "http://a.example.com"),
		Entry("subdomain #2", `This is a http://a.b.c.example.com link`, "http://a.b.c.example.com"),
		Entry("subdomain #4", `This is a http://a.b.c.d-e-f-very-long-subdomain.example.com link`, "http://a.b.c.d-e-f-very-long-subdomain.example.com"),

		Entry("folder w/ slash", `This is a http://example.com/sub/ link`, "http://example.com/sub/"),
		Entry("folder w/o slash", `This is a http://example.com/sub link`, "http://example.com/sub"),
		Entry("file", `This is a http://example.com/file.html link`, "http://example.com/file.html"),
		Entry("anchor on file", `This is a http://example.com/file.html#anchor link`, "http://example.com/file.html#anchor"),
		Entry("anchor on file w/ route", `This is a http://example.com/file.html#!anchor/sub link`, "http://example.com/file.html#!anchor/sub"),
		Entry("anchor on folder", `This is a http://example.com/file#anchor link`, "http://example.com/file#anchor"),
		Entry("anchor on folder w/ slash", `This is a http://example.com/file/#anchor link`, "http://example.com/file/#anchor"),

		Entry("anchor on query #1", `This is a http://example.com/file/index.php?foo=bar&baz=lorem#anchor link`, "http://example.com/file/index.php?foo=bar&baz=lorem#anchor"),
		Entry("anchor on query #2", `This is a http://example.com/file?foo=bar&baz=lorem#anchor link`, "http://example.com/file?foo=bar&baz=lorem#anchor"),
		Entry("anchor on query #3", `This is a http://example.com/?foo=bar&baz=lorem#anchor link`, "http://example.com/?foo=bar&baz=lorem#anchor"),

		Entry("weird #1", `foo http://موقع.وزارة-الاتصات.مصر bar`, `http://موقع.وزارة-الاتصات.مصر`),
		Entry("weird #2", `foo http://موقع.وزارة-الاتصات.مصر/asd.html`, `http://موقع.وزارة-الاتصات.مصر/asd.html`),
		Entry("weird #2", `foo https://müller.de/foo`, `https://müller.de/foo`),
		Entry("weird #3", `foo http://xn--4gbrim.xn----2c.xn--123/a/b/v.html#123`, `http://xn--4gbrim.xn----2c.xn--123/a/b/v.html#123`),
	)

	DescribeTable("Single link in multiline text",
		func(in string, link string) {
			found := Get(in)

			Expect(found).NotTo(BeNil())
			Expect(found.String()).To(Equal(link))
		},

		Entry("middle #1", `This 
scheme://userinfo:pass@host/path?query#fragment is
intended`, "scheme://userinfo:pass@host/path?query#fragment"),
		Entry("middle #2", `
scheme://userinfo@host/path?query#fragment`, "scheme://userinfo@host/path?query#fragment"),
		Entry("start", `scheme://userinfo@host/path?query#fragment
foo`, "scheme://userinfo@host/path?query#fragment"),
		Entry("end", `foo
scheme://userinfo@host/path?query#fragment`, "scheme://userinfo@host/path?query#fragment"),
	)

	DescribeTable("First link in Text",
		func(in string, link string) {
			found := Get(in)

			Expect(found).NotTo(BeNil())
			Expect(found.String()).To(Equal(link))
		},

		Entry("spec #1", `scheme://userinfo:pass@host/path?query#fragment scheme://userinfo:pass@host/path?query#fragment scheme3://userinfo:pass@host/path?query#fragment`, "scheme://userinfo:pass@host/path?query#fragment"),
		Entry("spec #2", `Text scheme://userinfo:pass@host/path?query#fragment scheme://userinfo:pass@host/path?query#fragment scheme3://userinfo:pass@host/path?query#fragment`, "scheme://userinfo:pass@host/path?query#fragment"),
		Entry("spec #3", `scheme://userinfo:pass@host/path?query#fragment Text scheme://userinfo:pass@host/path?query#fragment scheme3://userinfo:pass@host/path?query#fragment`, "scheme://userinfo:pass@host/path?query#fragment"),
		Entry("spec #4", `scheme://userinfo:pass@host/path?query#fragment Text scheme://userinfo:pass@host/path?query#fragment scheme3://userinfo:pass@host/path?query#fragment Text`, "scheme://userinfo:pass@host/path?query#fragment"),
	)

	DescribeTable("All links in Text",
		func(in string, links []string) {
			found := GetAll(in)

			Expect(found).NotTo(BeNil())

			for index, link := range found {
				Expect(link.String()).To(Equal(links[index]))
			}
		},

		Entry("basic", `http://foo.bar https://foo.bar`, []string{"http://foo.bar", "https://foo.bar"}),
		Entry("basic", `https://foo.bar http://foo.bar`, []string{"https://foo.bar", "http://foo.bar"}),
		Entry("basic", `With https://foo.bar Text http://foo.bar anywhere.`, []string{"https://foo.bar", "http://foo.bar"}),
		Entry("basic", `With
		 https://foo.bar Text and line
		 http://foo.bar breaks anywhere.`, []string{"https://foo.bar", "http://foo.bar"}),
	)
})
