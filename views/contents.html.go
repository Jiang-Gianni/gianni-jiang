// Code generated by qtc from "contents.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/contents.html:1
package views

//line views/contents.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/contents.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/contents.html:1
func StreamCommonContents(qw422016 *qt422016.Writer, section Section) {
//line views/contents.html:1
	qw422016.N().S(`
<div class="transition-opacity duration-500 ease-in-out opacity-0" hx-ext="class-tools" classes="remove opacity-0:0s">

    <!-- Section Title -->
    <h1 class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl">
        `)
//line views/contents.html:6
	qw422016.E().S(section.Title)
//line views/contents.html:6
	qw422016.N().S(`</h1>

    `)
//line views/contents.html:8
	for _, content := range section.Contents {
//line views/contents.html:8
		qw422016.N().S(`

    <br>

    `)
//line views/contents.html:12
		if content.Subtitle != "" {
//line views/contents.html:12
			qw422016.N().S(`
    <h2 class="text-4xl font-extrabold">`)
//line views/contents.html:13
			qw422016.E().S(content.Subtitle)
//line views/contents.html:13
			qw422016.N().S(`</h2>
    <br>
    `)
//line views/contents.html:15
		}
//line views/contents.html:15
		qw422016.N().S(`

    `)
//line views/contents.html:17
		if len(content.TextList) != 0 {
//line views/contents.html:17
			qw422016.N().S(`
    <div>`)
//line views/contents.html:18
			StreamText(qw422016, content.TextList)
//line views/contents.html:18
			qw422016.N().S(`</div>
    <br>
    `)
//line views/contents.html:20
		}
//line views/contents.html:20
		qw422016.N().S(`

    `)
//line views/contents.html:22
		if content.Img != "" {
//line views/contents.html:22
			qw422016.N().S(`
    <img class="mx-auto h-full" src="`)
//line views/contents.html:23
			qw422016.E().S(content.Img)
//line views/contents.html:23
			qw422016.N().S(`" alt="`)
//line views/contents.html:23
			qw422016.E().S(content.Img)
//line views/contents.html:23
			qw422016.N().S(`">
    <br>
    `)
//line views/contents.html:25
		}
//line views/contents.html:25
		qw422016.N().S(`

    `)
//line views/contents.html:27
		if content.RawHtml != "" {
//line views/contents.html:27
			qw422016.N().S(`
    `)
//line views/contents.html:28
			qw422016.N().S(content.RawHtml)
//line views/contents.html:28
			qw422016.N().S(`
    <br>
    `)
//line views/contents.html:30
		}
//line views/contents.html:30
		qw422016.N().S(`

    `)
//line views/contents.html:32
	}
//line views/contents.html:32
	qw422016.N().S(`
</div>
`)
//line views/contents.html:34
}

//line views/contents.html:34
func WriteCommonContents(qq422016 qtio422016.Writer, section Section) {
//line views/contents.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/contents.html:34
	StreamCommonContents(qw422016, section)
//line views/contents.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/contents.html:34
}

//line views/contents.html:34
func CommonContents(section Section) string {
//line views/contents.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/contents.html:34
	WriteCommonContents(qb422016, section)
//line views/contents.html:34
	qs422016 := string(qb422016.B)
//line views/contents.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/contents.html:34
	return qs422016
//line views/contents.html:34
}
