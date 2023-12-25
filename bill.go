package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// new bills
func newBill(name string) bill {
	localBill := bill{
		name: name,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}

	return localBill
}

// BILL FORMAT
func (bReceiverFunc *bill) format() string {
	// bReceiverFunc limits the format options to just the bill object
	formattedString := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range bReceiverFunc.items {
		formattedString += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total = total + v

	}

	// Adding tip
	formattedString += fmt.Sprintf("%-25v ...$%v\n", "tip:", bReceiverFunc.tip)

	// outputting the total line
	formattedString += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + bReceiverFunc.tip)

	return formattedString
}

// UPDATE TIP FUNC
func (bTipUpdate *bill) updateTip(tip float64){
	bTipUpdate.tip = tip
}

// ADD AN ITEM TO THE BILL
func (bAddItemToBill *bill) addItem(name string, price float64) {
	bAddItemToBill.items[name] = price
}