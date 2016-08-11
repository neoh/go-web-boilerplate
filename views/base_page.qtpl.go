// This file is automatically generated by qtc from "base_page.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line ../views/base_page.qtpl:1
package views

//line ../views/base_page.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// This is a base page template. All the other template pages implement this interface.
//

//line ../views/base_page.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../views/base_page.qtpl:4
type Page interface {
	//line ../views/base_page.qtpl:4
	Title() string
	//line ../views/base_page.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
	//line ../views/base_page.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
	//line ../views/base_page.qtpl:4
	Body() string
	//line ../views/base_page.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
	//line ../views/base_page.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line ../views/base_page.qtpl:4
}

// Page prints a page implementing Page interface.

//line ../views/base_page.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	//line ../views/base_page.qtpl:12
	qw422016.N().S(`
<html>
	<head>
	    <meta data-auth-token="52:AS6[rxSdsMd4RgYXJgeabsRAVBZ:0406139009]"/>

		<title>`)
	//line ../views/base_page.qtpl:17
	p.StreamTitle(qw422016)
	//line ../views/base_page.qtpl:17
	qw422016.N().S(`</title>
	</head>
	<body>
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
	//line ../views/base_page.qtpl:23
	p.StreamBody(qw422016)
	//line ../views/base_page.qtpl:23
	qw422016.N().S(`
	</body>
</html>
`)
//line ../views/base_page.qtpl:26
}

//line ../views/base_page.qtpl:26
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	//line ../views/base_page.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line ../views/base_page.qtpl:26
	StreamPageTemplate(qw422016, p)
	//line ../views/base_page.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line ../views/base_page.qtpl:26
}

//line ../views/base_page.qtpl:26
func PageTemplate(p Page) string {
	//line ../views/base_page.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
	//line ../views/base_page.qtpl:26
	WritePageTemplate(qb422016, p)
	//line ../views/base_page.qtpl:26
	qs422016 := string(qb422016.B)
	//line ../views/base_page.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
	//line ../views/base_page.qtpl:26
	return qs422016
//line ../views/base_page.qtpl:26
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line ../views/base_page.qtpl:31
type BasePage struct{}

//line ../views/base_page.qtpl:32
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line ../views/base_page.qtpl:32
qw422016.N().S(`This is a base title`) }

//line ../views/base_page.qtpl:32
//line ../views/base_page.qtpl:32
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line ../views/base_page.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line ../views/base_page.qtpl:32
	p.StreamTitle(qw422016)
	//line ../views/base_page.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line ../views/base_page.qtpl:32
}

//line ../views/base_page.qtpl:32
func (p *BasePage) Title() string {
	//line ../views/base_page.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
	//line ../views/base_page.qtpl:32
	p.WriteTitle(qb422016)
	//line ../views/base_page.qtpl:32
	qs422016 := string(qb422016.B)
	//line ../views/base_page.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
	//line ../views/base_page.qtpl:32
	return qs422016
//line ../views/base_page.qtpl:32
}

//line ../views/base_page.qtpl:33
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line ../views/base_page.qtpl:33
qw422016.N().S(`This is a base body`) }

//line ../views/base_page.qtpl:33
//line ../views/base_page.qtpl:33
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
	//line ../views/base_page.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line ../views/base_page.qtpl:33
	p.StreamBody(qw422016)
	//line ../views/base_page.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line ../views/base_page.qtpl:33
}

//line ../views/base_page.qtpl:33
func (p *BasePage) Body() string {
	//line ../views/base_page.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line ../views/base_page.qtpl:33
	p.WriteBody(qb422016)
	//line ../views/base_page.qtpl:33
	qs422016 := string(qb422016.B)
	//line ../views/base_page.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line ../views/base_page.qtpl:33
	return qs422016
//line ../views/base_page.qtpl:33
}