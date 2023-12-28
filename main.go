package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//helper function
func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err 
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	// fmt.Println("Created the bill - ", b.name, "'s bill")
	fmt.Printf("Created the bill - %v's bill\n", b.name)

	return b
}

func prompOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	if(opt != "s" && opt != "a" && opt != "t"){
		fmt.Println("ENTER VALID INPUT")
		prompOptions(b)
	}else{
		if(opt == "a"){
			name, _ := getInput("Item name: ", reader)
			price, _ := getInput("Item price: ", reader)

			p, err := strconv.ParseFloat(price, 64)
			if(err != nil){
				fmt.Println("The Price must be a number")
				prompOptions(b)
			}

			// update price if there's no error
			b.addItem(name, p)

			fmt.Println("Items added - ",  name, price)
			prompOptions(b)
		}else if(opt == "s"){
			fmt.Println("SAVE BILL")
		}else if(opt == "t"){
			tip, _ := getInput("Enter tip amount ($): ", reader)
			fmt.Println(tip)
		}
	}
}

func main() {
	// mybill := newBill("Yoma's Bill")

	mybill := createBill()
	prompOptions(mybill)

	fmt.Println(mybill)
}