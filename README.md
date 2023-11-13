# Sleeping Barber Problem Simulation
## Overview
This Go project simulates the classic concurrency problem known as the "Sleeping Barber Problem." The simulation demonstrates how a barbershop operates using concurrent programming concepts. It features multiple barbers serving clients who arrive at random intervals. The barbershop has a limited seating capacity; clients leave if there's no available seat. The project uses Go's concurrency features, like goroutines and channels, to manage the complex interactions between clients and barbers.

## Problem Description
In the Sleeping Barber Problem, a barber sleeps when there are no clients. When a client arrives, they either wake the barber if he's sleeping or wait if the barber is busy and seats are available. If the shop is full, the client leaves. The challenge is to coordinate the barber's and clients' actions without deadlocks or resource contention.

## Project Structure
### main.go: 
Contains the main function that initializes and runs the barbershop simulation.
### barberShop.go: 
Defines the BarberShop struct and associated methods for adding barbers and clients, and managing the shop's state.
## How It Works
The barbershop has a certain number of barbers and a seating capacity for waiting clients.
Barbers continuously check for clients and either go to sleep if none are available or start serving a client.
Clients arrive at random intervals and either take a seat, leave if the shop is full, or leave if the shop is closed.
The simulation runs for a predetermined amount of time before the shop closes for the day.
## Key Concepts Demonstrated
Concurrency: Managing multiple barbers and clients concurrently.
Synchronization: Coordinating actions between barbers and clients.
Resource Allocation: Managing limited resources (barber chairs and waiting seats).

## Setup and Running the Simulation
Installation: Ensure you have Go installed on your machine.
Running the Program:
Navigate to the project directory.
Run the command go run main.go barberShop.go.
Observe the simulation in the terminal.
Customization
You can modify the following parameters in main.go to alter the simulation:

seatingCapacity: Number of seats available for waiting clients.
arrivalRate: Controls the rate at which clients arrive.
cutDuration: Duration it takes for a barber to complete a haircut.
timeOpen: How long the barbershop stays open.
Dependencies
github.com/fatih/color: Used for colored terminal output to enhance readability.
