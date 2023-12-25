package main

type bill struct{
	name string
	items map[string]float64
	tip float64
}

// new bills 
func newBill(name string) bill{
	localBill := bill{
		name: name,
		items: map[string]float64{}, //empty map
		tip: 0,
	}

	return localBill
}