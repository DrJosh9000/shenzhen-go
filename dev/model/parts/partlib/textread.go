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

package partlib

import (
	"bufio"
	"os"
)

// FileLine represents a single line from a text file. It includes the file path,
// the line itself, and the line number starting from 0.
type FileLine struct {
	Path, Line string
	LineNum    int
}

// StreamTextFile reads files to read from an input channel, and streams
// text lines from the file as strings to the output.
func StreamTextFile(pathIn <-chan string, output chan<- FileLine, errors chan<- error) {
	for path := range pathIn {
		f, err := os.Open(path)
		if err != nil {
			errors <- err
			continue
		}
		defer f.Close()
		// TODO: switch to bufio.Reader since it can handle longer lines.
		sc := bufio.NewScanner(f)
		ln := 0
		for sc.Scan() {
			output <- FileLine{
				Path:    path,
				Line:    sc.Text(),
				LineNum: ln,
			}
			ln++
		}
		if err := sc.Err(); err != nil {
			errors <- err
			continue
		}
		if err := f.Close(); err != nil {
			errors <- err
		}
	}
}
