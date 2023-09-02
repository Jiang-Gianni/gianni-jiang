# Introduction
<!-- `{% func WebsiteIntroduction() %}` -->

This is a website that i made to showcase my projects.

You can jump directly to the section of the project you want to see by clicking the <!-- class="bg-teal-300 rounded-s" -->**menu on the left**

<br>

## CHATG Stack

To create this website I used the following tech stack:
<!-- class="list-disc list-inside" -->
* *CockroachDB*<!-- class="bg-orange-200" -->
* *HTMX*<!-- class="bg-green-200" -->
* *AlpineJS*<!-- class="bg-green-200" -->
* *TailwindCSS*<!-- class="bg-green-200" -->
* *Go*<!-- class="bg-blue-200" -->

<br>



<!-- class="mx-auto h-full" -->
```pikchr
D: cylinder height 3cm width 3cm fill Bisque rad 10px "CockroachDB" "(Database)"
arrow <-> right 200%
S: box rad 10px height 3cm width 3cm fill LightSkyBlue "Go" "(Server)"
arrow <-> right 200%
B: box height 3cm width 3cm fill LightGreen "HTMX" "Alpine JS" "Tailwind Css" "(Client)"
```

<br>

I am still missing a P and a T to finally complete the **CHATGPT** stack.

If you have some ideas that I might like let me know.
<!-- `{% endfunc %}` -->


# CockroachDB
<!-- `{% func WebsiteCockroachDB() %}` -->

