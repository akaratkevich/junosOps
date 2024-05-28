package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func ConnectAndExecute(device Device, commands []string) ([]string, error) {
	config := &ssh.ClientConfig{
		User: device.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(device.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%s", device.Host, device.Port)
	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	var results []string
	for _, cmd := range commands {
		session, err := conn.NewSession()
		if err != nil {
			return nil, fmt.Errorf("failed to create session: %v", err)
		}
		defer session.Close()

		output, err := session.CombinedOutput(cmd)
		if err != nil {
			return nil, fmt.Errorf("failed to run command: %v", err)
		}

		results = append(results, string(output))
	}

	return results, nil
}
