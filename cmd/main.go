package main

import (
	"log"
	"os/exec"

	"github.com/frncscsrcc/lazygen/generator"
)

func main() {
	filesToFormat := generator.Generate("entities.yaml")
	for _, file := range filesToFormat {
		if err := exec.Command("go", "fmt", file).Run(); err != nil {
			panic(err)
		} else {
			log.Printf("go fmt %s\n", file)
		}
	}
}
