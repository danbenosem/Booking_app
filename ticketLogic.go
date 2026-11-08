package main

import (
	 "fmt"
	
	 "Booking_app/models"
)


var TicketNumber int




func TicketLogic(user model.User){

	var remainingTicket int =Tickets-user.TicketNumber

	fmt.Printf("The current number of ticket is %v\n",remainingTicket)


	
}