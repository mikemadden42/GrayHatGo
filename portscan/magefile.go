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
	mg.Deps(Portscan03)
	mg.Deps(Portscan04)
	mg.Deps(Portscan05)
	mg.Deps(Portscan06)
	mg.Deps(Portscan07)
	mg.Deps(Portscan08)
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

// Build portscan03
func Portscan03() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan03.go")
}

// Build portscan04
func Portscan04() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan04.go")
}

// Build portscan05
func Portscan05() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan05.go")
}

// Build portscan06
func Portscan06() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan06.go")
}

// Build portscan07
func Portscan07() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan07.go")
}

// Build portscan08
func Portscan08() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "portscan08.go")
}

// Remove project artifacts
func Clean() {
	sh.Rm("portscan01")
	sh.Rm("portscan02")
	sh.Rm("portscan03")
	sh.Rm("portscan04")
	sh.Rm("portscan05")
	sh.Rm("portscan06")
	sh.Rm("portscan07")
	sh.Rm("portscan08")
}
