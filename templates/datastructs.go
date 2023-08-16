package views

type Content struct {
	Subtitle string
	TextList []string
	Img      string
	RawHtml  string
}

type Section struct {
	Title    string
	Contents []Content
}

type Project struct {
	Title          string
	Api            string
	SectionMap     map[string]Section
	DefaultSection string
	SortApis       []string
}
