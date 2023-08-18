package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteCockroachDBContents = views.Section{
	Title: "CockroachDB",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://www.cockroachlabs.com/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"CockroachDB is a commercial distributed PostgreSQL database management system and it is written in Go.",
			},
		},
		{
			TextList: []string{
				"I am using the database for some CRUD operations of a demo Todo application and some feedbacks.",
				"I am not using it to store all this text contents, since I feel it would be rather bothersome to update a table rows for each content change.",
			},
		},
		{
			TextList: []string{
				"With Cockroach DB the database cluster can be created in different regions in the world and they offer a great serverless free base plan: 10 GB and 50M request units per month.",
			},
		},
		{
			TextList: []string{
				"(Un)fortunately this website won't get many views and I won't ever reach that cap, so I don't have to worry about it.",
			},
		},
		{
			RawHtml: views.WebsiteHarold(),
		},
	},
}
