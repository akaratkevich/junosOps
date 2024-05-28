package internal

import (
	"flag"
	"fmt"
	"os"
)

func SetupFlags() (username *string, password *string, threshold *string, err error) {
	// Define flags
	username = flag.String("u", "", "Username for device access")
	password = flag.String("p", "", "Password for device access")
	threshold = flag.String("t", "2m", "Time threshold for interface flaps (e.g., 2m for 2 minutes, 2h for 2 hours, 2d for 2 days, 3M for 3 months, 33w4d for 33 weeks and 4 days)")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse the command line flags
	flag.Parse()

	// Validate the input flags
	err = validateFlags(username, password, threshold)
	return
}

func validateFlags(username *string, password *string, threshold *string) error {
	// Validate required flags
	if *username == "" {
		return fmt.Errorf("error: Username is required. Please provide a username with --u (eg. --u admin)")
	}
	if *password == "" {
		return fmt.Errorf("error: Password is required. Please provide a password with --p (eg. --p password)")
	}
	if *threshold == "" {
		return fmt.Errorf("error: Threshold is required. Please provide a time threshold with --t (e.g., --t 2m)")
	}

	// Validate the threshold format
	if _, err := ParseDuration(*threshold); err != nil {
		return fmt.Errorf("error: Invalid threshold format. %v", err)
	}
	return nil
}
