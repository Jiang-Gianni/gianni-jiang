// Code generated by qtc from "website.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <style>.bg-teal-300{--bg-opacity:1;background-color:#81e6d9;background-color:rgba(129,230,217,var(--bg-opacity))}.list-disc{list-style-type:disc}.list-inside{list-style-position:inside}.bg-orange-200{--bg-opacity:1;background-color:#feebc8;background-color:rgba(254,235,200,var(--bg-opacity))}.bg-green-200{--bg-opacity:1;background-color:#c6f6d5;background-color:rgba(198,246,213,var(--bg-opacity))}.bg-blue-200{--bg-opacity:1;background-color:#bee3f8;background-color:rgba(190,227,248,var(--bg-opacity))}.mx-auto{margin-left:auto;margin-right:auto}.h-full{height:100%}.bg-yellow-200{--bg-opacity:1;background-color:#fefcbf;background-color:rgba(254,252,191,var(--bg-opacity))}.bg-red-200{--bg-opacity:1;background-color:#fed7d7;background-color:rgba(254,215,215,var(--bg-opacity))}.animate-bounce{animation:bounce 1s infinite}@keyframes bounce{0%,100%{transform:translateY(-25%);animation-timing-function:cubic-bezier(.8,0,1,1)}50%{transform:none;animation-timing-function:cubic-bezier(0,0,.2,1)}}.font-bold{font-weight:700}.border-8{border-width:8px}.border-teal-400{--border-opacity:1;border-color:#4fd1c5;border-color:rgba(79,209,197,var(--border-opacity))}.text-red-700{--text-opacity:1;color:#c53030;color:rgba(197,48,48,var(--text-opacity))}</style>
// <h1 id="introduction">Introduction</h1>

//line views/website.html:3
package views

//line views/website.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/website.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/website.html:3
func StreamWebsiteIntroduction(qw422016 *qt422016.Writer) {
//line views/website.html:3
	qw422016.N().S(`
<p>This is a website that i made to showcase my projects.</p>
<p>You can jump directly to the section of the project you want to see by clicking the <strong  class="bg-teal-300 rounded-s" >menu on the left</strong></p>
<br>
<h2 id="chatg-stack">CHATG Stack</h2>
<p>To create this website I used the following tech stack:</p>
<ul  class="list-disc list-inside" >
<li><em><em  class="bg-orange-200" >CockroachDB</em></em></li>
<li><em><em  class="bg-green-200" >HTMX</em></em></li>
<li><em><em  class="bg-green-200" >AlpineJS</em></em></li>
<li><em><em  class="bg-green-200" >TailwindCSS</em></em></li>
<li><em><em  class="bg-blue-200" >Go</em></em></li>
</ul>
<br>
<div class="pikchr-svg mx-auto h-full"  style="max-width:802px"    ><svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 802.556 174.399">
<path d="M2,17L2,157A85 15 0 0 0 172 157L172,17A85 15 0 0 0 2 17A85 15 0 0 0 172 17"  style="fill:rgb(255,228,196);stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="87" y="88" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">CockroachDB</text>
<text x="87" y="108" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">(Database)</text>
<polygon points="172,87 183,82 183,91" style="fill:rgb(0,0,0)"/>
<polygon points="316,87 304,91 304,82" style="fill:rgb(0,0,0)"/>
<path d="M177,87L310,87"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
<path d="M331,172L471,172A15 15 0 0 0 486 157L486,17A15 15 0 0 0 471 2L331,2A15 15 0 0 0 316 17L316,157A15 15 0 0 0 331 172Z"  style="fill:rgb(135,206,250);stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="401" y="77" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Go</text>
<text x="401" y="97" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">(Server)</text>
<polygon points="486,87 497,82 497,91" style="fill:rgb(0,0,0)"/>
<polygon points="630,87 618,91 618,82" style="fill:rgb(0,0,0)"/>
<path d="M492,87L624,87"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
<path d="M630,172L800,172L800,2L630,2Z"  style="fill:rgb(144,238,144);stroke-width:2.16;stroke:rgb(0,0,0);" />
<text x="715" y="56" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">HTMX</text>
<text x="715" y="77" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Alpine JS</text>
<text x="715" y="97" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Tailwind Css</text>
<text x="715" y="117" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">(Client)</text>
</svg>
</div><br>
<p>I am still missing a P and a T to finally complete the <strong>CHATGPT</strong> stack.</p>
<p>If you have some ideas that I might like let me know.</p>
`)
//line views/website.html:39
}

