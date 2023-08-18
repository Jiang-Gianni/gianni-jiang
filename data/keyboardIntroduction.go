package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var KeyboardIntroductionContents = views.Section{
	Title: "Introduction",
	Contents: []views.Content{
		{
			TextList: []string{
				"If you are typing a couple (or more) hours a day, you should consider investing in a keyboard that would make typing more comfortable. It can be beneficial in the long run.",
				"That's why right after I started my first job I decided to buy a mechanical keyboard: REDRAGON Fizz K617.",
			},
		},
		{
			Subtitle: "REDRAGON Fizz K617",
			TextList: []string{
				"It is a 60% keyboard, meaning that it only has around 60% of the keys of a normal full keyboard: you need to press a combination of keys to access the missing keys.",
				"I changed the default switches into some silent light ones (Outemu silent white) so that I would feel less fatigue and I wouldn't make much noise from typing.",
			},
		},
		{
			Subtitle: "Colemak",
			TextList: []string{
				"Since I wanted to make typing even more ergonomic, I also decided to learn to type with a new layout: Colemak.",
				"It is one of the layouts optimized for touch typing in English: the most frequently used letters are moved to the home row to make them easier to press.",
				"Compared to the Qwerty layout, the positions of 17 of keys are changed but common shortcuts (like copy-pasta) are kept the same.",
			},
			RawHtml: views.Link("https://colemak.com/", "https://colemak.com/", "_blank"),
		},
	},
}
