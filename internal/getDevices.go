package internal

import (
	"bufio"
	"fmt"
	"os"
)

func GetDevices() ([]string, error) {
	var devices []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter the list of devices (end with Ctrl+D or an empty line):")

	for {
		if !scanner.Scan() {
			break
		}
		device := scanner.Text()
		if device == "" {
			break
		}
		devices = append(devices, device)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return devices, nil
}
