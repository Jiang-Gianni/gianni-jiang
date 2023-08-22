# Introduction
<!-- `{% func WebsiteIntroduction() %}` -->

This is a website that i made to showcase my projects.

You can jump directly to the section of the project you want to see by clicking the <!-- class="bg-teal-300 rounded-s" -->**menu on the left**

<br>

## CHATG Stack

To create this website I used the following tech stack:
<!-- class="list-disc list-inside" -->
* _<!-- class="bg-orange-200" -->*CockroachDB*_
* _<!-- class="bg-green-200" -->*HTMX*_
* _<!-- class="bg-green-200" -->*AlpineJS*_
* _<!-- class="bg-green-200" -->*TailwindCSS*_
* _<!-- class="bg-blue-200" -->*Go*_

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

**<!-- target="_blank" -->[https://www.cockroachlabs.com/](https://www.cockroachlabs.com/)**

<br>

CockroachDB is a commercial distributed PostgreSQL database management system and it is written in Go.

<br>

I am using the database for some CRUD operations of a <!-- hx-get="website/demo" hx-target="#contents" class="btn btn-accent" -->**Demo Todo application** and
some <!-- hx-get="website/feedback" hx-target="#contents" class="btn btn-accent" -->**feedbacks**.

<br>

I am not using it to store all this text contents, since I feel it would be rather bothersome to update a table rows for each content change.

<br>

With Cockroach DB the database cluster can be created in different regions in the world and they offer a great serverless free base plan: **10 GB and 50M request units per month**.

<br>

(Un)fortunately this website won't get many views and I won't ever reach that cap, so I don't have to worry about it.

<br>

<div x-data="{show:false}"><div>

**<!-- class="btn btn-primary" x-show="!show" @click="show=true" -->Me Right Now**
**<!-- class="mx-auto h-full" x-show="show"-->![Me Right Now](./assets/website/harold.jpg)**

<!-- `{% endfunc %}` -->


# HTMX
<!-- `{% func WebsiteHTMX() %}` -->
**<!-- target="_blank" -->[https://htmx.org/](https://htmx.org/)**

<br>

HTMX is a UI JavaScript library for the web that lets any HTML element make an HTTP request by simply adding some attributes.

<br>

If you right click on the menu on the left and select 'Inspect' you can see that there are some attributes like **`hx-get`** and **`hx-target`**. A click on the element will trigger a GET HTTP request to the endpoint value: <!-- class="bg-yellow-200" -->**the corresponding HTTP response contents will then substitute the element defined by the `hx-target` attribute**.

<br>

This means that the HTTP response body must be in form of HTML.

<br>

*<!-- class="bg-orange-200" -->**Any HTML element can trigger an HTTP request and all the HTTP verbs can be used (POST, GET, PUT, PATCH, and DELETE)*** in contrast to plain HTML where only POST (with 'form' tags) and GET (with 'a' tags) are possible.

<br>

## JavaScript size

There is no need on the client side to manipulate the response, **no json to be parsed**: this means that the page will load faster since <!-- class="bg-cyan-200" -->**there is no javascript (other than htmx) to be executed**.

If you press F12, go to 'Network', and press Ctrl + Shift + R to hard reload the page you'll see that <!-- class="bg-green-200" -->**the JavaScript size for htmx is around 15-16 kB**.

<br>

Now try to do the same on **<!-- target="_blank" -->[*https://create-react-template.vercel.app*](https://create-react-template.vercel.app)**

<br>

You'll notice <!-- class="bg-red-200" -->**almost 50 kB of JavaScript**.

Imagine what happens as the app scales: with React (or other similar SPA frameworks) even more JavaScript will be shipped to the client, while with HTMX it will stay the same.

<br>

## Framework

Another advantage of HTMX is that you just need those 15 kB of JavaScript.

No need for any **`npm install ...`** which can easily cause the local **`node_modules`** folder to exceed 100 MB of size when using frameworks like React.<!-- class="mx-auto h-full" -->![Node Modules](./assets/website/node-modules.png)

<br>


## Server Side Rendering

With HTMX the web contents are rendered on the server which means that **it is possible to use any backend language** as long as the HTML contents are delivered to the client through the HTTP response: the frontend is not bound to the use of JavaScript anymore.

<br>

Many languages also offer some templating tools, so generating the needed HTML with the data obtained from the database or other services gets easier (more on this later in the Go section).



<br>

<!-- `{% endfunc %}` -->

# Alpine JS
<!-- `{% func WebsiteAlpineJS() %}` -->
**<!-- target="_blank" -->[https://alpinejs.dev/](https://alpinejs.dev/)**
<br>

Alpine JS is a lightweight JavaScript framework: the compressed bundle is <!-- class="bg-green-200" -->**around 16 kB in size**.

<br>

Alpine JS lets you sprinkle a little bit of JavaScript where it is needed: if you inspect the menu element on the left (right click + 'Inspect') you'll see some attributes like <!-- class="bg-yellow-200" -->**'x-data', 'x-show', 'x-on' and 'x-transition'**.

<br>

After a click on the side menu, HTMX updates the content on the right and Alpine JS updates the local state of the menu, so that the current project and section are highlighted with different background color.

<br>

It is also possible to update the side menu through HTMX but I felt like in cases like this it was way easier to use a bit of JavaScript, especially with Alpine JS short syntax.

<br>

Combining Alpine JS, HTMX and Tailwind CSS (next section) makes it possible to generate the following component just by adding attributes to HTML elements.

<!-- `{%= AlpineCountToggle() %}` -->

<!-- `{% endfunc %}` -->

# TailwindCSS
<!-- `{% func WebsiteTailwindCSS() %}` -->
**<!-- target="_blank" -->[https://tailwindcss.com/](https://tailwindcss.com/)**
<br>

Tailwind CSS is an open source CSS framework that creates utility CSS classes.

<br>

It lets you add some classes to an HTML element rather than to write the element selector (tag / id / class) and its styling in a css file or inside a style tag: the HTML element gets more congested (especially with HTMX and Alpine) but I prefer writing 3-5 lines of attributes/values rather than to write 10-20 lines of CSS or JS (I am not fond of either).

## Example


The following displayed element is styled by adding the classes:

**`animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200`**

<br>

<div class="animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200">Hello World</div>

## CSS Size

It is possible to import the entire Tailwind library of utility classes by CDN but Tailwind offers a tool to optimize for production so that only the used classes are imported.

## DaysiUI

I am also using **<!-- class="btn btn-accent" target="_blank" -->[DaisyUI](https://daisyui.com/)** which is a library of prestyled components based on Tailwind CSS.

For example, all the buttons in this website are styled by simply adding the `btn` class.

<!-- `{% endfunc %}` -->


# Go
<!-- `{% func WebsiteGo() %}` -->
**<!-- target="_blank" -->[https://go.dev/](https://go.dev/)**
<br>

Go is a statically compiled language designed at Google and it has become my favorite language.

It is the tool that glues everything together in this project and I have used some of the community packages for this project

<br>

## sqlc

This tool generates Go structs and query functions with types and fields that are mapped using the input sql files that contain the `create table` commands and the queries.

<br>

What amazes me is that <!-- class="bg-green-200" -->**it also auto generates custom Go structs in case the query consists of joining multiple tables to match the fields and types**: you end up having the structs and the query functions with type safety without writing any Go boilerplate code.

<br>

**<!-- target="_blank" -->[https://github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)**

<br>

## quicktemplate

Go already has a template package in the standard library to generate text but I don't like that some of the expressions aren't Go-like: for example the 'if' equivalent in the template is **`{{if condition}} Stuff {{end}}`** and if you want to use a custom written function that returns a boolean as the condition then you need to put it in a map from Go, pass the map to the Go template variable and call the map key in the if condition inside the curly braces.

<br>

**<!-- target="_blank" -->[https://github.com/valyala/quicktemplate](https://github.com/valyala/quicktemplate)**

<br>

**<!-- class="bg-orange-200" -->*quicktemplate addresses this by letting you use the Go language inside the curly braces in the template files.***

<br>

In the template files you can define functions, structs and insert some raw Go code.

**<!-- class="bg-yellow-200" -->*The tool (executed from the command line) then generates `.go` files that contains all the templates definitions (functions, structs etc.)*** and can be exported and used inside the project.

<br>

*<!-- class="bg-cyan-200" -->**These template functions and structs are compiled in the final binary***: it is possible to run the final executable binary without any additional **`.html`** file dependency in the directory.

<br>

The drawback is that the intermediate parsing step (from `html` to `go`) is always needed even for small on-the-fly changes, reason why I am autorestarting the server on file save using [**`nodemon`**](https://nodemon.io/).

<!-- `{% endfunc %}` -->


# Demo Todo

<!-- `{% func WebsiteDemoTodo() %}` -->

Here is a demo of a Todo application that operates on the database.

<br>


Click the <!-- class="btn btn-accent" -->**Create New Todo** button to generate a new Todo or click one of the Todos in the table if you want
to <!-- class="btn btn-primary" -->**update**
or to <!-- class="btn btn-error" -->**delete** it.

<br>

<!-- `{%= TodoDemo() %}` -->

<!-- `{% endfunc %}` -->


# Demo Number
<!-- `{% func WebsiteDemoNumber() %}` -->
Simple HTTP call that retrieves some trivia from <!-- class="bg-blue-200" -->[**numbersapi**](http://numbersapi.com) of a random number.

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

Vercel is commonly used to deploy JavaScript based Web application, but <!-- class="bg-green-200" -->**it offers a Go Runtime with Serverless Functions**.

**<!-- class="bg-red-200" -->_The downside is that Serverless Function have a maximum execution duration, so it is not possible to maintain a WebSocket connection._**

<br>

**<!-- target="_blank" -->[https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go](https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go)**

<br>


## Render

Render has <!-- class="bg-green-200" -->**support for Docker and for many different programming languages**.

You can define a build step of the application and the start command and it is possible to maintain an active session while in use.

**<!-- class="bg-red-200" -->_The downside is that if the server doesn't receive any request for a period of time then the execution stops and the cold start may take time (even more than 10 seconds)._**

<br>

**<!-- target="_blank" -->[https://render.com/docs](https://render.com/docs)**

<br>
<!-- `{% endfunc %}` -->