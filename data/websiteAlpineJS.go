package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteAlpineJSContents = views.Section{
	Title: "Alpine JS",
	Contents: []views.Content{
		{
			RawHtml: views.WebsiteAlpineJS(),
		},
		// {
		// 	RawHtml: views.Link("https://alpinejs.dev/", "Link", "_blank"),
		// },
		// {
		// 	TextList: []string{
		// 		"Alpine JS is a lightweight JavaScript framework: the compressed bundle is around 16 kB in size.",
		// 		"I decided to add this tool to the stack because in certain cases I felt like it was way easier to track the state of part of the application on the client rather than on the server.",
		// 	},
		// },
		// {
		// 	TextList: []string{
		// 		"Alpine JS lets you sprinkle a little bit of JavaScript where it is needed: if you inspect the menu element on the left (right click + 'Inspect') you'll see some attributes like 'x-data', 'x-show', 'x-on' and 'x-transition'.",
		// 		"After a click on the side menu, HTMX updates the content on the right and Alpine JS updates the local state of the menu, so that the current project and section are highlighted with different background color.",
		// 	},
		// },
		// {
		// 	TextList: []string{
		// 		"It is also possible to update the side menu through HTMX but it would require to add an event listener with the 'hx-on' attributes and to then make another HTTP request when the event happens.",
		// 	},
		// },
		// {
		// 	TextList: []string{
		// 		"Combining Alpine JS, HTMX and Tailwind CSS (next section) makes it possible to generate the following component just by adding attributes to HTML elements and without writing much JavaScript.",
		// 	},
		// },
		// {
		// 	RawHtml: views.AlpineCountToggle(),
		// },
	},
}
