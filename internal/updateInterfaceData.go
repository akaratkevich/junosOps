package internal

import "strings"

func UpdateInterfaceData(data *InterfaceData, commandResults []string) {
	for _, result := range commandResults {
		lines := strings.Split(result, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Description:") {
				data.Description = strings.TrimSpace(strings.Split(line, ":")[1])
			} else if strings.Contains(line, "VLAN:") {
				data.DownSince = strings.TrimSpace(strings.Split(line, ":")[1])
			} else if strings.Contains(line, "Status:") {
				data.Status = strings.TrimSpace(strings.Split(line, ":")[1])
			}
		}
	}
}
