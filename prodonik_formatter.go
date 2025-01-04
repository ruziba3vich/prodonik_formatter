package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("arc", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running `arc status`:", err)
		return
	}

	var modifiedFiles []string
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "modified:") {
			path := strings.TrimSpace(strings.TrimPrefix(line, "modified:"))
			modifiedFiles = append(modifiedFiles, path)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading output:", err)
		return
	}

	for _, file := range modifiedFiles {
		fmt.Printf("Formatting file: %s\n", file)
		formatCmd := exec.Command("ya", "tool", "tt", "format", file)
		formatCmd.Stdout = os.Stdout
		formatCmd.Stderr = os.Stderr
		if err := formatCmd.Run(); err != nil {
			fmt.Printf("Error formatting file %s: %v\n", file, err)
		} else {
			fmt.Printf("Successfully formatted: %s\n", file)
		}
	}
}
