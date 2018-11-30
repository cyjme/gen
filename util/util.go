package util

import (
	"log"
	"os/exec"
)

func FormatSourceCode(filename string) {
	log.Println("format:", filename)
	cmd := exec.Command("gofmt", "-w", filename)
	if err := cmd.Run(); err != nil {
		log.Printf("Error while running gofmt: %s", err)
	}
}
