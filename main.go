package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Define global variables for simulation settings
var (
	seatingCapacity = 10                      // Capacity of the barber shop
	arrivalRate     = 100                     // Control rate of client arrival
	cutDuration     = 1000 * time.Millisecond // Duration for a haircut
	timeOpen        = 10 * time.Second        // Total open time for the shop
)

func main() {
	// Seed the random number generator for client arrival times
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Print a welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Red("----------------------------")

	// Create channels for client management and barber completion
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// Initialize the barber shop struct with its properties
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}
	color.Green("The shop is open for the day!")

	// Add barbers to the shop
	shop.addBarber("Ali")
	shop.addBarber("Amer")
	shop.addBarber("Jeff")

	// Prepare for shop closing after a set duration
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// Continuously add clients at random intervals until the shop starts closing
	c := 1
	go func() {
		for {
			// Calculate a random arrival time for clients
			randomeMillsecond := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomeMillsecond)):
				shop.addNewClint(fmt.Sprintf("Client #%d", c))
				c++
			}
		}
	}()

	// Wait until the shop is officially closed
	<-closed

	// Extra wait time to ensure all processes finish
	time.Sleep(5 * time.Second)
}
