package main

import "fmt"

func main() {
	var confranceName = "Go confrance"
	const confranceTickets = 50
	var reminingTickets = 50
	
	fmt.Printf("Welcome to our %v booking applicaion\n",confranceName )
	fmt.Print("we have %v and %v remining tickets now\n", confranceTickets, reminingTickets)
	fmt.Print("Get your ticket's here to attend")
}

