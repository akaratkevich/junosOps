package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ProcessDevice(device Device, command string, threshold time.Duration) (int, int, error) {
	results, err := ConnectAndExecute(device, command)
	if err != nil {
		return 0, 0, fmt.Errorf("Error connecting to device %s: %v", device.Host, err)
	}

	interfaceDataList, err := ParseInterfaceXML(results, device.Host)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to parse XML for device %s: %v", device.Host, err)
	}

	// Open a file for writing the interface data
	file, err := os.Create(fmt.Sprintf("%s_interfaces_audit.txt", device.Host))
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to create file for device %s: %v", device.Host, err)
	}
	defer file.Close()

	// Open file for writing configuration data
	configFileName := fmt.Sprintf("config_%s.txt", device.Host)
	configFile, err := os.Create(configFileName)
	if err != nil {
		log.Printf("Failed to create config file for device %s: %v", device.Host, err)
		return 0, 0, err
	}
	defer configFile.Close()

	// Write additional commands at the beginning of the config file
	preChecks := []string{
		"============ PRE - CHECKS ===========",
		"show configuration | display set",
		"show interfaces descriptions",
		"show interfaces terse | no-more",
		"show interfaces brief",
		"show bridge domain brief",
		"show bridge mac-table",
		"show lacp interfaces",
		"show lldp neighbors",
		"show evpn instance",
		"show evpn database state duplicate",
		"=====================================\n",
	}

	for _, cmd := range preChecks {
		_, err := fmt.Fprintf(configFile, "%s\n", cmd)
		if err != nil {
			log.Printf("Failed to write pre-checks commands to config file for device %s: %v", device.Host, err)
			return 0, 0, err
		}
	}

	count := 0

	// Write the updated InterfaceData if there is a Description, LastFlapped is longer than the threshold and the status is "down"
	for _, data := range interfaceDataList {
		if data.Description == "" {
			continue
		}

		if strings.ToLower(data.Status) != "down" {
			continue
		}

		flapped, err := ParseFlappedStamp(data.LastFlapped)
		if err != nil {
			log.Printf("Skipping interface %s on device %s: %v", data.Interface, device.Host, err)
			continue
		}

		if flapped > threshold {
			count++
			_, err := fmt.Fprintf(file, "Interface: %s\nDescription: %s\nStatus: %s\nLast Flapped: %s\n\n", data.Interface, data.Description, data.Status, data.LastFlapped)
			if err != nil {
				log.Printf("Failed to write to file for device %s: %v", device.Host, err)
				continue
			}
			_, err = fmt.Fprintf(configFile, "set interfaces %s description \"%s Decommissioned %s\"\n", data.Interface, data.Description, time.Now().Format("2006-01-02"))
			if err != nil {
				log.Printf("Failed to write to config file for device %s: %v", device.Host, err)
				continue
			}
			_, err = fmt.Fprintf(configFile, "set interfaces %s disable\n", data.Interface)
			if err != nil {
				log.Printf("Failed to write to config file for device %s: %v", device.Host, err)
				continue
			}
		}
	}

	return len(interfaceDataList), count, nil
}
