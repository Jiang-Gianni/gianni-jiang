package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var KeyboardSplitContents = views.Section{
	Title: "Split Keyboard",
	Contents: []views.Content{
		{
			TextList: []string{
				"After a year and a half of using the 60% keyboard, I looked into some custom split ergonomical keyboards.",
				"There are some issues (kind of) that split keyboard can improve.",
			},
		},
		{
			Subtitle: "Arms position",
			TextList: []string{
				"To type on a normal keyboard you have to bend your arms toward the center of your desk where the keyboard is located.",
				"With a split keyboard, your hands stay at shoulder distance (or at any distance of your preference) with less strain on your shoulder and arms.",
			},
		},
		{
			Subtitle: "Keys position",
			TextList: []string{
				"In a normal keyboard, the keys are staggered and in order to touch type you should assign a set of keys to each of your fingers.",
			},
			Img: "./assets/keyboard/layout.png",
		},
		{
			TextList: []string{
				"Try to type the word 'byzantine'. Did you find it comfortable to type the first three letters?",
				"Because of the staggered layout, B and Y are the furthest keys assigned to index fingers (relative to F and J which are the base positions). I often misstype their Colemak counterparts because of it.",
				"According to many touch typing guides, letter Z should be pressed with the left pinkie.",
				"Both hands are already bent toward the center, why strain the outer and weakest finger even more to type Z? I find it more comfortable to type it with my left ring finger.",
				"Split keyboards can be ortholinear, solving these problems.",
			},
		},
		{
			Subtitle: "Less keys and less finger movements",
			TextList: []string{
				"Some split keyboard are 40%. The idea is that there are multiple layers of keys that can be activated by holding some specific keys.",
				"This means that you need to press multiple keys together to type, for example, symbols or to use the F keys (F1, F2, F3 etc).",
				"The advantage is that the fingers don't move that much: it is easier (and even faster once you memorize the combo) to press multiple keys in easy to reach positions than to press a single key that requires to lift and move the entire hand.",
			},
		},
	},
}
