package main

import "fmt"


type User struct{

   firstName string
   LastName string
   email string
   ticketNumber int


}



func  CollectDetails() User{
	var user User


	fmt.Println("Welcome to Booking Application")
	fmt.Println("")
	fmt.Println("Enter your first Name:")
	fmt.Scanln(&user.firstName)

	fmt.Println("Enter your lastName")
	fmt.Scanln(&user.LastName)

	fmt.Println("Enter your Email")
	fmt.Scanln(&user.email)

	return user



	  
	}