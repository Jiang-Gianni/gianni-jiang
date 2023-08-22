# Introduction
<!-- `{% func GmtIntroduction() %}` -->

_"Hey, I am developing a web app."_

<br>

_"Cool! What are you using for your frontend? <!-- class="btn bg-blue-200" -->[**React**](https://react.dev/)
? <!-- class="btn bg-green-200" -->[**Vue**](https://vuejs.org/)
? <!-- class="btn bg-orange-200" -->[**Svelte**](https://svelte.dev/)?"_

<br>

<div x-data="{show:false}"><div>

**<!-- class="btn btn-primary" x-show="!show" @click="show=true" -->The obvious answer**
**<!-- class="mx-auto h-full" x-show="show"-->![Gigachad](./assets/gmt/gigachad.jpg)**

<!-- `{% endfunc %}` -->

# Why?
<!-- `{% func GmtWhy() %}` -->

_<!-- class="btn bg-blue-200" target="_blank" -->[**Markdown**](https://en.wikipedia.org/wiki/Markdown)_ is a markup language to write formatted text.
<br>

<br>

Writing Markdown files (**`.md`** extensions) using a text editor that has some plugins for it feels good compared to the other options: Markdown is not as plain as a `.txt` file, it is not slow at loading like MS Word, and best thing of all, the files can be converted into **`.html`**: they can be displayed in the browser and if the links are provided inside the contents then it becomes easy to navigate back and forth between them.

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

It is true that Markdown can be converted to HTML but the conversion just parse the contents into the corresponding HTML tags.

<br>

The interactions (JavaScript) and the styling (CSS) are both missing in the final converted **`.hml`** file.

<br>

Then I remembered how HTMX, Tailwind CSS and Alpine JS work: the interactions and the styling can be implemented by adding attributes to the HTML element.

<br>

But the problem is that in order to add attributes from Markdown you have to encapsulate the content inside an HTML element (yes, in Markdown you can use plain HTML) and add what you need in that element.

Another option is to use some <!-- target="_blank" class="bg-orange-200 animate-pulse" -->[**specific Markdown processors**](https://boringrails.com/tips/jekyll-css-class) that have
a <!-- target="_blank" class="bg-yellow-200 animate-pulse" -->[**set of rules**](https://kramdown.gettalong.org/syntax.html#block-ials) to accomplish the task.

<br>

**I didn't like the syntax so I decided to try and build a tool for it.**

<!-- `{% endfunc %}` -->


# Parser
<!-- `{% func GmtParser() %}` -->
The first step is to have a basic Markdown processor with functionalities that can be extended.

<br>

I started reading the book <!-- class="bg-teal-300 animate-pulse" target="_blank" -->[**Writing An Interpreter In Go**](https://interpreterbook.com/) that explains the steps to build a programming language interpreter from scratch:

<br>


<!-- class="list-disc list-inside" -->
* Define the **tokens** and the **keywords** to look for inside the contents
* Build a lexer that reads and tokenize the contents
* Parse the tokens and build an **Abstract Syntax Tree** (AST) of nodes from the sequence of tokens and the syntax of the content's programming language
* Render the elements of the tree into the new programming language

<br>

I tried to apply these concepts to parse Markdown into HTML.

<br>

_<!-- class="btn btn-primary" hx-get="gmt/result" hx-swap="outerHTML" -->**The Result**_

<!-- `{% endfunc %}` -->


<!-- `{% func GmtResult() %}` -->

After a weekend I <!-- class="bg-green-200" -->**successully**
... <!-- class="bg-red-200 transition-opacity duration-500 ease-in-out opacity-0" hx-ext="class-tools" classes="remove opacity-0:2s" -->**gave up**

<!-- `{% endfunc %}` -->


# Reason
<!-- `{% func GmtReason() %}` -->

The reason is that there are too many rules to be implemented in the parsing step.  <!-- target="_blank" class="bg-purple-200 animate-pulse" -->[**GitHub Flavored Markdown Spec**](https://github.github.com/gfm/) has a huge list of them.

<br>

Most importantly, there are already some high quality Markdown parser projects out there.

<br>

The one I chose to use is <!-- target="_blank" class="bg-orange-200 animate-pulse" -->[**goldmark**](https://github.com/yuin/goldmark/), written in Go. It has many parameters that can be set and many extensions already available.

<br>

My favorite feature is that it supports <!-- target="_blank" class="bg-green-200 animate-pulse" -->[**Pikchr**](https://pikchr.org/home/doc/trunk/homepage.md), a tool to sketch diagrams like the one on the first page of the website.

<br>

I tried <!-- target="_blank" class="bg-red-200 animate-pulse" -->[**Mermaid JS**](https://mermaid.js.org/intro/) in the past to draw the diagrams.

It has a nice syntax but it executes some JavaScript in order to render the contents.

I remember using Mermaid's [**Entity Relationship Diagram**](https://mermaid.js.org/syntax/entityRelationshipDiagram.html) to write a Markdown file with all the tables and columns of a database: the converted HTML file opened on the browser was taking too much time to render the contents.

<br>

**Pikchr** on the other hand needs an intermediate parsing step: the final result is an **`svg`** element, so the diagram is already ready to be displayed once it is opened in the browser.

<br>

<br>

<br>

Going back to the Markdown parsing tool... what is needed now is to find a mechanism to insert attributes from Markdown to the final converted HTML

<!-- `{% endfunc %}` -->


# Implementation
<!-- `{% func GmtImplementation() %}` -->

In the end I opted to make it so that the commented contents get *injected* to the next element using some regexes (this step happens after the first HTML is parsed by **goldmark**'s tool).

<br>

The repo: <!-- target="_blank" class="bg-teal-300 animate-pulse" -->[**Github**](https://github.com/Jiang-Gianni/gmt)

<br>

I also added some Tailwind CSS: I extracted the utility classes definition into the codebase.

The Tailwind styling definition is added in a **`style`** tag on top of the converted HTML in case they appear as classes.

<br>

I have not tested many Tailwind classes (and probably most of them won't work correctly yet), but the ones I needed to use (basic stuff like bg-color, text-color etc) seem to work fine.

<br>

## Example

```md
<!-- class="text-blue-700" -->
# My blue title
```

<br>

is parsed as:

<br>

```html
<style>.text-blue-700{--text-opacity:1;color:#2b6cb0;color:rgba(43,108,176,var(--text-opacity))}</style>
<h1 id="my-blue-title"  class="text-blue-700" >My blue title</h1>
```

and displayed as

<!-- class="text-blue-700" -->
# My blue title


## Refactoring

I then decided to rewrite the contents of this website in Markdown using my new tool.

This is the reason why you see all the <!-- class="btn btn-accent" -->**buttons**
and <!-- class="bg-orange-200" -->**highlighted texts**





<!-- `{% endfunc %}` -->