//line views/website.html:39
func WriteWebsiteIntroduction(qq422016 qtio422016.Writer) {
//line views/website.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:39
	StreamWebsiteIntroduction(qw422016)
//line views/website.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:39
}

//line views/website.html:39
func WebsiteIntroduction() string {
//line views/website.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:39
	WriteWebsiteIntroduction(qb422016)
//line views/website.html:39
	qs422016 := string(qb422016.B)
//line views/website.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:39
	return qs422016
//line views/website.html:39
}

// <h1 id="cockroachdb">CockroachDB</h1>

//line views/website.html:41
func StreamWebsiteCockroachDB(qw422016 *qt422016.Writer) {
//line views/website.html:41
	qw422016.N().S(`
<p><strong><a href="https://www.cockroachlabs.com/"  target="_blank" >https://www.cockroachlabs.com/</a></strong></p>
<br>
<p>CockroachDB is a commercial distributed PostgreSQL database management system and it is written in Go.</p>
<br>
<p>I am using the database for some CRUD operations of a <strong  hx-get="website/demo" hx-target="#contents" class="btn btn-accent" >Demo Todo application</strong> and
some <strong  hx-get="website/feedback" hx-target="#contents" class="btn btn-accent" >feedbacks</strong>.</p>
<br>
<p>I am not using it to store all this text contents, since I feel it would be rather bothersome to update a table rows for each content change.</p>
<br>
<p>With Cockroach DB the database cluster can be created in different regions in the world and they offer a great serverless free base plan: <strong>10 GB and 50M request units per month</strong>.</p>
<br>
<p>(Un)fortunately this website won't get many views and I won't ever reach that cap, so I don't have to worry about it.</p>
<br>
<div x-data="{show:false}"><div>
<p><strong><div  class="btn btn-primary" x-show="!show" @click="show=true" >Me Right Now</div><strong>
<strong><img src="./assets/website/harold.jpg" alt="Me Right Now"  class="mx-auto h-full" x-show="show"></strong></p>
`)
//line views/website.html:58
}

//line views/website.html:58
func WriteWebsiteCockroachDB(qq422016 qtio422016.Writer) {
//line views/website.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:58
	StreamWebsiteCockroachDB(qw422016)
//line views/website.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:58
}

//line views/website.html:58
func WebsiteCockroachDB() string {
//line views/website.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:58
	WriteWebsiteCockroachDB(qb422016)
//line views/website.html:58
	qs422016 := string(qb422016.B)
//line views/website.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:58
	return qs422016
//line views/website.html:58
}

// <h1 id="htmx">HTMX</h1>

//line views/website.html:60
func StreamWebsiteHTMX(qw422016 *qt422016.Writer) {
//line views/website.html:60
	qw422016.N().S(`
<p><strong><a href="https://htmx.org/"  target="_blank" >https://htmx.org/</a></strong></p>
<br>
<p>HTMX is a UI JavaScript library for the web that lets any HTML element make an HTTP request by simply adding some attributes.</p>
<br>
<p>If you right click on the menu on the left and select 'Inspect' you can see that there are some attributes like <strong><code>hx-get</code></strong> and <strong><code>hx-target</code></strong>. A click on the element will trigger a GET HTTP request to the endpoint value: <strong  class="bg-yellow-200" >the corresponding HTTP response contents will then substitute the element defined by the <code>hx-target</code> attribute</strong>.</p>
<br>
<p>This means that the HTTP response body must be in form of HTML.</p>
<br>
<p><em><strong  class="bg-orange-200" >Any HTML element can trigger an HTTP request and all the HTTP verbs can be used (POST, GET, PUT, PATCH, and DELETE)</strong></em> in contrast to plain HTML where only POST (with 'form' tags) and GET (with 'a' tags) are possible.</p>
<br>
<h2 id="javascript-size">JavaScript size</h2>
<p>There is no need on the client side to manipulate the response, <strong>no json to be parsed</strong>: this means that the page will load faster since <strong  class="bg-cyan-200" >there is no javascript (other than htmx) to be executed</strong>.</p>
<p>If you press F12, go to 'Network', and press Ctrl + Shift + R to hard reload the page you'll see that <strong  class="bg-green-200" >the JavaScript size for htmx is around 15-16 kB</strong>.</p>
<br>
<p>Now try to do the same on <strong><a href="https://create-react-template.vercel.app"  target="_blank" ><em>https://create-react-template.vercel.app</em></a></strong></p>
<br>
<p>You'll notice <strong  class="bg-red-200" >almost 50 kB of JavaScript</strong>.</p>
<p>Imagine what happens as the app scales: with React (or other similar SPA frameworks) even more JavaScript will be shipped to the client, while with HTMX it will stay the same.</p>
<br>
<h2 id="framework">Framework</h2>
<p>Another advantage of HTMX is that you just need those 15 kB of JavaScript.</p>
<p>No need for any <strong><code>npm install ...</code></strong> which can easily cause the local <strong><code>node_modules</code></strong> folder to exceed 100 MB of size when using frameworks like React.<img src="./assets/website/node-modules.png" alt="Node Modules"  class="mx-auto h-full" ></p>
<br>
<h2 id="server-side-rendering">Server Side Rendering</h2>
<p>With HTMX the web contents are rendered on the server which means that <strong>it is possible to use any backend language</strong> as long as the HTML contents are delivered to the client through the HTTP response: the frontend is not bound to the use of JavaScript anymore.</p>
<br>
<p>Many languages also offer some templating tools, so generating the needed HTML with the data obtained from the database or other services gets easier (more on this later in the Go section).</p>
<br>
`)
//line views/website.html:89
}

