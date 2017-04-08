// Copyright 2017 Google Inc.
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

package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type fakeServer struct {
	graph, node string
	x, y        int
}

func (f *fakeServer) SetPosition(req *SetPositionRequest) error {
	f.graph = req.Graph
	f.node = req.Node
	f.x = req.X
	f.y = req.Y
	return nil
}

func TestSetPositionRoundTrip(t *testing.T) {
	fs := &fakeServer{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Dispatch(fs, w, r)
	}))
	defer srv.Close()

	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("httptest is doing something weird: %v", err)
	}
	cl := NewClient(u)

	req := &SetPositionRequest{
		NodeRequest: NodeRequest{
			Request: Request{
				Graph: "graph",
			},
			Node: "node",
		},
		X: 300,
		Y: 500,
	}

	if err := cl.SetPosition(req); err != nil {
		t.Fatalf("SetPosition = %v, want nil error", err)
	}

	if got, want := *fs, (fakeServer{"graph", "node", 300, 500}); got != want {
		t.Errorf("fakeServer = %#v, want %#v", got, want)
	}
}
