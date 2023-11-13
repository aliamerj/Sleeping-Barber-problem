package main

import (
	"time"

	"github.com/fatih/color"
)

// Define the Barber Shop struct with its properties
type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

// Add a new barber to the shop and simulate their behavior
func (b *BarberShop) addBarber(barberName string) {
	b.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for client", barberName)
		for {
			// Check if there are clients; if not, the barber goes to sleep
			if len(b.ClientsChan) == 0 {
				color.Blue("Nobody is here , %s goes to sleep ", barberName)
				isSleeping = true
			}
			// Await new client or shop closure
			client, shopOpen := <-b.ClientsChan
			if shopOpen {
				// Wake up sleeping barber and start haircut
				if isSleeping {
					color.HiYellow("%s wakes %s up", client, barberName)
				}
				// Simulate the haircut process
				color.HiGreen("%s is cutting hair of %s", barberName, client)
				time.Sleep(b.HairCutDuration)
				color.Green("%s is finished cutting hair of %s", barberName, client)
			} else {
				// Barber leaves if the shop is closed
				color.Cyan("%s go home now ", barberName)
				b.BarberDoneChan <- true
				return
			}
		}
	}()
}

// Close the shop for the day, signaling all barbers to stop working
func (b *BarberShop) closeShopForDay() {
	color.HiCyan("Closing shop for day")
	close(b.ClientsChan) // Close the client channel
	b.Open = false       // Mark shop as closed
	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarberDoneChan // Wait for each barber to acknowledge closure
	}
	close(b.BarberDoneChan) // Close the barber done channel
	color.Red("The Barber Shop is now closed for today")
}

// Add a new client to the shop
func (b *BarberShop) addNewClint(clientName string) {
	// Announce the arrival of a new client
	color.Green("*** %s arrives", clientName)
	if b.Open {
		select {
		case b.ClientsChan <- clientName:
			// Client sits in the waiting room if there's space
			color.Blue("%s take a seat in waiting room", clientName)
		default:
			// Client leaves if the waiting room is full
			color.Red("The waiting room is full, so %s leave now", clientName)
		}
	} else {
		// Client leaves if the shop is already closed
		color.Red("The shop is already closed, so %s leave !", clientName)
	}
}