//line views/website.html:89
func WriteWebsiteHTMX(qq422016 qtio422016.Writer) {
//line views/website.html:89
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:89
	StreamWebsiteHTMX(qw422016)
//line views/website.html:89
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:89
}

//line views/website.html:89
func WebsiteHTMX() string {
//line views/website.html:89
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:89
	WriteWebsiteHTMX(qb422016)
//line views/website.html:89
	qs422016 := string(qb422016.B)
//line views/website.html:89
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:89
	return qs422016
//line views/website.html:89
}

// <h1 id="alpine-js">Alpine JS</h1>

//line views/website.html:91
func StreamWebsiteAlpineJS(qw422016 *qt422016.Writer) {
//line views/website.html:91
	qw422016.N().S(`
<p><strong><a href="https://alpinejs.dev/"  target="_blank" >https://alpinejs.dev/</a></strong>
<br></p>
<p>Alpine JS is a lightweight JavaScript framework: the compressed bundle is <strong  class="bg-green-200" >around 16 kB in size</strong>.</p>
<br>
<p>Alpine JS lets you sprinkle a little bit of JavaScript where it is needed: if you inspect the menu element on the left (right click + 'Inspect') you'll see some attributes like <strong  class="bg-yellow-200" >'x-data', 'x-show', 'x-on' and 'x-transition'</strong>.</p>
<br>
<p>After a click on the side menu, HTMX updates the content on the right and Alpine JS updates the local state of the menu, so that the current project and section are highlighted with different background color.</p>
<br>
<p>It is also possible to update the side menu through HTMX but I felt like in cases like this it was way easier to use a bit of JavaScript, especially with Alpine JS short syntax.</p>
<br>
<p>Combining Alpine JS, HTMX and Tailwind CSS (next section) makes it possible to generate the following component just by adding attributes to HTML elements.</p>
`)
//line views/website.html:103
	StreamAlpineCountToggle(qw422016)
//line views/website.html:103
	qw422016.N().S(`
`)
//line views/website.html:104
}

//line views/website.html:104
func WriteWebsiteAlpineJS(qq422016 qtio422016.Writer) {
//line views/website.html:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:104
	StreamWebsiteAlpineJS(qw422016)
//line views/website.html:104
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:104
}

//line views/website.html:104
func WebsiteAlpineJS() string {
//line views/website.html:104
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:104
	WriteWebsiteAlpineJS(qb422016)
//line views/website.html:104
	qs422016 := string(qb422016.B)
//line views/website.html:104
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:104
	return qs422016
//line views/website.html:104
}

// <h1 id="tailwindcss">TailwindCSS</h1>

//line views/website.html:106
func StreamWebsiteTailwindCSS(qw422016 *qt422016.Writer) {
//line views/website.html:106
	qw422016.N().S(`
<p><strong><a href="https://tailwindcss.com/"  target="_blank" >https://tailwindcss.com/</a></strong>
<br></p>
<p>Tailwind CSS is an open source CSS framework that creates utility CSS classes.</p>
<br>
<p>It lets you add some classes to an HTML element rather than to write the element selector (tag / id / class) and its styling in a css file or inside a style tag: the HTML element gets more congested (especially with HTMX and Alpine) but I prefer writing 3-5 lines of attributes/values rather than to write 10-20 lines of CSS or JS (I am not fond of either).</p>
<h2 id="example">Example</h2>
<p>The following displayed element is styled by adding the classes:</p>
<p><strong><code>animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200</code></strong></p>
<br>
<div class="animate-bounce bg-orange-200 font-bold border-8 border-teal-400 text-red-700 hover:bg-purple-200">Hello World</div>
<h2 id="css-size">CSS Size</h2>
<p>It is possible to import the entire Tailwind library of utility classes by CDN but Tailwind offers a tool to optimize for production so that only the used classes are imported.</p>
<h2 id="daysiui">DaysiUI</h2>
<p>I am also using <strong><a href="https://daisyui.com/"  class="btn btn-accent" target="_blank" >DaisyUI</a></strong> which is a library of prestyled components based on Tailwind CSS.</p>
<p>For example, all the buttons in this website are styled by simply adding the <code>btn</code> class.</p>
`)
//line views/website.html:122
}

