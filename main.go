package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetNetworkTime retrieves the current network time from the specified NTP server.
func GetNetworkTime(serverIP string) (time.Time, error) {
	ntpTime, err := ntp.Time(serverIP)
	if err != nil {
		return time.Now(), err
	}
	return ntpTime, nil
}

func main() {
	// You can use any valid NTP server IP or domain.
	// Here's an example with the NTP pool project.
	ntpServer := os.Getenv("SERVER")

	ntpTime, err := GetNetworkTime(ntpServer)
	if err != nil {
		fmt.Printf("Error retrieving network time(%s): %v\n", ntpServer, err)
	} else {
		fmt.Printf("Current network time(%s): %v\n", ntpServer, ntpTime)
	}
}
