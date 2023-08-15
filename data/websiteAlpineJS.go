package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteAlpineJSContents = views.Section{
	Title: "Alpine JS",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://alpinejs.dev/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"Alpine JS is a lightweight JavaScript framework",
			},
		},
	},
}
