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

package view

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/google/shenzhen-go/model"
)

// TODO: Replace these cobbled-together UIs with Polymer or something.
const channelEditorTemplateSrc = `<head>
	<title>Channel</title>
	<link type="text/css" rel="stylesheet" href="/.static/fonts.css">
	<link type="text/css" rel="stylesheet" href="/.static/main.css">
</head>
<body>
	<h1>Channel</h1>
	{{if not .New}}
	<a href="?channel={{.Index}}&clone">Clone</a> | 
	<a href="?channel={{.Index}}&delete">Delete</a>
	{{end}}
	<form method="post">
		<input type="hidden" name="New" value="{{.New}}>
		<div class="formfield">
			<label for="Type">Type</label>
			<input type="text" name="Type" required value="{{.Type}}">
		</div>
		<div class="formfield">
			<label for="Cap">Capacity</label>
			<input type="number" name="Cap" required pattern="^[0-9]+$" title="Must be a whole number, at least 0." value="{{.Cap}}">
		</div>
		<table>
			<thead>
				<tr>
					<th>Connection</th>
					<th></th>
				</tr>
			</thead>
			{{range $conn := .Connections -}}
			<tr>
				<td>
					<select>
					{{range $node := $.Graph.Nodes }}
					{{range $arg, $type := $node.InputArgs}}
					{{if eq $type $.Channel.Type}}
					{{$val := printf "%s.%s" $node $arg}}
						<option value="{{$val}}" {{if eq $val $conn.String}}selected{{end}}>{{$val}}</option>
					{{end -}}
					{{end -}}
					{{range $arg, $type := $node.OutputArgs}}
					{{if eq $type $.Channel.Type}}
					{{$val := printf "%s.%s" $node $arg}}
						<option value="{{$val}}" {{if eq $val $conn.String}}selected{{end}}>{{$val}}</option>
					{{end -}}
					{{end -}}
					{{end -}}
					</select>
				</td>
				<td><a href="javascript:void(0)" onclick="removeconn('{{$conn}}}')">Remove connection</td></tr>
			{{end}}
		</table>
		<div class="formfield hcentre">
			<input type="submit" value="Save">
			<input type="button" value="Return" onclick="window.location.href='?'">
		</div>
	</form>
</body>`

var (
	channelEditorTemplate = template.Must(template.New("channelEditor").Parse(channelEditorTemplateSrc))

	identifierRE = regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*$`)
)

// Channel handles viewing/editing a channel.
func Channel(g *model.Graph, name string, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)

	q := r.URL.Query()
	_, clone := q["clone"]
	_, del := q["delete"]

	var e *model.Channel
	if name == "new" {
		if clone || del {
			http.Error(w, "Asked for a new channel, but also to clone or delete the channel", http.StatusBadRequest)
			return
		}
		e = new(model.Channel)
	} else {
		e := g.Channels[name]
		if e == nil {
			http.Error(w, fmt.Sprintf("Channel %q not found", name), http.StatusNotFound)
			return
		}
	}

	switch {
	case clone:
		e2 := *e
		e = &e2
	case del:
		delete(g.Channels, e.Name)
		u := *r.URL
		u.RawQuery = ""
		log.Printf("redirecting to %v", &u)
		http.Redirect(w, r, u.String(), http.StatusSeeOther) // should cause GET
		return
	}

	var err error
	switch r.Method {
	case "POST":
		err = handleChannelPost(g, e, w, r)
	case "GET":
		err = channelEditorTemplate.Execute(w, &struct {
			*model.Graph
			*model.Channel
			New bool
		}{g, e, name == "new"})
	default:
		err = fmt.Errorf("unsupported verb %q", r.Method)
	}

	if err != nil {
		log.Printf("Could not handle request: %v", err)
		http.Error(w, "Could not handle request", http.StatusInternalServerError)
	}
}

func handleChannelPost(g *model.Graph, e *model.Channel, w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Validate.
	ci, err := strconv.Atoi(r.FormValue("Cap"))
	if err != nil {
		return err
	}
	if ci < 0 {
		return fmt.Errorf("invalid capacity [%d < 0]", ci)
	}

	// Update.
	nu := r.FormValue("New") == "true"
	name := r.FormValue("Name")
	e.Anonymous = (name == "")
	if e.Anonymous && e.Name == "" {
		for {
			name = fmt.Sprintf("anonymous_channel_%d", rand.Int())
			if _, exists := g.Channels[name]; !exists {
				break
			}
		}
	}
	e.Name = name
	e.Type = r.FormValue("Type")
	e.Capacity = ci
	if !nu {
		return channelEditorTemplate.Execute(w, e)
	}
	g.Channels[name] = e
	q := url.Values{
		"channel": []string{name},
	}
	u := *r.URL
	u.RawQuery = q.Encode()
	log.Printf("redirecting to %v", u)
	http.Redirect(w, r, u.String(), http.StatusSeeOther) // should cause GET
	return nil
}
