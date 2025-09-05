package main

import (
	"fmt"
	"log"

	"github.com/bugfix4u/govee-go/govee"
)

func main() {
	// Example usage of the Govee client
	// Create a new client from environment variables
	// The GOVEE_API_KEY environment variable must be set
	client, err := govee.NewFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Get Devices
	devices, err := client.GetDevices()
	if err != nil {
		log.Fatal(err)
	}

	deviceStateMap := make(map[string]*govee.DeviceInfo)

	//Get Device State
	for _, device := range devices {
		state, err := client.GetDeviceState(device)
		if err != nil {
			log.Printf("Error getting state for device %s: %v", device.Device, err)
			continue
		}
		deviceStateMap[device.Device] = state
	}

	for _, device := range devices {
		state, exists := deviceStateMap[device.Device]
		if exists {
			fmt.Printf("Model: %s\nDevice: %s\nName: %s\nType: %s\n", device.SKU, device.Device, device.DeviceName, device.Type)
			for _, capability := range state.Capabilities {
				fmt.Printf("Capability Type: %s\nInstance: %s\nValue: %v\n", capability.Type, capability.Instance, capability.State.Value)
			}
		}
		fmt.Print("\n------------------------------------------------------------------------------------------\n")
	}
	client.Close() // Close the client when done
}
