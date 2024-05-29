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
	username, password, threshold, err := internal.SetupFlags()
	if err != nil {
		logger.Fatal("Exiting the program due to setup failure", logger.Args("Reason", err))
		os.Exit(1)
	}

	// 2. Parse the threshold stamp
	thresholdSeconds, err := internal.ParseThreshold(*threshold)
	if err != nil {
		logger.Fatal("Failed to parse time threshold", logger.Args("Reason", err))
		os.Exit(1)
	}

	// 3. Get a list of devices from the user
	devices, err := internal.GetDevices()
	if err != nil {
		fmt.Printf("Failed to get input devices: %v\n", err)
		logger.Warn("Failed to get input devices", logger.Args("Reason", err))
		return
	}

	logger.Info("Working with the following devices:", logger.Args("Devices", devices))

	command := "show interfaces | display xml"

	// 4. Setup concurrency
	const workerCount = 10 // Define the number of workers
	jobs := make(chan internal.Device, len(devices))
	results := make(chan int, len(devices))

	var totalCount int
	var totalInterfacesProcessed int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Worker function
	worker := func() {
		for device := range jobs {
			interfacesProcessed, count, err := internal.ProcessDevice(device, command, thresholdSeconds)
			if err != nil {
				log.Printf("Failed to process device %s: %v", device.Host, err)
				results <- 0
				continue
			}
			mu.Lock()
			totalCount += count
			totalInterfacesProcessed += interfacesProcessed
			mu.Unlock()
			results <- count
		}
		wg.Done()
	}

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker()
	}

	// Send devices to workers
	go func() {
		for _, host := range devices {
			device := internal.Device{
				Host:     host,
				Port:     "22", // Default port
				Username: *username,
				Password: *password,
			}
			jobs <- device
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	for range results {
		// Results are already handled in the worker function
	}

	// ------------------- Reporting --------------------------------
	elapsedTime := time.Since(startTime)
	fmt.Println("\n----------------------------------------------------------------")
	pterm.FgLightYellow.Printf("Number of devices processed: %d\n", len(devices))
	pterm.FgLightYellow.Printf("Total interfaces processed: %d\n", totalInterfacesProcessed)
	pterm.FgLightYellow.Printf("Total count of 'down' interfaces with last flap exceeding the threshold: %d\n", totalCount)
	pterm.FgLightYellow.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Println("\n----------------------------------------------------------------\n")
	// Signature
	pterm.FgGray.Println("Developed by: Anton Karatkevich")
	pterm.FgGray.Println("Contact: karatkevich.anton@gmail.com")
	pterm.FgGray.Println("Project Repository: https://github.com/akaratkevich/junosOps.git")
}
