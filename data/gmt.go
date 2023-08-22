package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	GmtIntroduction   = "introduction"
	GmtWhy            = "why"
	GmtParser         = "parser"
	GmtReason         = "reason"
	GmtImplementation = "implementation"
)

var gmtSort = []string{
	GmtIntroduction,
	GmtWhy,
	GmtParser,
	GmtReason,
	GmtImplementation,
	FeedbackApi,
}

var Gmt = views.Project{
	Title:          "Gmt",
	Api:            GmtApi,
	SortApis:       gmtSort,
	DefaultSection: WebsiteIntroduction,
	SectionMap: map[string]views.Section{
		GmtIntroduction:   GmtIntroductionContents,
		GmtWhy:            GmtWhyContents,
		GmtParser:         GmtParserContents,
		GmtReason:         GmtReasonContents,
		GmtImplementation: GmtImplementationContents,
		FeedbackApi:       FeedbackContents(GmtApi),
	},
}

var GmtIntroductionContents = views.Section{
	Title: "Introduction",
	Contents: []views.Content{
		{
			RawHtml: views.GmtIntroduction(),
		},
	},
}

var GmtWhyContents = views.Section{
	Title: "Why",
	Contents: []views.Content{
		{
			RawHtml: views.GmtWhy(),
		},
	},
}

var GmtParserContents = views.Section{
	Title: "Parser",
	Contents: []views.Content{
		{
			RawHtml: views.GmtParser(),
		},
	},
}

var GmtReasonContents = views.Section{
	Title: "Reason",
	Contents: []views.Content{
		{
			RawHtml: views.GmtReason(),
		},
	},
}

var GmtImplementationContents = views.Section{
	Title: "Implementation",
	Contents: []views.Content{
		{
			RawHtml: views.GmtImplementation(),
		},
	},
}
