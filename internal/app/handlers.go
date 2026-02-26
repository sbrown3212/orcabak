package app

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func HandlerStatus() {
	cmd := exec.Command("git", "status")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error running 'git status': %s", err)
	}
	fmt.Print(out.String())
}
