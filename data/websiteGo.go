package data

import views "github.com/Jiang-Gianni/gianni-jiang/views"

var WebsiteGoContents = views.Section{
	Title: "Go",
	Contents: []views.Content{
		{
			RawHtml: views.Link("https://go.dev/", "Link", "_blank"),
		},
		{
			TextList: []string{
				"Go is a statically compiled language designed at Google and it has become my favorite language.",
				"It is the tool that glues everything together in this project and I have used some of the community packages for this project",
			},
		},
		{
			Subtitle: "sqlc",
			RawHtml:  views.Link("https://github.com/sqlc-dev/sqlc", "https://github.com/sqlc-dev/sqlc", "_blank") + "<br>",
			TextList: []string{
				"This tool generates Go structs and query functions mapped using the input sql files that contain the 'create table' commands and the queries.",
				"What amazes me is that it can also generate custom structs in case the query consists of joining multiple tables: you end up having the structs and the query functions with type safety without writing any Go boilerplate code.",
			},
		},
		{
			Subtitle: "quicktemplate",
			RawHtml:  views.Link("https://github.com/valyala/quicktemplate", "https://github.com/valyala/quicktemplate", "_blank"),
			TextList: []string{
				"Go already has a template package in the standard library to generate text: you use some keyword inside double curly braces to control how the output should look like based on the given input.",
				"What I don't like is that some of the expression aren't Go-like: for example the 'if' equivalent in the template is '{{if condition}} Stuff {{end}}' and if you want to use a custom written function that returns a boolean as the condition, you need to put it in a map, pass the map to the Go template variable and call the map key in the if condition inside the curly braces.",
			},
		},
		{
			TextList: []string{
				"quicktemplate addresses this by letting you use the Go language inside the curly braces in the template files.",
				"In the template files you can define functions, structs and insert some raw Go code.",
				"The tool (executed from the command line) then generates Go files that contains all the templates definitions (functions, structs etc.) and can be exported and used in the project.",
				"These template functions and structs get compiled in the final binary: it is possible to run the final executable without any additional [html] file dependency in the directory.",
				"The drawback is that the intermediate parsing step is always needed even for small on-the-fly changes.",
			},
		},
	},
}
