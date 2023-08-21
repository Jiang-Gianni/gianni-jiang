package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteTailwindCSSContents = views.Section{
	Title: "Tailwind CSS",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://tailwindcss.com/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"Tailwind CSS is an open source CSS framework that creates utility CSS classes.",
				"It lets you add some classes to an HTML element rather than to write the element selector (tag / id / class) and its styling in a css file or inside a style tag: the HTML element gets more congested (especially with HTMX and Alpine) but I prefer writing 3-5 lines of attributes/values rather than to write 10-20 lines of CSS or JS (I am not fond of either).",
			},
		},
		{
			TextList: []string{
				"",
			},
		},
		{
			Subtitle: "Example",
			TextList: []string{
				"The following displayed element is styled by adding the classes:",
				"'animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200'",
			},
			RawHtml: views.TailwindDemo(),
		},
		{
			Subtitle: "CSS Size",
			TextList: []string{
				"It is possible to import the entire Tailwind library of utility classes by CDN but Tailwind offers a tool to optimize for production so that only the used classes are imported.",
			},
		},
		{
			Subtitle: "DaysiUI",
			RawHtml:  views.Link("https://daisyui.com/", "https://daisyui.com/", "_blank"),
			TextList: []string{
				"I am also using DaysiUI which is a library of prestyled components based on Tailwind CSS.",
				"For example, the buttons in the 'Demo Todo' section are styled by simply adding the 'btn' class.",
			},
		},
	},
}
