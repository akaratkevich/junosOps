package internal

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ProcessDevice(device Device, command string) (int, error) {
	results, err := ConnectAndExecute(device, command)
	if err != nil {
		return 0, fmt.Errorf("error connecting to device %s: %v", device.Host, err)
	}

	interfaceDataList, err := ParseInterfaceXML(results, device.Host)
	if err != nil {
		return 0, fmt.Errorf("failed to parse XML for device %s: %v", device.Host, err)
	}

	// Open a file for writing the interface data
	file, err := os.Create(fmt.Sprintf("%s_interfaces.txt", device.Host))
	if err != nil {
		return 0, fmt.Errorf("failed to create file for device %s: %v", device.Host, err)
	}
	defer file.Close()

	count := 0

	// Write the updated InterfaceData if LastFlapped is longer than 2 minutes
	for _, data := range interfaceDataList {
		duration, err := ParseFlappedTime(data.LastFlapped)
		if err != nil {
			log.Printf("Skipping interface %s on device %s: %v", data.Interface, device.Host, err)
			continue
		}

		if duration > 2*time.Minute {
			count++
			_, err := fmt.Fprintf(file, "Interface: %s\nDescription: %s\nLast Flapped: %s\n\n", data.Interface, data.Description, data.LastFlapped)
			if err != nil {
				log.Printf("Failed to write to file for device %s: %v", device.Host, err)
			}
		}
	}

	return count, nil
}
