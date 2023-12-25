package main

import "fmt"

func main() {
	mybill := newBill("Yoma's Bill")

	mybill.updateTip(10)

	// displaying final bill 
	fmt.Println(mybill.format())
}