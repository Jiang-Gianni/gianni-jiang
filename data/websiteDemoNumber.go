package data

import (
	views "github.com/Jiang-Gianni/gianni-jiang/views"
)

var WebsiteDemoNumberContents = views.Section{
	Title: "Demo Number",
	Contents: []views.Content{
		{
			TextList: []string{
				"Simple HTTP call that retrieves some trivia from 'numbersapi' of a random number.",
			},
			RawHtml: views.Link("http://numbersapi.com", "http://numbersapi.com", "_blank"),
		},
		{
			RawHtml: views.NumbersApi(),
		},
	},
}
