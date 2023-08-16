// Code generated by qtc from "feedback.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/feedback.html:1
package views

//line views/feedback.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/feedback.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/feedback.html:1
func StreamFeedback(qw422016 *qt422016.Writer, project string) {
//line views/feedback.html:1
	qw422016.N().S(`
<div class="justify-center flex gap-20">

    <!-- Description -->
    <div class="form-control">
        <label class="label">
            <span class="label-text">Feedback</span>
        </label>
        <textarea name="description" type="text" placeholder="Description"
            class="textarea textarea-bordered textarea-lg w-full max-w-xs"></textarea>
    </div>

</div>

<!-- Submit Feedback -->
<div class="flex justify-center gap-20 p-4">
    <button hx-post="`)
//line views/feedback.html:17
	qw422016.E().S(project)
//line views/feedback.html:17
	qw422016.N().S(`/feedback" hx-include="[name='description']" hx-target="#submitted"
        class="btn btn-accent">
        Submit
    </button>
</div>

<!-- Submitted -->
<div id="submitted"></div>
`)
//line views/feedback.html:25
}

//line views/feedback.html:25
func WriteFeedback(qq422016 qtio422016.Writer, project string) {
//line views/feedback.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/feedback.html:25
	StreamFeedback(qw422016, project)
//line views/feedback.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/feedback.html:25
}

//line views/feedback.html:25
func Feedback(project string) string {
//line views/feedback.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/feedback.html:25
	WriteFeedback(qb422016, project)
//line views/feedback.html:25
	qs422016 := string(qb422016.B)
//line views/feedback.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/feedback.html:25
	return qs422016
//line views/feedback.html:25
}

//line views/feedback.html:27
func StreamSubmitted(qw422016 *qt422016.Writer) {
//line views/feedback.html:27
	qw422016.N().S(`
<div id="submitted" class="transition-opacity duration-500 ease-in-out p-4 mb-4 text-teal-600 rounded-lg bg-orange-100"
    hx-ext="class-tools" classes="add opacity-0:1s">
    Thank you for your submission
</div>

`)
//line views/feedback.html:33
}

//line views/feedback.html:33
func WriteSubmitted(qq422016 qtio422016.Writer) {
//line views/feedback.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/feedback.html:33
	StreamSubmitted(qw422016)
//line views/feedback.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/feedback.html:33
}

//line views/feedback.html:33
func Submitted() string {
//line views/feedback.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/feedback.html:33
	WriteSubmitted(qb422016)
//line views/feedback.html:33
	qs422016 := string(qb422016.B)
//line views/feedback.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/feedback.html:33
	return qs422016
//line views/feedback.html:33
}
