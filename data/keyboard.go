package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	KeyboardFirst  = "first"
	KeyboardSecond = "second"
)

var keyboardSort = []string{
	KeyboardFirst,
	KeyboardSecond,
}

var Keyboard = views.Project{
	Title:          "My Custom Keyboard",
	Api:            KeyboardApi,
	SortApis:       keyboardSort,
	DefaultSection: KeyboardFirst,
	SectionMap: map[string]views.Section{
		KeyboardFirst: {
			Title: "First Keyboard",
		},
		KeyboardSecond: {
			Title: "Second Keyboard",
		},
	},
}
