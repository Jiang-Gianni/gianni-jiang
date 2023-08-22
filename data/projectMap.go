package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	FeedbackApi = "feedback"
	WebsiteApi  = "website"
	KeyboardApi = "keyboard"
	GmtApi      = "gmt"
)

var SortApis = []string{
	WebsiteApi,
	KeyboardApi,
	GmtApi,
}

var ProjectMap = map[string]views.Project{
	WebsiteApi:  Website,
	KeyboardApi: Keyboard,
	GmtApi:      Gmt,
}
