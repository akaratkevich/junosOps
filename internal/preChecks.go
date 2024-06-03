package internal

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

const (
	PreChecksHeader = "************** PRE/POST - CHECKS **************"
	PreChecksFooter = "=====================================\n"
)

var preChecks = []string{
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
}

// PreChecks writes pre-check commands and header to the config file
func PreChecks(configFile *os.File, deviceHost string) error {
	currentTime := time.Now().Format("02-01-2006 15:04:05")
	username := "unknown"

	if usr, err := user.Current(); err == nil {
		username = usr.Username
	}

	header := fmt.Sprintf(`============ DEVICE CONFIGURATION ============
Device Hostname: %s
Date: %s
Generated by: %s
`, deviceHost, currentTime, username)

	_, err := fmt.Fprint(configFile, header)
	if err != nil {
		log.Printf("Failed to write header to config file for device %s: %v", deviceHost, err)
		return err
	}

	_, err = fmt.Fprintf(configFile, "%s\n", PreChecksHeader)
	if err != nil {
		log.Printf("Failed to write pre-checks header to config file for device %s: %v", deviceHost, err)
		return err
	}

	for _, cmd := range preChecks {
		_, err := fmt.Fprintf(configFile, "%s\n", cmd)
		if err != nil {
			log.Printf("Failed to write pre-checks commands to config file for device %s: %v", deviceHost, err)
			return err
		}
	}

	_, err = fmt.Fprintf(configFile, "%s\n", PreChecksFooter)
	if err != nil {
		log.Printf("Failed to write pre-checks footer to config file for device %s: %v", deviceHost, err)
		return err
	}

	return nil
}
