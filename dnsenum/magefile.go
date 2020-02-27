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
	mg.Deps(DnsEnum04)
	mg.Deps(DnsEnum05)
	mg.Deps(DnsEnum06)
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

// Build dnsenum04
func DnsEnum04() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum04.go")
}

// Build dnsenum05
func DnsEnum05() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum05.go")
}

// Build dnsenum06
func DnsEnum06() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum06.go")
}

// Remove project artifacts
func Clean() {
	projects := []string{"dnsenum01", "dnsenum02", "dnsenum03", "dnsenum04", "dnsenum05", "dnsenum06"}

	for _, project := range projects {
		sh.Rm(project)
	}
}
