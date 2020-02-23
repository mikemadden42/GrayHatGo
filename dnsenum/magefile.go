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
	mg.Deps(DnsEnum01)
	mg.Deps(DnsEnum02)
	mg.Deps(DnsEnum03)
	return nil
}

// Build dnsenum01
func DnsEnum01() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum01.go")
}

// Build dnsenum02
func DnsEnum02() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum02.go")
}

// Build dnsenum03
func DnsEnum03() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum03.go")
}

// Remove project artifacts
func Clean() {
	projects := []string{"dnsenum01", "dnsenum02", "dnsenum03"}

	for _, project := range projects {
		sh.Rm(project)
	}
}
