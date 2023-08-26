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

var WebsiteIntroductionContents = views.Section{
	Title: "Introduction",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteIntroduction(),
		},
	},
}

var WebsiteCockroachDBContents = views.Section{
	Title: "CockroachDB",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteCockroachDB(),
		},
	},
}

var WebsiteHTMXContents = views.Section{
	Title: "HTMX",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteHTMX(),
		},
	},
}

var WebsiteAlpineJSContents = views.Section{
	Title: "Alpine JS",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteAlpineJS(),
		},
	},
}

var WebsiteTailwindCSSContents = views.Section{
	Title: "Tailwind CSS",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteTailwindCSS(),
		},
	},
}

var WebsiteGoContents = views.Section{
	Title: "Go",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteGo(),
		},
	},
}

var WebsiteDemoTodoContents = views.Section{
	Title: "Demo Todo",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteDemoTodo(),
		},
	},
}

var WebsiteDemoNumberContents = views.Section{
	Title: "Demo Number",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteDemoNumber(),
		},
	},
}

var WebsiteDeployContents = views.Section{
	Title: "Deploy",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteDeploy(),
		},
	},
}
