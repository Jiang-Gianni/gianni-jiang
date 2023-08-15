package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	KeyboardIntroduction = "introduction"
	KeyboardSecond       = "second"
)

var keyboardSort = []string{
	KeyboardIntroduction,
	KeyboardSecond,
}

var Keyboard = views.Project{
	Title:          "My Custom Keyboard",
	Api:            KeyboardApi,
	SortApis:       keyboardSort,
	DefaultSection: KeyboardIntroduction,
	SectionMap: map[string]views.Section{
		KeyboardIntroduction: WebsiteIntroductionContents,
		KeyboardSecond:       WebsiteCockroachDBContents,
	},
}