**[https://www.cockroachlabs.com/](https://www.cockroachlabs.com/)**<!-- target="_blank" -->

<br>

CockroachDB is a commercial distributed PostgreSQL database management system and it is written in Go.

<br>

I am using the database for some CRUD operations of a **Demo Todo application**<!-- hx-get="website/demo" hx-target="#contents" class="btn btn-accent" --> and
some **feedbacks**<!-- hx-get="website/feedback" hx-target="#contents" class="btn btn-accent" -->.

<br>

I am not using it to store all this text contents, since I feel it would be rather bothersome to update a table rows for each content change.

<br>

With Cockroach DB the database cluster can be created in different regions in the world and they offer a great serverless free base plan: **10 GB and 50M request units per month**<!-- class="bg-green-200" -->.

<br>

(Un)fortunately this ~~destined to fail~~ website won't get many views and I won't ever reach that cap, so I don't have to worry about it.

<br>

<div x-data="{show:false}"><div>

**Me Right Now**<!-- class="btn btn-primary" x-show="!show" @click="show=true" -->
**![Me Right Now](./assets/website/harold.jpg)**<!-- class="mx-auto h-full" x-show="show" -->

<!-- `{% endfunc %}` -->


# HTMX
<!-- `{% func WebsiteHTMX() %}` -->
**[https://htmx.org/](https://htmx.org/)**<!-- target="_blank" -->

<br>

HTMX is a UI JavaScript library for the web that lets any HTML element make an HTTP request by simply adding some attributes.

<br>

If you right click on the menu on the left and select 'Inspect' you can see that there are some attributes like **`hx-get`** and **`hx-target`**. A click on the element will trigger a GET HTTP request to the endpoint value: **the corresponding HTTP response contents will then substitute the element defined by the `hx-target` attribute**<!-- class="bg-yellow-200" -->.

<br>

This means that the HTTP response body must be in form of HTML.

<br>

**Any HTML element can trigger an HTTP request and all the HTTP verbs can be used (POST, GET, PUT, PATCH, and DELETE)**<!-- class="bg-orange-200" --> in contrast to plain HTML where only POST (with 'form' tags) and GET (with 'a' tags) are possible.

<br>

## JavaScript size

There is no need on the client side to manipulate the response, **no json to be parsed**: this means that the page will load faster since **there is no javascript (other than htmx) to be executed**<!-- class="bg-cyan-200" -->.

If you press F12, go to 'Network', and press Ctrl + Shift + R to hard reload the page you'll see that **the JavaScript size for htmx is around 15-16 kB**<!-- class="bg-green-200" -->.

<br>

Now try to do the same on **[https://create-react-template.vercel.app](https://create-react-template.vercel.app)**<!-- target="_blank" class="animate-pulse bg-gray-200" -->

<br>

You'll notice **almost 50 kB of JavaScript**<!-- class="bg-red-200" -->.

Imagine what happens as the app scales: with React (or other similar SPA frameworks) even more JavaScript will be shipped to the client, while with HTMX it will stay the same.

<br>

## Framework

Another advantage of HTMX is that you just need those 15 kB of JavaScript.

No need for any **`npm install ...`** which can easily cause the local **`node_modules`** folder to exceed 100 MB of size when using frameworks like React.

**![Node Modules](./assets/website/node-modules.png)**<!-- class="mx-auto h-full" -->

<br>


## Server Side Rendering

With HTMX the web contents are rendered on the server which means that **it is possible to use any backend language**: all the HTML contents are delivered to the client through the HTTP response.

<br>

Many languages also offer some templating tools, so generating the needed HTML with the data obtained from the database or other services gets easier (more on this later in the Go section).

<br>

In my opinion it also requires **less effort to parse the server fetched data into ready-to-use HTML directly from the backend rather than exposing a json data API that needs to be consumed and parsed from the frontend**<!-- class="bg-purple-200" -->.

Same goes for managing **the state of the application**<!-- class="bg-indigo-200" -->, most of the times.


<br>

<!-- `{% endfunc %}` -->

# Alpine JS
<!-- `{% func WebsiteAlpineJS() %}` -->
**[https://alpinejs.dev/](https://alpinejs.dev/)**<!-- target="_blank" -->
<br>

Alpine JS is a lightweight JavaScript framework: the compressed bundle is **around 16 kB in size**<!-- class="bg-green-200" -->.

<br>

Alpine JS lets you sprinkle a little bit of JavaScript where it is needed: if you inspect the menu element on the left (right click + 'Inspect') you'll see some attributes like **'x-data', 'x-show', 'x-on' and 'x-transition'**<!-- class="bg-yellow-200" -->.

<br>

After a click on the side menu, HTMX updates the content on the right and Alpine JS updates the local state of the menu, so that the current project and section are highlighted with different background color.

<br>

It is also possible to update the side menu through HTMX but I felt like in cases like this it was way easier to use a bit of JavaScript, especially with Alpine JS short syntax.

<br>

Combining Alpine JS, HTMX and Tailwind CSS (next section) makes it possible to generate the following component **just by adding attributes to HTML elements**<!-- class="bg-indigo-200" -->.

<!-- `{%= AlpineCountToggle() %}` -->

<!-- `{% endfunc %}` -->

# TailwindCSS
<!-- `{% func WebsiteTailwindCSS() %}` -->
**[https://tailwindcss.com/](https://tailwindcss.com/)**<!-- target="_blank" -->
<br>

Tailwind CSS is an open source CSS framework that creates utility CSS classes.

<br>

It lets you add some classes to an HTML element rather than to write the element selector (tag / id / class) and its styling in a css file or inside a style tag: the HTML element gets more congested (especially with HTMX and Alpine) but I prefer writing 3-5 lines of attributes and values rather than to write 10-20 lines (or more) of CSS or JS (I am not fond of either).

## Example


The following displayed element is styled by adding the classes:

**`animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200`**

<br>

<div class="animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200">Hello World</div>

## CSS Size

It is possible to import the entire Tailwind library of utility classes by CDN but Tailwind offers a tool to optimize for production so that only the used classes are imported.

## DaisyUI

I am also using **[DaisyUI](https://daisyui.com/)**<!-- class="btn btn-accent" target="_blank" --> which is a library of prestyled components based on Tailwind CSS.

For example, all the buttons in this website are styled by simply adding the `btn` class.

<!-- `{% endfunc %}` -->


# Go
<!-- `{% func WebsiteGo() %}` -->
**[https://go.dev/](https://go.dev/)**<!-- target="_blank" -->
<br>

Go is a statically compiled language designed at Google and it has become my favorite language.

It is the tool that glues everything together in this project and I have used some of the community packages for this project

<br>

## sqlc

This tool generates Go structs and query functions with types and fields that are mapped using the input sql files that contain the `create table` commands and the queries.

<br>

What I really find cool is that **it also auto generates custom Go structs in case the query consists of extracting columns from multiple different tables to match the corresponding fields and types**<!-- class="bg-green-200" -->: you end up having the structs and the query functions with type safety without writing any Go boilerplate code.

<br>

**[https://github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)**<!-- target="_blank" -->

<br>

## quicktemplate

Go already has a template package in the standard library to generate text but I don't like that some of the expressions aren't Go-like: for example the 'if' equivalent in the template is **`{{if condition}} Stuff {{end}}`** and if you want to use a custom written Go function that returns a boolean as the `condition` then you need to put it in a map from Go, pass the map to the Go template variable and call the map key in the if condition inside the curly braces.

<br>

**[https://github.com/valyala/quicktemplate](https://github.com/valyala/quicktemplate)**<!-- target="_blank" -->

<br>

**quicktemplate addresses this by letting you use the Go language inside the curly braces in the template files.**<!-- class="bg-orange-200" -->

<br>

In the template files you can define functions, structs and insert some raw Go code.

**The tool (executed from the command line) then generates .go files that contains all the templates definitions (functions, structs etc.)**<!-- class="bg-yellow-200" --> and can be exported and used inside the project.

<br>

**These template functions and structs are compiled in the final binary**<!-- class="bg-cyan-200" -->: it is possible to run the final executable binary and start the server without being dependent on any **`.html`** file in the directory that needs to be read and parsed (with all the performance implications).

<br>

The drawback is that the intermediate parsing step (from `html` to `go`) is always needed even for small on-the-fly changes, reason why I am autorestarting the server on file save using [**`nodemon`**](https://nodemon.io/).


<br>


## chi

chi is a lightweight, idiomatic and composable router for building Go HTTP services.

**[https://go-chi.io/#/README](https://go-chi.io/#/README)**<!-- target="_blank" -->

<br>

It makes it easy to group subroutes and to apply some common middlewares.

<br>

What I also like about chi is the [**Walk**](https://pkg.go.dev/github.com/go-chi/chi#Walk)<!-- target="_blank" class="animate-pulse" --> function: combining this with the standard library [**reflect**](https://pkg.go.dev/reflect)<!-- target="_blank" class="animate-pulse" --> and [**runtime**](https://pkg.go.dev/runtime)<!-- target="_blank" class="animate-pulse" --> packages makes it easy to **automatically generate the documentation of the application's APIs with all the exact file and line references of the functions in the codebase**<!-- class="bg-teal-200" -->.

<br>

Example for this application: **[https://github.com/Jiang-Gianni/gianni-jiang/blob/main/docRoutes.md](https://github.com/Jiang-Gianni/gianni-jiang/blob/main/docRoutes.md)**<!-- target="_blank" -->


<br>

<!-- `{% endfunc %}` -->


# Demo Todo

<!-- `{% func WebsiteDemoTodo() %}` -->

Here is a demo of a Todo application that operates on the database.

<br>


Click the **Create New Todo**<!-- class="btn btn-accent" --> button to generate a new Todo or click one of the Todos in the table if you want to **update**<!-- class="btn btn-primary" --> or to **delete**<!-- class="btn btn-error" --> it.

<br>

<!-- `{%= TodoDemo() %}` -->

<!-- `{% endfunc %}` -->


# Demo Number
<!-- `{% func WebsiteDemoNumber() %}` -->
Simple HTTP call that retrieves some trivia from [**numbersapi**](http://numbersapi.com)<!-- class="bg-blue-200 animate-pulse" --> of a random number.

<br>

<!-- `{%= NumbersApi() %}` -->

<br>

<!-- `{% endfunc %}` -->


# Deploy
<!-- `{% func WebsiteDeploy() %}` -->

In order to make this app publicly available, I needed to deploy it somewhere.

<br>

After researching and trying various services I think that the best platforms **with free base plans and Go language support** are **Vercel** and **Render**.

<br>

## Vercel

Vercel is commonly used to deploy JavaScript based Web application, but **it offers a Go Runtime with Serverless Functions**<!-- class="bg-green-200" -->.

**The downside is that Serverless Function have a maximum execution duration, so it is not possible to maintain a WebSocket connection.**<!-- class="bg-red-200" -->

<br>

**[https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go](https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go)**<!-- target="_blank" class="animate-pulse bg-gray-200" -->

<br>


## Render

Render has **support for Docker and for many different programming languages**<!-- class="bg-green-200" -->.

You can define a build step of the application and the start command and it is possible to maintain an active session while in use.

**The downside is that if the server doesn't receive any request for a period of time then the execution stops and the cold start may take time (even more than 10 seconds).**<!-- class="bg-red-200" -->

<br>

**[https://render.com/docs](https://render.com/docs)**<!-- target="_blank" class="animate-pulse bg-gray-200" -->

<br>
<!-- `{% endfunc %}` -->