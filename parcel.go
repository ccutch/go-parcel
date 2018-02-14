package main

import "os/exec"

func buildFrontend() {
	cmd := exec.Command("npm", "run", "build")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func watchFrontend() {
	cmd := exec.Command("npm", "run", "watch")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
