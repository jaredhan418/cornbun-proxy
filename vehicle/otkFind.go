package vehicle

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/teslamotors/vehicle-command/pkg/account"
	"github.com/teslamotors/vehicle-command/pkg/protocol"
)

func OtkFind(vin string, oauthToken string, keyPath string) {
	logger := log.New(os.Stderr, "", 0)

	// Specify the user-agent header value used in HTTP requests to Tesla's servers. The default
	// value is constructed from your package name and account.LibraryVersion.
	userAgent := "cornbun.fun/1.0.0"

	if vin == "" {
		logger.Printf("Must specify VIN")
		return
	}

	// Since commands are authenticated end-to-end, they need to be authorized with a private key.
	// The corresponding public key must be enrolled on the vehicle's keychain. See the README.md
	// file in the root directory for pointers on setting all this up.
	privateKey, err := protocol.LoadPrivateKey(keyPath)
	if err != nil {
		logger.Printf("Failed to load private key: %s", err)
		return
	}

	if err != nil {
		logger.Printf("Failed to load OAuth token: %s", err)
		return
	}

	// For simplicity, allow 30 seconds to wake up the vehicle, connect to it, and unlock. In
	// practice you'd want a fresh timeout for each command.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// This example program sends commands over the Internet, which requires a Tesla account login
	// token. The protocol can also work over BLE; see other programs in the example directory.
	acct, err := account.New(string(oauthToken), userAgent)
	if err != nil {
		logger.Printf("Authentication error: %s", err)
		return
	}

	car, err := acct.GetVehicle(ctx, vin, privateKey, nil)
	if err != nil {
		logger.Printf("Failed to fetch vehicle info from account: %s", err)
		return
	}

	// Some commands can be sent while the vehicle is offline and some accounts have multiple
	// vehicles. So connecting to the vehicle is a separate step.
	fmt.Println("Connecting to car...")
	if err := car.Connect(ctx); err != nil {
		logger.Printf("Failed to connect to vehicle: %s\n", err)
		return
	}

	// The above code authenticates with Tesla. However, most commands require the client to
	// authenticate directly to the car as well. StartSession() performs a handshake with the
	// vehicle that allows subsequent commands to be authenticated.
	if err := car.StartSession(ctx, nil); err != nil {
		logger.Printf("Failed to perform handshake with vehicle: %s\n", err)
		return
	}

	fmt.Println("Start FlashLight Set Sentry...")
	if err := car.FlashLights(ctx); err != nil {
		if protocol.MayHaveSucceeded(err) {
			logger.Printf("Open Frunk command sent, but client could not confirm receipt: %s\n", err)
		} else {
			logger.Printf("Failed to Open vehicle frunk: %s\n", err)
		}
		return
	}
	if err := car.SetSentryMode(ctx, true); err != nil {
		if protocol.MayHaveSucceeded(err) {
			logger.Printf("Open Frunk command sent, but client could not confirm receipt: %s\n", err)
		} else {
			logger.Printf("Failed to Open vehicle frunk: %s\n", err)
		}
		return
	}
	fmt.Println("Vehicle FlashLight Set Sentry!")
}