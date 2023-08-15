package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	WebsiteIntroduction = "introduction"
	WebsiteCockroachDB  = "cockroachdb"
	WebsiteHTMX         = "htmx"
	WebsiteAlpineJS     = "alpinejs"
)

var websiteSort = []string{
	WebsiteIntroduction,
	WebsiteCockroachDB,
	WebsiteHTMX,
	WebsiteAlpineJS,
}

var Website = views.Project{
	Title:          "Website",
	Api:            WebsiteApi,
	SortApis:       websiteSort,
	DefaultSection: WebsiteIntroduction,
	SectionMap: map[string]views.Section{
		WebsiteIntroduction: WebsiteIntroductionContents,
		WebsiteCockroachDB:  WebsiteCockroachDBContents,
		WebsiteHTMX:         WebsiteHTMXContents,
		WebsiteAlpineJS:     WebsiteAlpineJSContents,
	},
}
