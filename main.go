package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	currentTime := time.Now()
	git_command := fmt.Sprintf("git add . && git commit -m '%s' && git push", currentTime.Format("02-01-06"))
	cmd := exec.Command("cmd", "/C", git_command)
	cmd.Dir = "" // Adicione o diretório do obsidian

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Worked!")
}
