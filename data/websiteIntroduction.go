package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteIntroductionContents = views.Section{
	Title: "Introduction",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteIntroduction(),
		},
		// {
		// 	TextList: []string{
		// 		"This is a website that i made to showcase my projects.",
		// 		"You can jump directly to the section of the project you want to see by clicking the menu on the left.",
		// 	},
		// },
		// {
		// 	Subtitle: "CHATG stack",
		// 	RawHtml: views.Text(
		// 		[]string{
		// 			"To create this website I used the following tech stack:",
		// 		},
		// 	) + views.List(
		// 		[]string{
		// 			"CockroachDB",
		// 			"HTMX",
		// 			"AlpineJS",
		// 			"TailwindCSS",
		// 			"Go",
		// 		},
		// 	),
		// },
		// {
		// 	TextList: []string{
		// 		"I am still missing a P and a T to finally complete the CHATGPT stack.",
		// 		"If you have some ideas that I might like let me know.",
		// 	},
		// },
	},
}
