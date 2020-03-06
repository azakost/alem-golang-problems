package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("curl", "https://google.com").Output()
	fmt.Print(string(out))
}
