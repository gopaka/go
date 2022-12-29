package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v remaining tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{}
	for {
	      var firstName string
	      var lastName string
	      var email string
	      var userTickets uint

		  
	      //ask the user their name

	      fmt.Printf("\nPlease entter your First name\n")
	      fmt.Scan(&firstName)
          
	      fmt.Printf("\nPlease enter your Last name\n")
	       fmt.Scan(&lastName)
       
	       fmt.Printf("\nPlease enter yout email id\n")
	       fmt.Scan(&email)
       
	       fmt.Printf("\nHow many ticket do you want\n")
	       fmt.Scan(&userTickets)
 
		   if userTickets <= remainingTickets{
	          remainingTickets = remainingTickets - userTickets
	          // bookings = firstName + " " + lastName      - used in array
	          bookings = append(bookings, firstName+" "+lastName) //slice
    
		      // get first name from slice
    
		      firstNames := []string{}
		      for _, booking := range bookings{
		 	     var names = strings.Fields(booking)
		 	     firstNames = append(firstNames, ".", names[0])
		        }
              fmt.Printf("\nName of bookings are %v", firstNames)
          
	          //   fmt.Printf("\nThank you for booking tickets with us %v %v", firstName, lastName)
		      fmt.Printf("\nThe Remaining ticket is %v", remainingTickets)
    
		      if remainingTickets == 0 {
				fmt.Printf("\n Booking is full")
				break
			  }
		 	   
		    } else {	 
			    fmt.Printf("\nYou cant book more ticket than remaining ticket %v", remainingTickets)
		 	}
			

		
			
		  
		}  
	}

    
