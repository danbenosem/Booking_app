package main

import "fmt"

func main() {

        

    fmt.Println("Welcome to Booking Application")
    const tickets = 50


    fmt.Println("How many tickets do you want to buy?")
    var user_input int
    fmt.Scan(user_input)
    remaining_ticket:= 50-user_input

    fmt.Print("The number of tickets remaining is %v",remaining_ticket)






}