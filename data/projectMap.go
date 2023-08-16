package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	WebsiteApi  = "website"
	KeyboardApi = "keyboard"
)

var SortApis = []string{
	WebsiteApi,
	KeyboardApi,
}

var ProjectSort = []string{
	KeyboardIntroduction,
	KeyboardSixty,
}

var ProjectMap = map[string]views.Project{
	WebsiteApi:  Website,
	KeyboardApi: Keyboard,
}
