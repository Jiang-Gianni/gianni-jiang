package data

import "github.com/Jiang-Gianni/gianni-jiang/views"

const (
	KeyboardIntroduction = "introduction"
	KeyboardSplit        = "split"
	KeyboardCircuit      = "circuit"
	KeyboardAssembling   = "assembling"
	KeyboardFirmware     = "firmware"
)

var keyboardSort = []string{
	KeyboardIntroduction,
	KeyboardSplit,
	KeyboardCircuit,
	KeyboardAssembling,
	KeyboardFirmware,
	FeedbackApi,
}

var Keyboard = views.Project{
	Title:          "Custom Keyboard",
	Api:            KeyboardApi,
	SortApis:       keyboardSort,
	DefaultSection: KeyboardIntroduction,
	SectionMap: map[string]views.Section{
		KeyboardIntroduction: KeyboardIntroductionContents,
		KeyboardSplit:        KeyboardSplitContents,
		KeyboardCircuit:      KeyboardCircuitContents,
		KeyboardAssembling:   KeyboardAssemblingContents,
		KeyboardFirmware:     KeyboardFirmwareContents,
		FeedbackApi:          FeedbackContents(KeyboardApi),
	},
}

var KeyboardIntroductionContents = views.Section{
	Title: "Introduction",
	Contents: []views.Content{
		{
			RawHtml: views.KeyboardIntroduction(),
		},
	},
}

var KeyboardSplitContents = views.Section{
	Title: "Split Keyboard",
	Contents: []views.Content{
		{
			RawHtml: views.KeyboardSplit(),
		},
	},
}

var KeyboardCircuitContents = views.Section{
	Title: "Circuit",
	Contents: []views.Content{
		{
			RawHtml: views.KeyboardCircuit(),
		},
	},
}

var KeyboardAssemblingContents = views.Section{
	Title: "Assembling",
	Contents: []views.Content{
		{
			RawHtml: views.KeyboardAssembling(),
		},
	},
}

var KeyboardFirmwareContents = views.Section{
	Title: "Firmware",
	Contents: []views.Content{
		{
			RawHtml: views.KeyboardFirmware(),
		},
	},
}
