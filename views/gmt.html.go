// Code generated by qtc from "gmt.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <style></style>

//line views/gmt.html:2
package views

//line views/gmt.html:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/gmt.html:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/gmt.html:2
func StreamGmtIntroduction(qw422016 *qt422016.Writer) {
//line views/gmt.html:2
	qw422016.N().S(`
`)
//line views/gmt.html:3
}

//line views/gmt.html:3
func WriteGmtIntroduction(qq422016 qtio422016.Writer) {
//line views/gmt.html:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/gmt.html:3
	StreamGmtIntroduction(qw422016)
//line views/gmt.html:3
	qt422016.ReleaseWriter(qw422016)
//line views/gmt.html:3
}

//line views/gmt.html:3
func GmtIntroduction() string {
//line views/gmt.html:3
	qb422016 := qt422016.AcquireByteBuffer()
//line views/gmt.html:3
	WriteGmtIntroduction(qb422016)
//line views/gmt.html:3
	qs422016 := string(qb422016.B)
//line views/gmt.html:3
	qt422016.ReleaseByteBuffer(qb422016)
//line views/gmt.html:3
	return qs422016
//line views/gmt.html:3
}