//line views/website.html:122
func WriteWebsiteTailwindCSS(qq422016 qtio422016.Writer) {
//line views/website.html:122
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:122
	StreamWebsiteTailwindCSS(qw422016)
//line views/website.html:122
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:122
}

//line views/website.html:122
func WebsiteTailwindCSS() string {
//line views/website.html:122
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:122
	WriteWebsiteTailwindCSS(qb422016)
//line views/website.html:122
	qs422016 := string(qb422016.B)
//line views/website.html:122
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:122
	return qs422016
//line views/website.html:122
}

// <h1 id="go">Go</h1>

//line views/website.html:124
func StreamWebsiteGo(qw422016 *qt422016.Writer) {
//line views/website.html:124
	qw422016.N().S(`
<p><strong><a href="https://go.dev/"  target="_blank" >https://go.dev/</a></strong>
<br></p>
<p>Go is a statically compiled language designed at Google and it has become my favorite language.</p>
<p>It is the tool that glues everything together in this project and I have used some of the community packages for this project</p>
<br>
<h2 id="sqlc">sqlc</h2>
<p>This tool generates Go structs and query functions with types and fields that are mapped using the input sql files that contain the <code>create table</code> commands and the queries.</p>
<br>
<p>What amazes me is that <strong  class="bg-green-200" >it also auto generates custom Go structs in case the query consists of joining multiple tables to match the fields and types</strong>: you end up having the structs and the query functions with type safety without writing any Go boilerplate code.</p>
<p><strong><a href="https://github.com/sqlc-dev/sqlc"  target="_blank" >https://github.com/sqlc-dev/sqlc</a></strong></p>
<br>
<h2 id="quicktemplate">quicktemplate</h2>
<p>Go already has a template package in the standard library to generate text but I don't like that some of the expressions aren't Go-like: for example the 'if' equivalent in the template is <strong><code>{{if condition}} Stuff {{end}}</code></strong> and if you want to use a custom written function that returns a boolean as the condition then you need to put it in a map from Go, pass the map to the Go template variable and call the map key in the if condition inside the curly braces.</p>
<br>
<p><strong><a href="https://github.com/valyala/quicktemplate"  target="_blank" >https://github.com/valyala/quicktemplate</a></strong></p>
<br>
<p><strong><em  class="bg-orange-200" >quicktemplate addresses this by letting you use the Go language inside the curly braces in the template files.</em></strong></p>
<br>
<p>In the template files you can define functions, structs and insert some raw Go code.</p>
<p><strong><em  class="bg-yellow-200" >The tool (executed from the command line) then generates <code>.go</code> files that contains all the templates definitions (functions, structs etc.)</em></strong> and can be exported and used inside the project.</p>
<br>
<p><em><strong  class="bg-cyan-200" >These template functions and structs are compiled in the final binary</strong></em>: it is possible to run the final executable binary without any additional <strong><code>.html</code></strong> file dependency in the directory.</p>
<br>
<p>The drawback is that the intermediate parsing step (from <code>html</code> to <code>go</code>) is always needed even for small on-the-fly changes, reason why I am autorestarting the server on file save using <a href="https://nodemon.io/"><strong><code>nodemon</code></strong></a>.</p>
`)
//line views/website.html:149
}

//line views/website.html:149
func WriteWebsiteGo(qq422016 qtio422016.Writer) {
//line views/website.html:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:149
	StreamWebsiteGo(qw422016)
//line views/website.html:149
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:149
}

//line views/website.html:149
func WebsiteGo() string {
//line views/website.html:149
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:149
	WriteWebsiteGo(qb422016)
//line views/website.html:149
	qs422016 := string(qb422016.B)
//line views/website.html:149
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:149
	return qs422016
//line views/website.html:149
}

// <h1 id="demo-todo">Demo Todo</h1>

