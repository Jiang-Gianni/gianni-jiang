package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteCockroachDBContents = views.Section{
	Title: "CockroachDB",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://www.cockroachlabs.com/", "https://www.cockroachlabs.com/"),
		},
		{
			TextList: []string{
				"CockroachDB is a commercial distributed PostgreSQL database management system and it is written in Go.",
			},
		},
		{
			TextList: []string{
				"Great serverless free base plan (10 GB and 50M request units per month) and the database cluster can be created in different regions in the world (only one in Europe at the moment of writing).",
			},
		},
		{
			TextList: []string{
				"I am using the database for some CRUD operations of a demo Todo application and some feedbacks.",
				"I am not currently using it to store all this text contents, since I feel it would be rather bothersome to update the database for each content change.",
			},
		},
	},
}
