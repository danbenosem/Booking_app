package main

import ("fmt"
    "Booking_app/models"
 
)






func  CollectDetails() model.User{
	var user model.User


	fmt.Println("Welcome to Booking Application")
	fmt.Println("")
	fmt.Println("Enter your first Name:")
	fmt.Scanln(&user.FirstName)

	fmt.Println("Enter your lastName")
	fmt.Scanln(&user.LastName)

	fmt.Println("Enter your Email")
	fmt.Scanln(&user.Email)

	fmt.Println("Enter the number of ticket")
	fmt.Scanln(&user.TicketNumber)




	return user



	  
	}