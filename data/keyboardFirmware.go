package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var KeyboardFirmwareContents = views.Section{
	Title: "Firmware",
	Contents: []views.Content{
		{
			TextList: []string{
				"The last step is to install a firmware so that the MCU can send key presses when a switch is pushed.",
				"The most widely used firmware for custom keyboard is QMK (based on C)...",
			},
			RawHtml: views.Link("https://github.com/qmk/qmk_firmware", "https://github.com/qmk/qmk_firmware", "_blank"),
		},
		{
			TextList: []string{
				"... but I decided to go with PRK (based on Ruby).",
				"That's the reason I put an image of Poro (from LOL) on the MCU since PORO can also stand for 'Plain Old Ruby Object'",
			},
			RawHtml: views.Link("https://github.com/picoruby/prk_firmware/wiki/Tutorial", "https://github.com/picoruby/prk_firmware/wiki/Tutorial", "_blank"),
		},
		{
			TextList: []string{
				"The PRK firmware is very easy to install and all that is needed after that is to upload a 'keymap.rb' file that contains the definition of the pins, the matrix and the assigment of the keys.",
			},
		},
		{
			Subtitle: "Keymap",
			TextList: []string{
				"The layout I'm using is basically Colemak but with some modifications.",
				"I am missing the central 6 letter keys assigned to the index fingers which in Colemak are G, D, B, J, H and K (in Qwerty: T, G, B, Y, H, N).",
				"G, D, B and H have replaced comma, dot, slash and semicolons of the standard Colemak.",
				"J, K, numbers, symbols, F keys and navigations have been placed in secondary layers.",
				"In the thumb cluster I have Ctrl, Space, Shift, Backspace and two keys to activate the secondary layers.",
			},
		},
	},
}
