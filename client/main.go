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

// The client script is for interacting with a graph in an SVG (via DOM).
// The browser gets it from the server (served "statically") and it makes
// gRPC-Web API calls into the server.
package main

import (
	"fmt"
	"log"

	"github.com/google/shenzhen-go/client/view"
	"github.com/google/shenzhen-go/jsutil"
	pb "github.com/google/shenzhen-go/proto/js"
	"github.com/gopherjs/gopherjs/js"
)

var (
	graphPath = js.Global.Get("graphPath").String()
	apiURL    = js.Global.Get("apiURL").String()
)

func main() {
	doc := jsutil.CurrentDocument()
	client := pb.NewShenzhenGoClient(apiURL)
	initial := js.Global.Get("GraphJSON").String()
	if err := view.Setup(doc, client, initial); err != nil {
		log.Fatalf("Couldn't load graph: %v", err)
	}

}
