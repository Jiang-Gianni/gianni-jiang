package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	KeyboardIntroduction = "introduction"
	KeyboardSixty        = "sixty"
)

var keyboardSort = []string{
	KeyboardIntroduction,
	KeyboardSixty,
}

var Keyboard = views.Project{
	Title:          "Custom Keyboard",
	Api:            KeyboardApi,
	SortApis:       keyboardSort,
	DefaultSection: KeyboardIntroduction,
	SectionMap: map[string]views.Section{
		KeyboardIntroduction: KeyboardIntroductionContents,
		KeyboardSixty:        KeyboardSixtyContents,
	},
}
