// Code generated by qtc from "alpineMain.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/alpineMain.html:1
package views

//line views/alpineMain.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/alpineMain.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/alpineMain.html:1
func StreamAlpineMain(qw422016 *qt422016.Writer) {
//line views/alpineMain.html:1
	qw422016.N().S(`
<h1
    class="mb-5 text-3xl font-bold leading-tight text-teal-500 dark:text-teal-500 sm:text-4xl sm:leading-tight md:text-5xl md:leading-tight text-center">
    Alpine</h1>
<p>Hello Alpine</p>

`)
//line views/alpineMain.html:7
	StreamAlpineCountToggle(qw422016)
//line views/alpineMain.html:7
	qw422016.N().S(`

`)
//line views/alpineMain.html:9
	StreamAlpineSearchInput(qw422016)
//line views/alpineMain.html:9
	qw422016.N().S(`

`)
//line views/alpineMain.html:11
}

//line views/alpineMain.html:11
func WriteAlpineMain(qq422016 qtio422016.Writer) {
//line views/alpineMain.html:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/alpineMain.html:11
	StreamAlpineMain(qw422016)
//line views/alpineMain.html:11
	qt422016.ReleaseWriter(qw422016)
//line views/alpineMain.html:11
}

//line views/alpineMain.html:11
func AlpineMain() string {
//line views/alpineMain.html:11
	qb422016 := qt422016.AcquireByteBuffer()
//line views/alpineMain.html:11
	WriteAlpineMain(qb422016)
//line views/alpineMain.html:11
	qs422016 := string(qb422016.B)
//line views/alpineMain.html:11
	qt422016.ReleaseByteBuffer(qb422016)
//line views/alpineMain.html:11
	return qs422016
//line views/alpineMain.html:11
}

//line views/alpineMain.html:15
func StreamAlpineCountToggle(qw422016 *qt422016.Writer) {
//line views/alpineMain.html:15
	qw422016.N().S(`
<br>
<strong>Count Toggle Transitions</strong>
<div x-data="{ open: false, count: 0 }" class="transition ease-out duration-1000 bg-orange-200 bg-green-200"
    hx-ext="class-tools" classes="toggle bg-green-200:1s, toggle bg-orange-200:1s">
    <button @click="open = ! open;" x-text="open?'Hide':'Show'" class="btn btn-primary"></button>
    <br>
    <button @click="count++;" class="btn btn-secondary">Count</button>
    <br>
    <div x-show="open" x-transition>
        <span x-text="'Count is ' + count "></span>
    </div>
</div>
<br>
`)
//line views/alpineMain.html:29
}

//line views/alpineMain.html:29
func WriteAlpineCountToggle(qq422016 qtio422016.Writer) {
//line views/alpineMain.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/alpineMain.html:29
	StreamAlpineCountToggle(qw422016)
//line views/alpineMain.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/alpineMain.html:29
}

//line views/alpineMain.html:29
func AlpineCountToggle() string {
//line views/alpineMain.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/alpineMain.html:29
	WriteAlpineCountToggle(qb422016)
//line views/alpineMain.html:29
	qs422016 := string(qb422016.B)
//line views/alpineMain.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/alpineMain.html:29
	return qs422016
//line views/alpineMain.html:29
}

//line views/alpineMain.html:33
func StreamAlpineSearchInput(qw422016 *qt422016.Writer) {
//line views/alpineMain.html:33
	qw422016.N().S(`
<br>
<div x-data="{
        search: '',
        items: ['foo', 'bar', 'baz'],
        get filteredItems() {
            return this.items.filter(
                i => i.startsWith(this.search)
            )
        }
    }">

    <strong>Search Input with getters</strong>
    <br>
    <input x-model="search" placeholder="Search...">
    <ul>
        <template x-for="item in filteredItems" :key="item">
            <li x-text="item"></li>
        </template>
    </ul>

    <strong>Search Input with x-show</strong>
    <br>
    <input x-model="search" placeholder="Search...">
    <ul>
        <template x-for="item in items" :key="item">
            <li x-text="item" x-show="item.startsWith(search)"></li>
        </template>
    </ul>


</div>
`)
//line views/alpineMain.html:65
}

//line views/alpineMain.html:65
func WriteAlpineSearchInput(qq422016 qtio422016.Writer) {
//line views/alpineMain.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/alpineMain.html:65
	StreamAlpineSearchInput(qw422016)
//line views/alpineMain.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/alpineMain.html:65
}

//line views/alpineMain.html:65
func AlpineSearchInput() string {
//line views/alpineMain.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/alpineMain.html:65
	WriteAlpineSearchInput(qb422016)
//line views/alpineMain.html:65
	qs422016 := string(qb422016.B)
//line views/alpineMain.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/alpineMain.html:65
	return qs422016
//line views/alpineMain.html:65
}