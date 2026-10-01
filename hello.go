package main

import "fmt"

func main() {
	confranceName := "Go confrance"
	const confranceTickets int = 50
	var reminingTickets uInt= 7

	//fmt.Printf("confranceName is %T, confranceTickets is %T and reminingTickets is %T", confranceName, confranceTickets, reminingTickets)
	reminingTickets = -1
	fmt.Printf("Welcome to our %v booking applicaion\n",confranceName )
	fmt.Printf("we have %v and %v remining tickets now\n", confranceTickets, reminingTickets)
	fmt.Print("Get your ticket's here to attend!\n")

	var userName string
	var userTickets int  

	userName = "Kiarash"
	userTickets = 25
	fmt.Printf("Your user name is %v and your ticket number is %v.\n", userName, userTickets)
}