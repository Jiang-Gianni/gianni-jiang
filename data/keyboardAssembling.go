package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var KeyboardAssemblingContents = views.Section{
	Title: "Assembling",
	Contents: []views.Content{
		{
			TextList: []string{
				"Here some pictures of the assembling process.",
			},
		},
		{
			Img: "./assets/keyboard/a1.jpg",
		},
		{
			Img: "./assets/keyboard/a2.jpg",
		},
		{
			Img: "./assets/keyboard/a3.jpg",
		},
		{
			Img: "./assets/keyboard/a4.jpg",
		},
		{
			Img: "./assets/keyboard/a5.jpg",
		},
		{
			Img: "./assets/keyboard/a6.jpg",
		},
		{
			Img: "./assets/keyboard/a7.jpg",
		},
		{
			TextList: []string{
				"In the last picture if you are wondering what the extra small little cable is for, well... let's say that there was a 'Happy Little Accident' during the design of the PCB.",
				"An RJ45 PCB connection has 8 pins and they are positioned in 2 staggered rows, sort of like a parallelogram.",
				"Since it is not symmetric, I added an extra pin hole so that the board could be flipped and be used both as left and right split but I connected the pinholes to the MCU as if there were 9 available connecting pins.",
			},
		},
	},
}
