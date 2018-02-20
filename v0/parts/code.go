// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parts

import (
	"encoding/json"
	"go/format"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/google/shenzhen-go/v0/source"
)

const codePartEditTemplateSrc = `
<input type="hidden" id="hhead" name="Head" value="">
<input type="hidden" id="hbody" name="Body" value="">
<input type="hidden" id="htail" name="Tail" value="">
<script>
	function switchto(e) {
		h = document.getElementById('headtab');
		b = document.getElementById('bodytab');
		t = document.getElementById('tailtab');
		x = document.getElementById(e);
		h.style.display = 'none';
		b.style.display = 'none';
		t.style.display = 'none';
		x.style.display = 'block';
	}
</script>
<a href="javascript:void(0)" onclick="switchto('headtab')">Head</a> |
<a href="javascript:void(0)" onclick="switchto('bodytab')">Body</a> |
<a href="javascript:void(0)" onclick="switchto('tailtab')">Tail</a>
<div id="headtab" style="display:none">
	<h4>Head</h4>
	<pre class="codeedit" id="head">{{.Node.ImplHead}}</pre>
</div>
<div id="bodytab" style="display:block">
	<h4>Body</h4>
	<pre class="codeedit" id="body">{{.Node.ImplBody}}</pre>
</div>
<div id="tailtab" style="display:none">
	<h4>Tail</h4>
	<pre class="codeedit" id="tail">{{.Node.ImplTail}}</pre>
</div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.2.6/ace.js" type="text/javascript" charset="utf-8"></script>
<script>
    var theme = "ace/theme/chrome";
	var lang = "ace/mode/golang";
    var head = ace.edit("head");
	var body = ace.edit("body");
	var tail = ace.edit("tail");
    head.setTheme(theme);
    head.getSession().setMode(lang);
	head.getSession().setUseSoftTabs(false);
    body.setTheme(theme);
    body.getSession().setMode(lang);
	body.getSession().setUseSoftTabs(false);
    tail.setTheme(theme);
    tail.getSession().setMode(lang);
	tail.getSession().setUseSoftTabs(false);
	this.parent.onsubmit = function() {
		document.getElementById("hhead").value = head.getValue();
		document.getElementById("hbody").value = body.getValue();
		docuemnt.getElementById("htail").value = tail.getValue();
	};
</script>
`

// Code is a component containing arbitrary code.
type Code struct {
	Head, Body, Tail string
}

type jsonCode struct {
	Head []string `json:"head"`
	Body []string `json:"body"`
	Tail []string `json:"tail"`
}

// MarshalJSON encodes the Code component as JSON.
func (c *Code) MarshalJSON() ([]byte, error) {
	k := &jsonCode{
		Head: strings.Split(c.Head, "\n"),
		Body: strings.Split(c.Body, "\n"),
		Tail: strings.Split(c.Tail, "\n"),
	}
	stripCR(k.Head)
	stripCR(k.Body)
	stripCR(k.Tail)
	return json.Marshal(k)
}

// UnmarshalJSON decodes the Code component from JSON.
func (c *Code) UnmarshalJSON(j []byte) error {
	var mp jsonCode
	if err := json.Unmarshal(j, &mp); err != nil {
		return err
	}
	h := strings.Join(mp.Head, "\n")
	b := strings.Join(mp.Body, "\n")
	t := strings.Join(mp.Tail, "\n")
	if err := c.refresh(h, b, t); err != nil {
		// refresh would error if it can't get used channels,
		// say, because of a syntax error preventing parsing.
		// Returning error would stop the graph being loaded,
		// and the user might need to fix their syntax error,
		// via the interface...
		log.Printf("Couldn't format or determine channels used: %v", err)
	}
	return nil
}

// LineCount is number of lines in c.Head, c.Body, c.Tail
// (conveneince function for templates.)
func (c *Code) LineCount() struct{ H, B, T int } {
	return struct{ H, B, T int }{
		H: strings.Count(c.Head, "\n") + 3,
		B: strings.Count(c.Body, "\n") + 3,
		T: strings.Count(c.Tail, "\n") + 3,
	}
}

