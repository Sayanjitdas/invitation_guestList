package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	type InvitedUser struct {
		FirstName   string
		LastName    string
		phoneNumber string
	}

	var nameList []InvitedUser

	loop := true

	fmt.Println("[+]WELCOME TO INVITATION LIST KEEPER")

	for loop {

		var choice int

		//ENTRY PRINT STATEMENT
		fmt.Println("[+]TOTAL GUESTS =>", len(nameList))
		fmt.Println("[+]PRESS 1 TO SEARCH FOR AN INVITATION")
		fmt.Println("[+]PRESS 2 TO ENTER NEW INVITATION")
		fmt.Println("[+]PRESS 3 TO VIEW LIST OF ALL GUESTS")
		fmt.Println("[+]PRESS 0 TO QUIT")

		//INPUTS
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			clearScreen()
			//name variable
			var searchNumber string

			fmt.Println("[+]ENTER THE PHONE NUMBER YOU WANT TO SEARCH")
			fmt.Scanln(&searchNumber)

			for _, name := range nameList {
				if name.phoneNumber == searchNumber {
					fmt.Println(name.FirstName + " " + name.LastName)
					fmt.Println("[+] --INVITED--")
					fmt.Println()
				}
			}

		case 2:
			clearScreen()

			//user name input
			var firstName string
			var lastName string
			var phoneNumber string
			var yesNo int
			var user InvitedUser

			fmt.Println("[+]ENTER YOUR NAME")
			fmt.Printf("FIRSTNAME: ")
			fmt.Scanln(&firstName)
			fmt.Printf("LASTNAME: ")
			fmt.Scanln(&lastName)
			fmt.Printf("PHONENUMBER: ")
			fmt.Scanln(&phoneNumber)

			user.FirstName = firstName
			user.LastName = lastName
			user.phoneNumber = phoneNumber

			fmt.Println("[+]YOUR ENTRY")
			fmt.Println("\t FIRSTNAME : ", user.FirstName)
			fmt.Println("\t LASTNAME: ", user.LastName)
			fmt.Println("\t PHONE NUMBER: ", user.phoneNumber)
			fmt.Println("[+] DO YOU CONFIRM ?? 1 TO CONFIRM / 0 TO CANCEL")
			fmt.Scanln(&yesNo)

			if yesNo == 1 {
				nameList = append(nameList, user)
				fmt.Println("[+] --CONFIRMED--")
			} else {
				fmt.Println("[-] --CANCELLED--")
			}

		case 3:
			clearScreen()
			if len(nameList) > 0 {
				fmt.Println("ID" + strings.Repeat(" ", 30-2) + "FIRSTNAME" + strings.Repeat(" ", 30-9) + "LASTNAME" + strings.Repeat(" ", 30-8) + "PHONENUMBER")
				fmt.Println(strings.Repeat("-", 100))
				for i, user := range nameList {
					t := strconv.Itoa(i)
					i := len(t)
					f := len(user.FirstName)
					l := len(user.LastName)
					id := t + strings.Repeat(" ", 30-i)
					firstname := user.FirstName + strings.Repeat(" ", 30-f)
					lastname := user.LastName + strings.Repeat(" ", 30-l)
					fmt.Println(id + firstname + lastname + user.phoneNumber)
				}
			} else {
				fmt.Println("[-] --GUEST LIST EMPTY--")
				fmt.Println()
			}

		case 0:
			fmt.Println("[+] QUITTING...")
			loop = false
		}

	}
}
