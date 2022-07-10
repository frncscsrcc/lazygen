package main

import (
	"os/exec"

	"github.com/frncscsrcc/lazygen/generator"
)

func main() {
	filesToFormat := generator.Generate("entities.yaml")
	for _, file := range filesToFormat {
		if err := exec.Command("go", "fmt", file).Run(); err != nil {
			panic(err)
		}
	}
}
