package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func runPowerShellCommand(command string) string {
	out, err := exec.Command("powershell.exe", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {

	service := "wuauserv"
	response := runPowerShellCommand("Get-Service " + service)

	matched, err := regexp.MatchString("Stopped", response)
	fmt.Println(matched)
	if err != nil {
		log.Fatal(err)
	}

	if matched {
		runPowerShellCommand("Start-Service " + service)
	}
}
