package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ProcessDevice(device Device, command string, threshold time.Duration) (int, error) {
	results, err := ConnectAndExecute(device, command)
	if err != nil {
		return 0, fmt.Errorf("Error connecting to device %s: %v", device.Host, err)
	}

	interfaceDataList, err := ParseInterfaceXML(results, device.Host)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse XML for device %s: %v", device.Host, err)
	}

	// Open a file for writing the interface data
	file, err := os.Create(fmt.Sprintf("%s_interfaces.txt", device.Host))
	if err != nil {
		return 0, fmt.Errorf("Failed to create file for device %s: %v", device.Host, err)
	}
	defer file.Close()

	count := 0

	// Write the updated InterfaceData if LastFlapped is longer than the threshold
	for _, data := range interfaceDataList {
		lastFlappedSeconds, err := strconv.ParseInt(data.LastFlapped, 10, 64)
		if err != nil {
			log.Printf("Skipping interface %s on device %s: %v", data.Interface, device.Host, err)
			continue
		}

		if time.Duration(lastFlappedSeconds)*time.Second > threshold {
			_, err := fmt.Fprintf(file, "Interface: %s\nDescription: %s\nLast Flapped: %s\n\n", data.Interface, data.Description, data.LastFlapped)
			if err != nil {
				log.Printf("Failed to write to file for device %s: %v", device.Host, err)
			}
			count++
		}
	}
	return count, nil
}
