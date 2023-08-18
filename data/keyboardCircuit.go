package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var KeyboardCircuitContents = views.Section{
	Title: "Circuit",
	Contents: []views.Content{
		{
			TextList: []string{
				"After learning how these custom split keyboards works I decided to design and to build one myself and I am using it to type all of this.",
				"There are two types: hand wired and PCB (printed circuit board).",
				"I went for the latter as it looked easier.",
				"The tools I used to generate the PCB design are 'ergogen' and 'KiCad' and most of the design was inspired by the 'cheapino' keyboard.",
				"I also added a rotary encoder in the design but after the assembling it didn't work as I wanted it to.",
			},
			RawHtml: views.Link("https://github.com/ergogen/ergogen", "https://github.com/ergogen/ergogen", "_blank") + "<br>" + views.Link("https://www.kicad.org/", "https://www.kicad.org/", "_blank") + "<br>" + views.Link("https://github.com/tompi/cheapino", "https://github.com/tompi/cheapino", "_blank"),
		},
		{
			Subtitle: "Design",
			TextList: []string{
				"Here is my final circuit design.",
				"It can be used for both left and right splits and the microcontroller can be installed on either side.",
				"30 keys in total, 15 on each side, 3 keys for each finger.",
			},
			Img: "./assets/keyboard/kicad.png",
		},
		{
			Subtitle: "How it works",
			TextList: []string{
				"On each key there is a switch that, once pressed, connects the two side electrically completing the circuit.",
				"Each side of the switch is connected to a pin of the MCU (RP2040 Zero in my case) that elaborates what needs to happen when a specific key is pressed.",
			},
			RawHtml: views.Link("https://www.waveshare.com/wiki/RP2040-Zero", "https://www.waveshare.com/wiki/RP2040-Zero", "_blank"),
		},
		{
			TextList: []string{
				"Usually the microcontroller has a very limited amount of pins, less than the number of keyboard keys: in order to make it work a diode is connected to each key and a matrix arrangment of the connections is used (one pin of the MCU for each of the rows and columns of the keyboard).",
				"On each cycle the MCU scans the pins, checking if there are other pins with a signal. The scanning direction is set to 'row to column' or 'column to row'.",
				"Each switch is connected to a row pin and a column pin so that it can be identified when pressed.",
				"If (row1, col1) and (row2, col2) keys are pressed together it is possible for the MCU to also register (row1, col2) or (row2, col1): a diode on each key allows the current to flow only on one direction and prevents this effect called 'ghosting'.",
			},
		},
		{
			TextList: []string{
				"Given N number of free MCU pins, the maximum number of keyboard keys that can be registered using a regular matrix scanning method is (N/2)*(N/2) with (N/2) rows and (N/2) columns.",
				"This number can be doubled using a duplex matrix: the scanning goes both way and for each (row, col) pair there are 2 switches. The difference between these 2 switches is the orientation of the diodes (row2col or col2row).",
			},
			RawHtml: views.Link("https://kbd.news/The-Japanese-duplex-matrix-1391.html", "https://kbd.news/The-Japanese-duplex-matrix-1391.html", "_blank"),
		},
		{
			TextList: []string{
				"An RJ45 cable connects the two side: this allows for 8 connection pins to be used.",
			},
		},
		{
			Subtitle: "JLCPCB",
			TextList: []string{
				"After finishing the PCB design I needed to have it printed so I uploaded it to JLCPCB.",
				"I really like their service because they offer a special deal if the PCB is less than 10cm X 10cm in size: the total cost of 5 boards + shipping (to Italy) + taxes was only around 4€.",
			},
			RawHtml: views.Link("https://jlcpcb.com/", "https://jlcpcb.com/", "_blank"),
		},
		{
			Subtitle: "Aliexpress",
			TextList: []string{
				"The rest of the materials (soldering tools, diodes, switches, keycaps etc.) were bought from Aliexpress",
			},
		},
	},
}
