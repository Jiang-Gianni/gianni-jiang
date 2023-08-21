package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteHTMXContents = views.Section{
	Title: "HTMX",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://htmx.org/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"HTMX is a UI JavaScript library for the web that lets any HTML element make an HTTP request by simply adding some attributes.",
			},
		},
		{
			TextList: []string{
				"If you right click on the menu on the left and select 'Inspect' you can see that there are some attributes like 'hx-get' and 'hx-target'. A click on the element will trigger an HTTP request: the corresponding HTTP response contents will then substitute the element defined by the 'hx-target' attribute.",
			},
		},
		{
			TextList: []string{
				"This means that the HTTP response body can be any HTML content.",
			},
		},
		{
			Subtitle: "JavaScript size",
			TextList: []string{
				"There is no need on the client side to manipulate the response, no json to be parsed: this means that the page will load faster since there is no javascript (other than htmx) to be executed.",
				"If you press F12, go to 'Network', and press Ctrl + Shift + R to hard reload the page you'll see that the JavaScript size for htmx is around 15 kB.",
				"Now try to do the same on",
				"https://create-react-template.vercel.app",
				"You'll notice almost 50 kB of JavaScript. Imagine what happens as the app scales: with React (or other similar SPA frameworks) even more JavaScript will be shipped to the client.",
			},
		},
		{
			Subtitle: "Framework",
			TextList: []string{
				"Another advantage of HTMX is that you just need those 15 kB of JavaScript.",
				"No need for any 'npm install ...' which can easily cause the local 'node_modules' folder to exceed 100 MB of size.",
			},
			Img: "./assets/website/node-modules.png",
		},
		{
			Subtitle: "Server Side Rendering",
			TextList: []string{
				"With HTMX the web contents are rendered on the server which means that it is possible to use any backend language as long as the HTML contents are delivered to the client through the HTTP response. Many languages also offer some templating tools, so generating the needed HTML with the data obtained from the database or other services gets easier.",
			},
		},
	},
}
