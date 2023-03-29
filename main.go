package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	hargaVip := 30
	hargaRegular := 20
	hargaStudent := 10

	totalPrice := (VIP * hargaVip) + (regular * hargaRegular) + (student * hargaStudent)
	totalTicket := VIP + regular + student
	diskon := 0

	if totalPrice >= 100 {
		// jika tanggal ganjil
		if day % 2 != 0 {
			if totalTicket >= 5 {
				diskon = (totalPrice * 25/100)
			} else {
				diskon = (totalPrice * 15/100)
			}
		} else {
			if totalTicket >= 5 {
				diskon = (totalPrice * 20/100)
			} else {
				diskon = (totalPrice * 10/100)
			}
		}
	}

	return float32(totalPrice - diskon) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
}