//line views/website.html:151
func StreamWebsiteDemoTodo(qw422016 *qt422016.Writer) {
//line views/website.html:151
	qw422016.N().S(`
<p>Here is a demo of a Todo application that operates on the database.</p>
<br>
<p>Click the <strong  class="btn btn-accent" >Create New Todo</strong> button to generate a new Todo or click one of the Todos in the table if you want
to <strong  class="btn btn-primary" >update</strong>
or to <strong  class="btn btn-error" >delete</strong> it.</p>
<br>
`)
//line views/website.html:158
	StreamTodoDemo(qw422016)
//line views/website.html:158
	qw422016.N().S(`
`)
//line views/website.html:159
}

//line views/website.html:159
func WriteWebsiteDemoTodo(qq422016 qtio422016.Writer) {
//line views/website.html:159
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:159
	StreamWebsiteDemoTodo(qw422016)
//line views/website.html:159
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:159
}

//line views/website.html:159
func WebsiteDemoTodo() string {
//line views/website.html:159
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:159
	WriteWebsiteDemoTodo(qb422016)
//line views/website.html:159
	qs422016 := string(qb422016.B)
//line views/website.html:159
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:159
	return qs422016
//line views/website.html:159
}

// <h1 id="demo-number">Demo Number</h1>

//line views/website.html:161
func StreamWebsiteDemoNumber(qw422016 *qt422016.Writer) {
//line views/website.html:161
	qw422016.N().S(`
<p>Simple HTTP call that retrieves some trivia from <a href="http://numbersapi.com"  class="bg-blue-200" ><strong>numbersapi</strong></a> of a random number.</p>
<br>
`)
//line views/website.html:164
	StreamNumbersApi(qw422016)
//line views/website.html:164
	qw422016.N().S(`
<br>
`)
//line views/website.html:166
}

//line views/website.html:166
func WriteWebsiteDemoNumber(qq422016 qtio422016.Writer) {
//line views/website.html:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:166
	StreamWebsiteDemoNumber(qw422016)
//line views/website.html:166
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:166
}

//line views/website.html:166
func WebsiteDemoNumber() string {
//line views/website.html:166
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:166
	WriteWebsiteDemoNumber(qb422016)
//line views/website.html:166
	qs422016 := string(qb422016.B)
//line views/website.html:166
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:166
	return qs422016
//line views/website.html:166
}

// <h1 id="deploy">Deploy</h1>

//line views/website.html:168
func StreamWebsiteDeploy(qw422016 *qt422016.Writer) {
//line views/website.html:168
	qw422016.N().S(`
<p>In order to make this app publicly available, I needed to deploy it somewhere.</p>
<br>
<p>After researching and trying various services I think that the best platforms <strong>with free base plans and Go language support</strong> are <strong>Vercel</strong> and <strong>Render</strong>.</p>
<br>
<h2 id="vercel">Vercel</h2>
<p>Vercel is commonly used to deploy JavaScript based Web application, but <strong  class="bg-green-200" >it offers a Go Runtime with Serverless Functions</strong>.</p>
<p><strong><em  class="bg-red-200" >The downside is that Serverless Function have a maximum execution duration, so it is not possible to maintain a WebSocket connection.</em></strong></p>
<br>
<p><strong><a href="https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go"  target="_blank" >https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go</a></strong></p>
<br>
<h2 id="render">Render</h2>
<p>Render has <strong  class="bg-green-200" >support for Docker and for many different programming languages</strong>.</p>
<p>You can define a build step of the application and the start command and it is possible to maintain an active session while in use.</p>
<p><strong><em  class="bg-red-200" >The downside is that if the server doesn't receive any request for a period of time then the execution stops and the cold start may take time (even more than 10 seconds).</em></strong></p>
<br>
<p><strong><a href="https://render.com/docs"  target="_blank" >https://render.com/docs</a></strong></p>
<br>
`)
//line views/website.html:186
}

//line views/website.html:186
func WriteWebsiteDeploy(qq422016 qtio422016.Writer) {
//line views/website.html:186
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/website.html:186
	StreamWebsiteDeploy(qw422016)
//line views/website.html:186
	qt422016.ReleaseWriter(qw422016)
//line views/website.html:186
}

//line views/website.html:186
func WebsiteDeploy() string {
//line views/website.html:186
	qb422016 := qt422016.AcquireByteBuffer()
//line views/website.html:186
	WriteWebsiteDeploy(qb422016)
//line views/website.html:186
	qs422016 := string(qb422016.B)
//line views/website.html:186
	qt422016.ReleaseByteBuffer(qb422016)
//line views/website.html:186
	return qs422016
//line views/website.html:186
}