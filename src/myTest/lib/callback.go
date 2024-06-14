package lib

import (
	"fmt"
	"time"
)

// Callback function signature
type SecretUpdateCallback func(secret string)

// Function to monitor dynamic secrets
func MonitorSecrets(callback SecretUpdateCallback) {
	// Simulating periodic retrieval of secrets
	for {
		// Retrieve secret from secret manager
		secret := "new_secret"
		// Trigger callback with updated secret
		callback(secret)
		// Wait for a specified interval before retrieving the next secret
		time.Sleep(5 * time.Second)
	}
}

// Example callback function
func UpdateSecret(secret string) {
	fmt.Println("Received updated secret:", secret)
	// Perform actions with the updated secret
}

func CallbackTest() {
	// Launching a goroutine to monitor secrets and trigger callbacks
	go MonitorSecrets(UpdateSecret)
	// Main function continues execution
	fmt.Println("Main function executing...")
	// Waiting indefinitely to keep the program running
}
