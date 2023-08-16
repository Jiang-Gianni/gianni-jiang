// Code generated by qtc from "todo.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/todo.html:1
package views

//line views/todo.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/todo.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/todo.html:1
func StreamTodoDemo(qw422016 *qt422016.Writer) {
//line views/todo.html:1
	qw422016.N().S(`
<div id="todo-contents" hx-get="todo" hx-trigger="load once"
    class="flex flex-col overflow-x-auto sm:-mx-6 lg:-mx-8 min-w-full py-2 sm:px-6 lg:px-8 overflow-hidden">
</div>
`)
//line views/todo.html:5
}

//line views/todo.html:5
func WriteTodoDemo(qq422016 qtio422016.Writer) {
//line views/todo.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/todo.html:5
	StreamTodoDemo(qw422016)
//line views/todo.html:5
	qt422016.ReleaseWriter(qw422016)
//line views/todo.html:5
}

//line views/todo.html:5
func TodoDemo() string {
//line views/todo.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line views/todo.html:5
	WriteTodoDemo(qb422016)
//line views/todo.html:5
	qs422016 := string(qb422016.B)
//line views/todo.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line views/todo.html:5
	return qs422016
//line views/todo.html:5
}