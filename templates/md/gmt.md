# Introduction
<!-- `{% func GmtIntroduction() %}` -->

_"Hey, I am developing a web app."_

<br>

_"Cool! What are you using for your frontend? [**React**](https://react.dev/)<!-- class="btn bg-blue-200" -->
? [**Vue**](https://vuejs.org/)<!-- class="btn bg-green-200" -->
? [**Svelte**](https://svelte.dev/)<!-- class="btn bg-orange-200" -->?"_

<br>

<div x-data="{show:false}"><div>

**The obvious answer**<!-- class="btn btn-primary" x-show="!show" @click="show=true" -->
**![Gigachad](./assets/gmt/gigachad.jpg)**<!-- class="mx-auto h-full" x-show="show"-->

<!-- `{% endfunc %}` -->

# Why?
<!-- `{% func GmtWhy() %}` -->

[**Markdown**](https://en.wikipedia.org/wiki/Markdown)<!-- class="btn bg-blue-200 animate-pulse" target="_blank" --> is a markup language to write formatted text.
<br>

<br>

Writing Markdown files (**`.md`** extension) using a text editor that has some plugins for it feels way better compared to many other options: Markdown is not as plain as a `.txt` file, it is not slow at loading like MS Word, and best thing of all, the files can be converted into **`.html`**: they can be displayed in the browser and if the links are provided inside the contents then it becomes easy to navigate back and forth between files.

<br>

It allows for some shorthand syntax like
```md
**bold**, _italic_, ~~strikethrough~~
```

<br>

which, after using a Markdown processor are then transformed into html elements

```md
<strong>bold</strong>, <em>italic</em>, <del>strikethrough</del>
```

<br>

and will look like **bold**, _italic_, ~~strikethrough~~ on the browser.


<br>


## Markdown instead of HTML

It is true that Markdown can be converted to HTML but the conversion just parses the contents into the corresponding HTML tags.

<br>

The interactions (JavaScript) and the styling (CSS) are both missing in the final converted **`.hml`** file.

<br>

Then I remembered how HTMX, Tailwind CSS and Alpine JS work: the interactions and the styling can be implemented by adding attributes to the HTML element.

<br>

But the problem is that in order to add attributes from Markdown you have to encapsulate the content inside an HTML element (yes, in Markdown you can use plain HTML) and add what you need to that element.

Another option is to use some [**specific Markdown processors**](https://boringrails.com/tips/jekyll-css-class)<!-- target="_blank" class="bg-orange-200 animate-pulse" --> that have
a [**set of rules**](https://kramdown.gettalong.org/syntax.html#block-ials)<!-- target="_blank" class="bg-yellow-200 animate-pulse" --> to accomplish the task.

<br>

**I didn't like the syntax so I decided to try and build a tool for it.**

<!-- `{% endfunc %}` -->


# Parser
<!-- `{% func GmtParser() %}` -->
The first step is to have a basic Markdown processor with functionalities that can be extended.

<br>

I started reading the book [**Writing An Interpreter In Go**](https://interpreterbook.com/)<!-- class="bg-teal-300 animate-pulse" target="_blank" --> that explains the steps to build a programming language interpreter from scratch:

<br>


<!-- class="list-disc list-inside" -->
* Define the **tokens** and the **keywords** to look for
* Build a lexer that reads and tokenize the contents (=splitting the contents into fragments based on the tokens/keywords)
* Parse the tokens and build an **Abstract Syntax Tree** (AST) of nodes from the sequence of tokens and the syntax of the content's programming language
* Render the elements of the tree into the new programming language

<br>

I tried to apply these concepts to parse Markdown into HTML.

<br>

**The Result**<!-- class="btn btn-primary" hx-get="gmt/result" hx-swap="outerHTML" -->

<!-- `{% endfunc %}` -->


<!-- `{% func GmtResult() %}` -->

After a weekend I **successully**<!-- class="bg-green-200" -->
... **gave up**<!-- class="bg-red-200 transition-opacity duration-500 ease-in-out opacity-0" hx-ext="class-tools" classes="remove opacity-0:2s" -->

<!-- `{% endfunc %}` -->


# Reason
<!-- `{% func GmtReason() %}` -->

The reason is that there are too many rules to be implemented in the parsing step. [**GitHub Flavored Markdown Spec**](https://github.github.com/gfm/)<!-- target="_blank" class="bg-purple-200 animate-pulse" --> has a huge list of them.

<br>

Most importantly, there are already some high quality Markdown parser projects out there.

<br>

The one I chose to use is [**goldmark**](https://github.com/yuin/goldmark/)<!-- target="_blank" class="bg-orange-200 animate-pulse" -->, written in Go. It has many parameters that can be set and many extensions already available.

<br>

My favorite feature is that it supports [**Pikchr**](https://pikchr.org/home/doc/trunk/homepage.md)<!-- target="_blank" class="bg-green-200 animate-pulse" -->, a tool to sketch diagrams like the one on the first page of the website.

<br>

I tried [**Mermaid JS**](https://mermaid.js.org/intro/)<!-- target="_blank" class="bg-red-200 animate-pulse" --> in the past to draw the diagrams.

It has a nice syntax but it executes some JavaScript in order to render the contents.

<br>

I remember using Mermaid's [**Entity Relationship Diagram**](https://mermaid.js.org/syntax/entityRelationshipDiagram.html)<!-- class="bg-pink-200 animate-pulse" --> to write a Markdown file with all the tables and columns of a database (the client connection was too slow and I had to be connected on VPN, so I wanted to have a local list of the structure): the converted HTML file opened on the browser was taking too much time to render the contents. I think there were around 140 tables.

<br>

**Pikchr** on the other hand needs an intermediate parsing step (there is no JavaScript executed on the browser): the final result is an **`svg`** element, so the diagram is ready to be displayed when opened in the browser.

<br>

<br>

<br>

Going back to the Markdown parsing tool... what is needed now is to find a **mechanism to insert attributes**<!-- class="bg-cyan-200" --> from Markdown to the final converted HTML

<!-- `{% endfunc %}` -->


# Implementation
<!-- `{% func GmtImplementation() %}` -->

In the end I opted to make it so that the commented contents get *injected* to the previous or next element (based on cases) using some regexes (this step happens after the first HTML is parsed by **goldmark**'s tool).

<br>

The repo: [**Github**](https://github.com/Jiang-Gianni/gmt)<!-- target="_blank" class="bg-teal-300 animate-pulse" -->

<br>

I also added some Tailwind CSS: I extracted the utility classes definition into the codebase.

(Hence the name `gmt`: Go, Markdown, Tailwind)

The Tailwind styling definition is added in a **`style`** tag on top of the converted HTML in case they appear as classes.

<br>

I have not tested many Tailwind classes and probably not all of them work correctly at the moment.

The ones I needed to use (basic stuff like bg-color, text-color etc) seem to work fine, so I am satisfied with those.

<br>

## Examples

```md
# My blue title
<!-- class="text-blue-700" -->
```

<br>

is parsed as:

<br>

```html
<style>.text-blue-700{--text-opacity:1;color:#2b6cb0;color:rgba(43,108,176,var(--text-opacity))}</style>
<h1 id="my-blue-title"  class="text-blue-700" >My blue title</h1 id="my-blue-title">
```

<br>


and displayed as (depending on applied `h1` base styling)

# My blue title
<!-- class="text-blue-700" -->


<br>


I also made it so that comments beteen two backticks `<!-- `\`comment\`` `--> don't get injected but simply inserted in place so that other templating tools can be applied
([**quicktemplate**](https://github.com/valyala/quicktemplate/)<!-- class="bg-cyan-200 animate-pulse" --> is the one I am currently using).

<br>



<!-- class="mx-auto h-full" -->
```pikchr
box "Markdown" height 3cm width 3cm rad 10px
arrow -> right 200% "goldMark" above
box "HTML" height 3cm width 3cm rad 10px
arrow -> right 200% "gmt" above
box "HTML templates" "with attributes" "inside tags" height 3cm width 3cm rad 10px

arrow from previous box.s \
          down 2cm \
          then right 300% "Tailwind" center

arrow from previous box.e right 200% "quicktemplate" above

G: box "Go files" height 3cm width 3cm rad 10px

box "CSS file" height 3cm width 3cm rad 10px at 4cm below G
```


## Refactoring

I then decided to rewrite the text contents of this website in Markdown using my new tool.

This is the reason why you see many of those (maybe annoying) **buttons**<!-- class="btn btn-accent" --> and **highlighted texts**<!-- class="bg-orange-200" -->


<!-- `{% endfunc %}` -->