package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	WebsiteIntroduction = "introduction"
	WebsiteCockroachDB  = "cockroachdb"
	WebsiteHTMX         = "htmx"
	WebsiteAlpineJS     = "alpinejs"
	WebsiteTailwindCSS  = "tailwindcss"
	WebsiteGo           = "go"
	WebsiteTodo         = "demo"
	WebsiteNumber       = "number"
	WebsiteDeploy       = "deploy"
)

var websiteSort = []string{
	WebsiteIntroduction,
	WebsiteCockroachDB,
	WebsiteHTMX,
	WebsiteAlpineJS,
	WebsiteTailwindCSS,
	WebsiteGo,
	WebsiteTodo,
	WebsiteNumber,
	WebsiteDeploy,
	FeedbackApi,
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
		WebsiteTailwindCSS:  WebsiteTailwindCSSContents,
		WebsiteGo:           WebsiteGoContents,
		WebsiteTodo:         WebsiteDemoTodoContents,
		WebsiteNumber:       WebsiteDemoNumberContents,
		WebsiteDeploy:       WebsiteDeployContents,
		FeedbackApi:         FeedbackContents(WebsiteApi),
	},
}
