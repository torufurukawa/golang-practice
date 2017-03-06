package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "ok.go", "OK")
	stdout, err := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Scan() => ", scanner.Scan())
	fmt.Println("line: ", scanner.Text())
}
