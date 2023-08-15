// Code generated by qtc from "render.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/render.html:1
package views

//line views/render.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/render.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/render.html:1
func StreamText(qw422016 *qt422016.Writer, textList []string) {
//line views/render.html:1
	qw422016.N().S(`
`)
//line views/render.html:2
	for _, text := range textList {
//line views/render.html:2
		qw422016.N().S(`
<p>`)
//line views/render.html:3
		qw422016.E().S(text)
//line views/render.html:3
		qw422016.N().S(`</p>
`)
//line views/render.html:4
	}
//line views/render.html:4
	qw422016.N().S(`
`)
//line views/render.html:5
}

//line views/render.html:5
func WriteText(qq422016 qtio422016.Writer, textList []string) {
//line views/render.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:5
	StreamText(qw422016, textList)
//line views/render.html:5
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:5
}

//line views/render.html:5
func Text(textList []string) string {
//line views/render.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:5
	WriteText(qb422016, textList)
//line views/render.html:5
	qs422016 := string(qb422016.B)
//line views/render.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:5
	return qs422016
//line views/render.html:5
}

//line views/render.html:8
func StreamList(qw422016 *qt422016.Writer, textList []string) {
//line views/render.html:8
	qw422016.N().S(`
<ul class="list-disc list-inside">
    `)
//line views/render.html:10
	for _, text := range textList {
//line views/render.html:10
		qw422016.N().S(`
    <li>`)
//line views/render.html:11
		qw422016.E().S(text)
//line views/render.html:11
		qw422016.N().S(`</li>
    `)
//line views/render.html:12
	}
//line views/render.html:12
	qw422016.N().S(`
</ul>
`)
//line views/render.html:14
}

//line views/render.html:14
func WriteList(qq422016 qtio422016.Writer, textList []string) {
//line views/render.html:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:14
	StreamList(qw422016, textList)
//line views/render.html:14
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:14
}

//line views/render.html:14
func List(textList []string) string {
//line views/render.html:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:14
	WriteList(qb422016, textList)
//line views/render.html:14
	qs422016 := string(qb422016.B)
//line views/render.html:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:14
	return qs422016
//line views/render.html:14
}

//line views/render.html:16
func StreamLink(qw422016 *qt422016.Writer, href string, text string, target string) {
//line views/render.html:16
	qw422016.N().S(`
<a href="`)
//line views/render.html:17
	qw422016.E().S(href)
//line views/render.html:17
	qw422016.N().S(`" target="`)
//line views/render.html:17
	qw422016.E().S(target)
//line views/render.html:17
	qw422016.N().S(`">`)
//line views/render.html:17
	qw422016.E().S(text)
//line views/render.html:17
	qw422016.N().S(`</a>
`)
//line views/render.html:18
}

//line views/render.html:18
func WriteLink(qq422016 qtio422016.Writer, href string, text string, target string) {
//line views/render.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:18
	StreamLink(qw422016, href, text, target)
//line views/render.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:18
}

//line views/render.html:18
func Link(href string, text string, target string) string {
//line views/render.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:18
	WriteLink(qb422016, href, text, target)
//line views/render.html:18
	qs422016 := string(qb422016.B)
//line views/render.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:18
	return qs422016
//line views/render.html:18
}

//line views/render.html:20
func StreamCode(qw422016 *qt422016.Writer, text string) {
//line views/render.html:20
	qw422016.N().S(`
<code>`)
//line views/render.html:21
	qw422016.E().S(text)
//line views/render.html:21
	qw422016.N().S(`</code>
`)
//line views/render.html:22
}

//line views/render.html:22
func WriteCode(qq422016 qtio422016.Writer, text string) {
//line views/render.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:22
	StreamCode(qw422016, text)
//line views/render.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:22
}

//line views/render.html:22
func Code(text string) string {
//line views/render.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:22
	WriteCode(qb422016, text)
//line views/render.html:22
	qs422016 := string(qb422016.B)
//line views/render.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:22
	return qs422016
//line views/render.html:22
}

//line views/render.html:24
func StreamStrong(qw422016 *qt422016.Writer, text string) {
//line views/render.html:24
	qw422016.N().S(`
<strong>`)
//line views/render.html:25
	qw422016.E().S(text)
//line views/render.html:25
	qw422016.N().S(`</strong>
`)
//line views/render.html:26
}

//line views/render.html:26
func WriteStrong(qq422016 qtio422016.Writer, text string) {
//line views/render.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:26
	StreamStrong(qw422016, text)
//line views/render.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:26
}

//line views/render.html:26
func Strong(text string) string {
//line views/render.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:26
	WriteStrong(qb422016, text)
//line views/render.html:26
	qs422016 := string(qb422016.B)
//line views/render.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:26
	return qs422016
//line views/render.html:26
}

//line views/render.html:28
func StreamImg(qw422016 *qt422016.Writer, src string) {
//line views/render.html:28
	qw422016.N().S(`
<img class="mx-auto h-full" src="`)
//line views/render.html:29
	qw422016.E().S(src)
//line views/render.html:29
	qw422016.N().S(`" alt="`)
//line views/render.html:29
	qw422016.E().S(src)
//line views/render.html:29
	qw422016.N().S(`">
`)
//line views/render.html:30
}

//line views/render.html:30
func WriteImg(qq422016 qtio422016.Writer, src string) {
//line views/render.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/render.html:30
	StreamImg(qw422016, src)
//line views/render.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/render.html:30
}

//line views/render.html:30
func Img(src string) string {
//line views/render.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/render.html:30
	WriteImg(qb422016, src)
//line views/render.html:30
	qs422016 := string(qb422016.B)
//line views/render.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/render.html:30
	return qs422016
//line views/render.html:30
}
