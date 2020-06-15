// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regtest

import (
	"testing"

	"golang.org/x/tools/internal/lsp/fake"
)

const simpleProgram = `
-- go.mod --
module mod.com

go 1.12
-- main.go --
package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
}`

func TestHoverSerialization(t *testing.T) {
	runner.Run(t, simpleProgram, func(t *testing.T, env *Env) {
		// Hover on an empty line.
		env.OpenFile("main.go")
		content, pos := env.Hover("main.go", fake.Pos{Line: 3, Column: 0})
		if content != nil {
			t.Errorf("got non-empty response for empty hover: %v: %v", pos, *content)
		}
	})
}
