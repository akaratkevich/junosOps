package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"junosOps/internal"
	"log"
	"os"
	"sync"
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
	username, password, thresholdStr, err := internal.SetupFlags()
	if err != nil {
		logger.Fatal("Exiting the program due to setup failure", logger.Args("Reason", err))
		os.Exit(1)
	} else {
		logger.Trace("Successfully passed the parameters for setup.")
	}

	// Parse the threshold duration
	threshold, err := internal.ParseDuration(*thresholdStr)
	if err != nil {
		logger.Fatal("Failed to parse time threshold", logger.Args("Reason", err))
		os.Exit(1)
	}

	// 2. Get a list of devices from the user
	devices, err := internal.GetDevices()
	if err != nil {
		fmt.Printf("Failed to get input devices: %v\n", err)
		logger.Warn("Failed to get input devices", logger.Args("Reason:", err))
		return
	}

	logger.Info("Working with the following devices:", logger.Args("Devices", devices))

	// 3. Setup concurrency
	var totalCount int
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	// 4. Connect and execute
	command := "show interfaces extensive | display xml"

	for _, host := range devices {
		device := internal.Device{
			Host:     host,
			Port:     "22", // Default port
			Username: *username,
			Password: *password,
		}

		wg.Add(1)
		go func(device internal.Device) {
			defer wg.Done()
			count, err := internal.ProcessDevice(device, command, threshold)
			if err != nil {
				log.Printf("Failed to process device %s: %v", device.Host, err)
				return
			}
			mu.Lock()
			totalCount += count
			mu.Unlock()
		}(device)
	}

	wg.Wait()

	logger.Info("Total interfaces that flapped for more than the threshold:", logger.Args("Count", totalCount))

	// ------------------- Reporting --------------------------------
	elapsedTime := time.Since(startTime)
	fmt.Println("\n----------------------------------------------------------------")
	pterm.FgLightYellow.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Println("\n----------------------------------------------------------------")
	// Signature
	pterm.FgGray.Println("Developed by: Anton Karatkevich")
	pterm.FgGray.Println("Contact: anton.karatkevich@virginmediao2.com")
	pterm.FgGray.Println("Project Repository: https://github.com/akaratkevich/junosOps.git")
}
