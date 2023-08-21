package data

import (
	views "github.com/Jiang-Gianni/gianni-jiang/views"
)

var WebsiteDemoTodoContents = views.Section{
	Title: "Demo Todo",
	Contents: []views.Content{
		{
			TextList: []string{
				"Here is a demo of a Todo application that operates on the database.",
				"Click the 'Create New Todo' button to generate a new Todo or click one of the Todos in the table to update / delete it.",
			},
		},
		{
			RawHtml: views.TodoDemo(),
		},
	},
}
