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
