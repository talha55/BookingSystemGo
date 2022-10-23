package main

import "fmt"

func main (){
	const companyName = "FiveRivers Technologies"
	const totalBookings = 50
	var remainingBookings = totalBookings
	fmt.Printf("Welcom To %v \n", companyName)
	fmt.Printf("We have %v booking slots available out of %v \n\n", remainingBookings, totalBookings)
	var customerName string
	var selectedSlots int
	customerName = "Talha Muneer"
	selectedSlots = 4
	fmt.Printf("Hi %v,\nYou've selected %v Slots.\n", customerName, selectedSlots)
}
