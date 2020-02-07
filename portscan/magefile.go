//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var ldflags = "-s -w"

// Create project artifacts
func Build() error {
	mg.Deps(Clean)
	mg.Deps(Portscan01)
	mg.Deps(Portscan02)
	return nil
}

// Build portscan01
func Portscan01() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan01.go")
}

// Build portscan02
func Portscan02() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan02.go")
}

// Build portscan02
func Portscan03() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan03.go")
}

// Remove project artifacts
func Clean() {
	sh.Rm("portscan01")
	sh.Rm("portscan02")
	sh.Rm("portscan03")
}
