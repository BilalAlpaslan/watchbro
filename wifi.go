package main

import (
	"log"
	"os/exec"
)

func wifiOnOff(b bool) {
	status := "enable" 
	if !b {
		status = "disable"
	}

	cmd := exec.Command("netsh", "interface", "set", "interface", "Wi-Fi", "admin="+status)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}