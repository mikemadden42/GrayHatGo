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
	return nil
}

// Build portscan01
func DnsEnum01() error {
	return sh.Run("go", "build", "-ldflags", ldflags, "dnsenum01.go")
}

// Remove project artifacts
func Clean() {
	projects := []string{"dnsenum01"}

	for _, project := range projects {
		sh.Rm(project)
	}
}
