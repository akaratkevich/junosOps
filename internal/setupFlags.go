package internal

import (
	"flag"
	"fmt"
	"os"
)

// Return pointers to the values of the flags
func SetupFlags() (username *string, password *string, err error) {
	// Define flags
	username = flag.String("u", "", "Username for device access")
	password = flag.String("p", "", "Password for device access")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse the command line flags
	flag.Parse()

	// Validate the input flags for all other cases
	err = validateFlags(username, password)
	return
}

// Check if necessary flags are provided; return an error if any are missing.
func validateFlags(username *string, password *string) error {
	// Validate required flags
	if *username == "" {
		return fmt.Errorf("error: Username is required. Please provide a username with --u (eg. --u admin)")
	}
	if *password == "" {
		return fmt.Errorf("error: Password is required. Please provide a password with --p (eg. --p password)")
	}
	return nil
}
