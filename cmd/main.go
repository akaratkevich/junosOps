package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"junosOps/internal"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	// ------------------- Logging to a file --------------------------------
	logFile, err := internal.SetupLogging("junosOps-application.log")
	if err != nil {
		log.Fatalf("Error opening log filePath: %v", err)
	}
	defer logFile.Close()
	// Set the output of logs to the filePath
	log.SetOutput(logFile)
	// ---- !!! FROM THIS POINT ON, ALL LOG MESSAGES WILL BE WRITTEN TO THE FILE !!! ----

	// ------------------- Logging to screen --------------------------------
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// 1. Get username and password

	username, password, err := internal.SetupFlags()

	if err != nil {
		logger.Fatal("Exiting the program due to setup failure", logger.Args("Reason", err))
		os.Exit(1)
	} else {
		logger.Trace("Successfully passed the parameters for setup.")
	}

	// 2. Get a list of devices from the user

	devices, err := internal.GetDevices()
	if err != nil {
		fmt.Printf("Failed to get input devices: %v\n", err)
		logger.Warn("Failed to get input devices", logger.Args("Reason:", err))
		return
	}

	logger.Info("Working with the following devices:", logger.Args("Devices", devices))

	// 3. Setup concurrency TODO:

	// 4. Connect and execute
	command := "show interfaces extensive | display xml"

	for _, host := range devices {
		device := internal.Device{
			Host:     host,
			Port:     "22", // Default port
			Username: *username,
			Password: *password,
		}

		results, err := internal.ConnectAndExecute(device, command)
		if err != nil {
			log.Fatalf("Error connecting to device %s: %v", device.Host, err)
		}

		interfaceDataList, err := internal.ParseInterfaceXML(results, device.Host)
		if err != nil {
			log.Fatalf("Failed to parse XML for device %s: %v", device.Host, err)
		}

		// Open a file for writing the interface data
		file, err := os.Create(fmt.Sprintf("%s_interfaces.txt", device.Host))
		if err != nil {
			log.Fatalf("Failed to create file for device %s: %v", device.Host, err)
		}
		defer file.Close()

		// Write the updated InterfaceData if LastFlapped is longer than 2 minutes
		for _, data := range interfaceDataList {
			duration, err := internal.ParseFlappedTime(data.LastFlapped)
			if err != nil {
				log.Printf("Skipping interface %s on device %s: %v", data.Interface, device.Host, err)
				continue
			}

			if duration > 2*time.Minute {
				_, err := fmt.Fprintf(file, "Interface: %s\nDescription: %s\nLast Flapped: %s\n\n", data.Interface, data.Description, data.LastFlapped)
				if err != nil {
					log.Printf("Failed to write to file for device %s: %v", device.Host, err)
				}
			}
		}
	}
	// ------------------- Reporting --------------------------------
	elapsedTime := time.Since(startTime)
	fmt.Println("\n----------------------------------------------------------------")
	pterm.FgLightYellow.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Println("\n----------------------------------------------------------------")

}
