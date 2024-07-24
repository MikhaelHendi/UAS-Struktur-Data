package Model

import (
	"MVC_member/Database"
	"MVC_member/Node"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := &Node.TableMember{
		Data: Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	if Database.DBmember == nil {
		Database.DBmember = newDataMember
	} else {
		tempLL := Database.DBmember
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	return Database.DBmember
}

func DeleteMember(id int) bool {
	if Database.DBmember == nil {
		return false
	}
	if Database.DBmember.Data.Id == id {
		Database.DBmember = Database.DBmember.Next
		return true
	}

	tempLL := Database.DBmember
	for tempLL.Next != nil {
		if tempLL.Next.Data.Id == id {
			tempLL.Next = tempLL.Next.Next
			return true
		}
		tempLL = tempLL.Next
	}
	return false
}

func SearchMember(id int) *Node.TableMember {
	tempLL := Database.DBmember
	for tempLL != nil {
		if tempLL.Data.Id == id {
			return tempLL
		}
		tempLL = tempLL.Next
	}
	return nil
}

func UpdateMember(member Node.Member) bool {
	addr := SearchMember(member.Id)
	if addr == nil {
		return false
	}
	addr.Data.Nama = member.Nama
	addr.Data.Alamat = member.Alamat
	addr.Data.Point = member.Point
	return true
}
