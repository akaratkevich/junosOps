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

	commands := []string{
		"show interface description",
		"show vlan",
		"show interface status",
	}

	for _, host := range devices {
		device := internal.Device{
			Host:     host,
			Port:     "22", // Default port
			Username: *username,
			Password: *password,
		}

		results, err := internal.ConnectAndExecute(device, commands)
		if err != nil {
			log.Fatalf("Error connecting to device %s: %v", device.Host, err)
		}

		data := &internal.InterfaceData{Node: device.Host}
		internal.UpdateInterfaceData(data, results)

		// Print the updated InterfaceData
		fmt.Printf("Device: %s\n", device.Host)
		fmt.Printf("Description: %s\n", data.Description)
		fmt.Printf("VLAN: %s\n", data.DownSince)
		fmt.Printf("Status: %s\n", data.Status)
	}
	// ------------------- Reporting --------------------------------
	elapsedTime := time.Since(startTime)
	fmt.Println("\n----------------------------------------------------------------")
	pterm.FgLightYellow.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Println("\n----------------------------------------------------------------")

}
