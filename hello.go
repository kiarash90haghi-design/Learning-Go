package main

import "fmt"

func main() {
	var confranceName = "Go confrance"
	const confranceTickets = 50
	var reminingTickets = 50
	
	fmt.Println("Welcome to our ", confranceName, " booking applicaion")
	fmt.Println("we have ", confranceTickets, " confrance tickets and ", reminingTickets, " remining tickets now")
	fmt.Print("Get your ticket's here to attend")
}

