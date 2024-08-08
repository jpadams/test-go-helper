// Test of using a helper golang program

package main

import (
	"dagger/test/internal/dagger"
)

type Test struct{}

// Returns Container with output/ of processing src/ with go program in helper/
func (m *Test) Foo(src *dagger.Directory, helper *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:latest").
		WithMountedDirectory("/src", src).
		WithMountedDirectory("/helper", helper).
		WithWorkdir("/helper").
		WithExec([]string{"go", "run", "main.go"})
}
