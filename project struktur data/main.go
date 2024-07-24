package main

import (
	"MVC_member/View"
	"fmt"
)

func menu() {
	for {
		fmt.Println("Menu Program")
		fmt.Println("1. Insert Member")
		fmt.Println("2. Update Member")
		fmt.Println("3. Delete Member")
		fmt.Println("4. Read All Member")
		fmt.Println("5. Search Member")
		fmt.Println("6. Exit")
		fmt.Println("----------------")
		fmt.Print("masukkan menu pilihan anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			View.InsertMember()
		case 2:
			View.UpdateMember()
		case 3:
			View.DeleteMember()
		case 4:
			View.ReadAllMember()
		case 5:
			View.SearchMember()
		case 6:
			fmt.Println("exit program")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func main() {
	menu()
}
