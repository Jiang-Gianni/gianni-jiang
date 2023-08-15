package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteHTMXContents = views.Section{
	Title: "HTMX",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://htmx.org/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"HTMX is a UI JavaScript library for the web that lets any HTML element make an HTTP request by simply adding some attributes.",
			},
		},
		{
			TextList: []string{
				"If you right click on the menu on the left and select 'Inspect' you can see that there are some attributes like 'hx-get' and 'hx-target'",
			},
		},
	},
}
