package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	WebsiteIntroduction = "introduction"
	WebsiteSecond       = "cockroachdb"
)

var websiteSort = []string{
	WebsiteIntroduction,
	WebsiteSecond,
}

var Website = views.Project{
	Title:          "This Website",
	Api:            WebsiteApi,
	SortApis:       websiteSort,
	DefaultSection: WebsiteIntroduction,
	SectionMap: map[string]views.Section{
		WebsiteIntroduction: WebsiteIntroductionContents,
		WebsiteSecond:       WebsiteCockroachDBContents,
	},
}