// AssociateEditor adds a "part_view" template to the given template.
func (c *Code) AssociateEditor(tmpl *template.Template) error {
	_, err := tmpl.New("part_view").Parse(codePartEditTemplateSrc)
	return err
}

// Channels returns nil - all channel usage is expected to be auto extracted.
func (c *Code) Channels() (read, written source.StringSet) { return nil, nil }

// Clone returns a copy of this Code part.
func (c *Code) Clone() interface{} {
	return &Code{
		Head: c.Head,
		Body: c.Body,
		Tail: c.Tail,
	}
}

// Help returns a helpful explanation.
func (*Code) Help() template.HTML {
	return `<p>
	A Code part runs (executes) any Go code that you write.
	</p><p>
	It consists of 3 parts: a Head, a Body, and a Tail. 
	The Head runs first, and only runs once, no matter what number Multiplicity is set to. 
	The Body runs next. The number of concurrent copies of the Body that run is set by Multiplicity.
	Finally, when all copies of the Body return, the Tail runs. 
	</p><p>
	The Head and Tail are useful for operations that should only be done once. For example, any 
	output channels written to in the Body can be correctly closed (if desired) in the Tail.
	</p><p>
	Each instance of the Body can use the int parameters <code>instanceNumber</code> and <code>multiplicity</code>
	to distinguish which instance is running and how many are running, if necessary. 
	<code>0 <= instanceNumber < multiplicity</code>
	</p><p>
	Any channels referred to will automatically be detected and shown in the graph, and
	when channels are renamed, these will be safely updated in the Code where they are
	referred to.
	</p><p>
	The <code>return</code> statement is allowed but optional in Code. There are no values that
	need to be returned.
	Using <code>return</code> in the Head will prevent the Body or Tail from executing, but 
	using <code>return</code> in the Body won't affect whether the Tail is executed.
	</p>
	`
}

// Impl returns the implementation of the goroutine.
func (c *Code) Impl() (Head, Body, Tail string) {
	return c.Head, c.Body, c.Tail
}

// Imports returns a nil slice.
func (*Code) Imports() []string { return nil }

// RenameChannel does fancy footwork to rename the channel in the code,
// with a side-effect of nicely formatting it. If a rename issue occurs
// e.g. because the user's code has a syntax error, the rename is aborted
// and logged.
func (c *Code) RenameChannel(from, to string) {
	h, err := source.RenameChannel(c.Head, "Head", from, to)
	if err != nil {
		log.Printf("Couldn't do rename on Head: %v", err)
		return
	}
	b, err := source.RenameChannel(c.Body, "Body", from, to)
	if err != nil {
		log.Printf("Couldn't do rename on Body: %v", err)
		return
	}
	t, err := source.RenameChannel(c.Tail, "Tail", from, to)
	if err != nil {
		log.Printf("Couldn't do rename on Tail: %v", err)
		return
	}
	c.Head, c.Body, c.Tail = h, b, t
}

// TypeKey returns "Code".
func (*Code) TypeKey() string { return "Code" }

func (c *Code) refresh(h, b, t string) error {
	// At least save what the user entered.
	c.Head, c.Body, c.Tail = h, b, t

	// Try to format it.
	hf, err := format.Source([]byte(h))
	if err != nil {
		return err
	}
	bf, err := format.Source([]byte(b))
	if err != nil {
		return err
	}
	tf, err := format.Source([]byte(t))
	if err != nil {
		return err
	}

	c.Head, c.Body, c.Tail = string(hf), string(bf), string(tf)
	return nil
}

// Update sets relevant fields based on the given Request.
func (c *Code) Update(r *http.Request) error {
	h, b, t := c.Head, c.Body, c.Tail
	if r != nil {
		h, b, t = r.FormValue("Head"), r.FormValue("Body"), r.FormValue("Tail")
	}
	return c.refresh(h, b, t)
}
