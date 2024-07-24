package View

import (
	"MVC_member/Model"
	"MVC_member/Node"
	"fmt"
)

func InsertMember() {
	var id int
	var nama, alamat string
	var point float32
	fmt.Print("Enter ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Name: ")
	fmt.Scan(&nama)
	fmt.Print("Enter Address: ")
	fmt.Scan(&alamat)
	fmt.Print("Enter Points: ")
	fmt.Scan(&point)
	Model.InsertMember(id, nama, alamat, point)
	fmt.Println("Member inserted successfully.")
}

func UpdateMember() {
	var id int
	var nama, alamat string
	var point float32
	fmt.Print("Enter ID of the member to update: ")
	fmt.Scan(&id)
	fmt.Print("Enter new Name: ")
	fmt.Scan(&nama)
	fmt.Print("Enter new Address: ")
	fmt.Scan(&alamat)
	fmt.Print("Enter new Points: ")
	fmt.Scan(&point)
	member := Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point}
	if Model.UpdateMember(member) {
		fmt.Println("Member updated successfully.")
	} else {
		fmt.Println("Member not found.")
	}
}

func DeleteMember() {
	var id int
	fmt.Print("Enter ID of the member to delete: ")
	fmt.Scan(&id)
	if Model.DeleteMember(id) {
		fmt.Println("Member deleted successfully.")
	} else {
		fmt.Println("Member not found.")
	}
}

func ReadAllMember() {
	members := Model.ReadAllMember()
	if members == nil {
		fmt.Println("No members found.")
		return
	}
	for members != nil {
		fmt.Printf("ID: %d, Name: %s, Address: %s, Points: %.2f\n", members.Data.Id, members.Data.Nama, members.Data.Alamat, members.Data.Point)
		members = members.Next
	}
}

func SearchMember() {
	var id int
	fmt.Print("Enter ID of the member to search: ")
	fmt.Scan(&id)
	member := Model.SearchMember(id)
	if member == nil {
		fmt.Println("Member not found.")
	} else {
		fmt.Printf("ID: %d, Name: %s, Address: %s, Points: %.2f\n", member.Data.Id, member.Data.Nama, member.Data.Alamat, member.Data.Point)
	}
}
