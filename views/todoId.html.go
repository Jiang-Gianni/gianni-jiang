// Code generated by qtc from "todoId.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/todoId.html:1
package views

//line views/todoId.html:1
import "github.com/Jiang-Gianni/gthc/db"

//line views/todoId.html:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/todoId.html:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/todoId.html:2
func StreamTodoId(qw422016 *qt422016.Writer, todo db.GetTodoRow, statusList []db.Status) {
//line views/todoId.html:2
	qw422016.N().S(`
<div class="justify-center flex gap-20">

    <!-- Description -->
    <div class="form-control">
        <label class="label">
            <span class="label-text">Description</span>
        </label>
        <textarea name="description" type="text" placeholder="Todo Description"
            class="textarea textarea-bordered textarea-lg w-full max-w-xs">`)
//line views/todoId.html:11
	qw422016.E().S(todo.Description)
//line views/todoId.html:11
	qw422016.N().S(`</textarea>
    </div>

    <!-- Status -->
    `)
//line views/todoId.html:15
	if len(statusList) > 0 {
//line views/todoId.html:15
		qw422016.N().S(`
    <div class="form-control w-full max-w-xs">
        <label class="label">
            <span class="label-text">Todo Status</span>
        </label>
        <select class="select select-bordered" name="status_id">
            `)
//line views/todoId.html:21
		for _, status := range statusList {
//line views/todoId.html:21
			qw422016.N().S(`
            <option value="`)
//line views/todoId.html:22
			qw422016.E().V(status.ID)
//line views/todoId.html:22
			qw422016.N().S(`" `)
//line views/todoId.html:22
			if status.ID == todo.StatusID {
//line views/todoId.html:22
				qw422016.N().S(`selected`)
//line views/todoId.html:22
			}
//line views/todoId.html:22
			qw422016.N().S(`>`)
//line views/todoId.html:22
			qw422016.E().S(status.Description)
//line views/todoId.html:23
			qw422016.N().S(`
            </option>

            `)
//line views/todoId.html:26
		}
//line views/todoId.html:26
		qw422016.N().S(`
        </select>
    </div>
    `)
//line views/todoId.html:29
	}
//line views/todoId.html:29
	qw422016.N().S(`

</div>

`)
//line views/todoId.html:33
	if todo.ID != 0 {
//line views/todoId.html:33
		qw422016.N().S(`
<!-- Updating or deleting a Todo -->

<div class="flex justify-center gap-20 p-4">

    <button hx-post="todo`)
//line views/todoId.html:38
		qw422016.E().V(todo.ID)
//line views/todoId.html:38
		qw422016.N().S(`" hx-include="[name='description'], [name='status_id']"
        hx-target="#todo-contents" class="btn btn-primary btn-wide">Update</button>

    <button class="btn btn-error btn-wide" onclick="del_btn.showModal()">Delete</button>
    <dialog id="del_btn" class="modal">
        <form method="dialog" class="modal-box">
            <h3 class="font-bold text-lg">Delete</h3>
            <p class="py-4">Click Confirm to delete this todo. Press ESC key or click Cancel to close</p>
            <div class="modal-action">
                <button class="btn">Cancel</button>
                <button hx-delete="todo`)
//line views/todoId.html:48
		qw422016.E().V(todo.ID)
//line views/todoId.html:48
		qw422016.N().S(`" hx-target="#todo-contents" class="btn btn-error">Confirm</button>
            </div>
        </form>
    </dialog>

</div>

`)
//line views/todoId.html:55
	} else {
//line views/todoId.html:55
		qw422016.N().S(`

<!-- Creating a new Todo -->

<div class="flex justify-center gap-20 p-4">

    <button hx-post="todo" hx-include="[name='description']" hx-target="#todo-contents" class="btn btn-accent">
        Create
    </button>

</div>

`)
//line views/todoId.html:67
	}
//line views/todoId.html:67
	qw422016.N().S(`
`)
//line views/todoId.html:68
}

//line views/todoId.html:68
func WriteTodoId(qq422016 qtio422016.Writer, todo db.GetTodoRow, statusList []db.Status) {
//line views/todoId.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/todoId.html:68
	StreamTodoId(qw422016, todo, statusList)
//line views/todoId.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/todoId.html:68
}

//line views/todoId.html:68
func TodoId(todo db.GetTodoRow, statusList []db.Status) string {
//line views/todoId.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/todoId.html:68
	WriteTodoId(qb422016, todo, statusList)
//line views/todoId.html:68
	qs422016 := string(qb422016.B)
//line views/todoId.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/todoId.html:68
	return qs422016
//line views/todoId.html:68
}
