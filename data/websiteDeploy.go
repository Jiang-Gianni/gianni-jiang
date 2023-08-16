package data

import (
	views "github.com/Jiang-Gianni/gianni-jiang/views"
)

var WebsiteDeployContents = views.Section{
	Title: "Deploy",
	Contents: []views.Content{
		{
			TextList: []string{
				"In order to make this app public, I needed to deploy it somewhere.",
				"After researching and trying various services I think that the best platforms with free base plans and Go language support are Vercel and Render.",
			},
		},
		{
			Subtitle: "Vercel",
			TextList: []string{
				"Vercel is commonly used to deploy JavaScript based Web application, but it offers a Go Runtime with Serverless Functions, so all I had to do was creating a 'vercel.json' file in my project to redirect the HTTP request, and a go file in the 'api' directory that contains the HTTP handler functions like shown in their documentation.",
				"The downside is that Serverless Function have a maximum execution duration, so it is not possible to maintain a WebSocket connection.",
			},
			RawHtml: views.Link("https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go", "https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go", "_blank"),
		},
		{
			Subtitle: "Render",
			TextList: []string{
				"Render has support for Docker and for many different programming languages. You can define a build step of the application and the start command.",
				"With Render it is possible to use WebSocket.",
				"The downside is that if the server doesn't receive any request for a period of time then the execution stops and the cold start may take time (even more than 10 seconds)",
			},
			RawHtml: views.Link("https://render.com/docs", "https://render.com/docs", "_blank"),
		},
	},
}
