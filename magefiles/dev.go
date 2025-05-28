package main

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/mg"
)

type Dev mg.Namespace

// InstallDependencies installs development tools required for the project.
func (Dev) InstallDependencies() error {
	commands := [][]string{
		{"brew", "install", "pre-commit"},
		{"brew", "install", "bufbuild/buf/buf"},
		{"pre-commit", "install"},
	}

	for _, cmd := range commands {
		fmt.Printf("Running: %s\n", cmd)
		c := exec.Command(cmd[0], cmd[1:]...)
		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to run %v: %w", cmd, err)
		}
	}

	return nil
}

// BufGenerate builds the image and then produces results according to `buf.gen.yaml`.
func (Dev) BufGenerate() error {
	fmt.Println("Running buf generate...")
	c := exec.Command("buf", "generate")
	if err := c.Run(); err != nil {
		return fmt.Errorf("failed to run buf generate: %w", err)
	}
	return nil
}

// BufBuild compiles protobuf files into an internal image format, verifying that the definitions are syntactically and semantically correct (e.g., resolving types, imports, etc.).
func (Dev) BufBuild() error {
	fmt.Println("Running buf build...")
	c := exec.Command("buf", "build")
	if err := c.Run(); err != nil {
		return fmt.Errorf("failed to run buf build: %w", err)
	}
	return nil
}
