package data

import (
	views "github.com/Jiang-Gianni/gianni-jiang/views"
)

var WebsiteDemoTodoContents = views.Section{
	Title: "Demo Todo",
	Contents: []views.Content{
		{
			TextList: []string{
				"Todo apps have become the 'Hello World' of web development. Here is a demo that operates on the database.",
			},
		},
		{
			RawHtml: views.TodoDemo(),
		},
	},
}
