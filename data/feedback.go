package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var FeedbackContents = func(project string) views.Section {
	return views.Section{
		Title: "Feedback",
		Contents: []views.Content{
			{
				TextList: []string{
					"Feel free to leave some feedbacks and suggestions so that I can make some improvements.",
				},
			},
			{
				RawHtml: views.Feedback(project),
			},
		},
	}
}
